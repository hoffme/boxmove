package mongo

import (
	"errors"

	"github.com/hoffme/boxmove/pkg/storage/box"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type box_s struct {
	connection     *conn
	collectionName string
}

func newBoxStore(connection *conn, collectionName string) *box_s {
	return &box_s{
		connection:     connection,
		collectionName: collectionName,
	}
}

func (b *box_s) collection() *mongo.Collection {
	return b.connection.Collection(b.collectionName)
}

func (b *box_s) Find(id string) (*box.DTO, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	dto := &box.DTO{}

	err = b.collection().FindOne(b.connection.Ctx, bson.D{{
		Key:   "_id",
		Value: oid,
	}}).Decode(dto)

	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (b *box_s) Search(filter *box.Filter) ([]*box.DTO, error) {
	bsonFilter := bson.D{}

	if len(filter.ID) > 0 {
		oids := []primitive.ObjectID{}

		for _, id := range filter.ID {
			oid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return nil, err
			}

			oids = append(oids, oid)
		}

		bsonFilter = append(bsonFilter, bson.E{Key: "_id", Value: oids})
	}
	if len(filter.Name) > 0 {
		bsonFilter = append(bsonFilter, bson.E{Key: "name", Value: bson.E{Key: "$text", Value: filter.Name}})
	}
	if len(filter.ParentID) > 0 {
		bsonFilter = append(bsonFilter, bson.E{Key: "route.0", Value: filter.Name})
	}
	if len(filter.AncestorID) > 0 {
		bsonFilter = append(bsonFilter, bson.E{Key: "route", Value: bson.E{Key: "$all", Value: filter.AncestorID}})
	}

	options := options.Find()

	if len(filter.Order) > 0 {
		orderValue := -1

		if filter.Ascendent {
			orderValue = 1
		}

		options.SetSort(bson.D{{Key: filter.Order, Value: orderValue}})
	}

	if filter.Limit > 0 {
		options.SetLimit(int64(filter.Limit))
	}
	if filter.Start > 0 {
		options.SetSkip(int64(filter.Start))
	}

	cursor, err := b.collection().Find(b.connection.Ctx, bsonFilter, options)
	if err != nil {
		return nil, err
	}
	defer func() { _ = cursor.Close(b.connection.Ctx) }()

	result := []*box.DTO{}

	for cursor.Next(b.connection.Ctx) {
		var dto *box.DTO

		err := cursor.Decode(dto)
		if err != nil {
			return nil, err
		}

		result = append(result, dto)
	}

	return result, nil
}

func (b *box_s) Create(fields *box.Fields) (*box.DTO, error) {
	if fields == nil {
		return nil, errors.New("invalid params")
	}

	dto := &box.DTO{
		Name:  fields.Name,
		Route: fields.Route,
		Meta:  fields.Meta,
	}

	result, err := b.collection().InsertOne(b.connection.Ctx, bson.D{
		{Key: "name", Value: fields.Name},
		{Key: "route", Value: fields.Route},
		{Key: "meta", Value: fields.Meta},
	})
	if err != nil {
		return nil, err
	}

	dto.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return dto, nil
}

func (b *box_s) Update(id string, fields *box.Fields) (*box.DTO, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if fields == nil {
		return nil, errors.New("invalid fields")
	}

	dto, err := b.Find(id)
	if err != nil {
		return nil, err
	}

	result, err := b.collection().UpdateOne(
		b.connection.Ctx,
		bson.M{"_id": oid},
		bson.D{{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: fields.Name},
				{Key: "route", Value: fields.Route},
				{Key: "meta", Value: fields.Meta},
			},
		}},
	)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("box not found")
	}

	dto.Name = fields.Name
	dto.Route = fields.Route
	dto.Meta = fields.Meta

	return dto, nil
}

func (b *box_s) Delete(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := b.collection().DeleteOne(
		b.connection.Ctx,
		bson.M{"_id": oid},
	)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("box not found")
	}

	return nil
}

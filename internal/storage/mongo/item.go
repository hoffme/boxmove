package mongo

import (
	"errors"

	"github.com/hoffme/boxmove/pkg/storage/item"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type item_s struct {
	connection     *conn
	collectionName string
}

func newItemStore(connection *conn, collectionName string) *item_s {
	return &item_s{
		connection:     connection,
		collectionName: collectionName,
	}
}

func (i *item_s) collection() *mongo.Collection {
	return i.connection.Collection(i.collectionName)
}

func (i *item_s) Find(id string) (*item.DTO, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	dto := &item.DTO{}

	err = i.collection().FindOne(i.connection.Ctx, bson.D{{
		Key:   "_id",
		Value: oid,
	}}).Decode(dto)

	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (i *item_s) Search(filter *item.Filter) ([]*item.DTO, error) {
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

	cursor, err := i.collection().Find(i.connection.Ctx, bsonFilter, options)
	if err != nil {
		return nil, err
	}
	defer func() { _ = cursor.Close(i.connection.Ctx) }()

	result := []*item.DTO{}

	for cursor.Next(i.connection.Ctx) {
		var dto *item.DTO

		err := cursor.Decode(dto)
		if err != nil {
			return nil, err
		}

		result = append(result, dto)
	}

	return result, nil
}

func (i *item_s) Create(fields *item.Fields) (*item.DTO, error) {
	if fields == nil {
		return nil, errors.New("invalid params")
	}

	dto := &item.DTO{
		Name: fields.Name,
		Meta: fields.Meta,
	}

	result, err := i.collection().InsertOne(i.connection.Ctx, bson.D{
		{Key: "name", Value: fields.Name},
		{Key: "meta", Value: fields.Meta},
	})
	if err != nil {
		return nil, err
	}

	dto.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return dto, nil
}

func (i *item_s) Update(id string, fields *item.Fields) (*item.DTO, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if fields == nil {
		return nil, errors.New("invalid fields")
	}

	dto, err := i.Find(id)
	if err != nil {
		return nil, err
	}

	result, err := i.collection().UpdateOne(
		i.connection.Ctx,
		bson.M{"_id": oid},
		bson.D{{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: fields.Name},
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
	dto.Meta = fields.Meta

	return dto, nil
}

func (i *item_s) Delete(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := i.collection().DeleteOne(
		i.connection.Ctx,
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

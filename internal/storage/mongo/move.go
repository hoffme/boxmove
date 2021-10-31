package mongo

import (
	"errors"

	"github.com/hoffme/boxmove/pkg/storage/move"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type move_s struct {
	connection     *conn
	collectionName string
}

func newMoveStore(connection *conn, collectionName string) *move_s {
	return &move_s{
		connection:     connection,
		collectionName: collectionName,
	}
}

func (m *move_s) collection() *mongo.Collection {
	return m.connection.Collection(m.collectionName)
}

func (m *move_s) Find(id string) (*move.DTO, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	dto := &move.DTO{}

	err = m.collection().FindOne(m.connection.Ctx, bson.D{{
		Key:   "_id",
		Value: oid,
	}}).Decode(dto)

	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (m *move_s) Search(filter *move.Filter) ([]*move.DTO, error) {
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
	if len(filter.FromID) > 0 {
		bsonFilter = append(bsonFilter, bson.E{Key: "from_id", Value: filter.FromID})
	}
	if len(filter.ToID) > 0 {
		bsonFilter = append(bsonFilter, bson.E{Key: "to_id", Value: filter.ToID})
	}
	if len(filter.ItemID) > 0 {
		bsonFilter = append(bsonFilter, bson.E{Key: "item_id", Value: filter.ItemID})
	}
	if !filter.DateFrom.IsZero() || !filter.DateTo.IsZero() {
		dateFilter := bson.D{}

		if !filter.DateFrom.IsZero() {
			dateFilter = append(dateFilter, bson.E{Key: "$gte", Value: filter.DateFrom})
		}
		if !filter.DateTo.IsZero() {
			dateFilter = append(dateFilter, bson.E{Key: "$lt", Value: filter.DateTo})
		}

		bsonFilter = append(bsonFilter, bson.E{Key: "date", Value: dateFilter})
	}
	if filter.QuantityMin > 0 || filter.QuantityMax > 0 {
		quantityFilter := bson.D{}

		if filter.QuantityMin > 0 {
			quantityFilter = append(quantityFilter, bson.E{Key: "$gte", Value: filter.QuantityMin})
		}
		if filter.QuantityMax > 0 {
			quantityFilter = append(quantityFilter, bson.E{Key: "$lt", Value: filter.QuantityMax})
		}

		bsonFilter = append(bsonFilter, bson.E{Key: "quantity", Value: quantityFilter})
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

	cursor, err := m.collection().Find(m.connection.Ctx, bsonFilter, options)
	if err != nil {
		return nil, err
	}
	defer func() { _ = cursor.Close(m.connection.Ctx) }()

	result := []*move.DTO{}

	for cursor.Next(m.connection.Ctx) {
		var dto *move.DTO

		err := cursor.Decode(dto)
		if err != nil {
			return nil, err
		}

		result = append(result, dto)
	}

	return result, nil
}

func (m *move_s) Create(fields *move.Fields) (*move.DTO, error) {
	if fields == nil {
		return nil, errors.New("invalid params")
	}

	dto := &move.DTO{
		FromID:   fields.FromID,
		ToID:     fields.ToID,
		ItemID:   fields.ItemID,
		Date:     fields.Date,
		Quantity: fields.Quantity,
		Meta:     fields.Meta,
	}

	result, err := m.collection().InsertOne(m.connection.Ctx, bson.D{
		{Key: "from_id", Value: fields.FromID},
		{Key: "to_id", Value: fields.ToID},
		{Key: "item_id", Value: fields.ItemID},
		{Key: "date", Value: fields.Date},
		{Key: "quantity", Value: fields.Quantity},
		{Key: "meta", Value: fields.Meta},
	})
	if err != nil {
		return nil, err
	}

	dto.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return dto, nil
}

func (m *move_s) Update(id string, fields *move.Fields) (*move.DTO, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if fields == nil {
		return nil, errors.New("invalid fields")
	}

	dto, err := m.Find(id)
	if err != nil {
		return nil, err
	}

	result, err := m.collection().UpdateOne(
		m.connection.Ctx,
		bson.M{"_id": oid},
		bson.D{{
			Key: "$set",
			Value: bson.D{
				{Key: "from_id", Value: fields.FromID},
				{Key: "to_id", Value: fields.ToID},
				{Key: "item_id", Value: fields.ItemID},
				{Key: "date", Value: fields.Date},
				{Key: "quantity", Value: fields.Quantity},
				{Key: "meta", Value: fields.Meta},
			},
		}},
	)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("move not found")
	}

	dto.FromID = fields.FromID
	dto.ToID = fields.ToID
	dto.ItemID = fields.ItemID
	dto.Date = fields.Date
	dto.Quantity = fields.Quantity
	dto.Meta = fields.Meta

	return dto, nil
}

func (m *move_s) Delete(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := m.collection().DeleteOne(
		m.connection.Ctx,
		bson.M{"_id": oid},
	)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("move not found")
	}

	return nil
}

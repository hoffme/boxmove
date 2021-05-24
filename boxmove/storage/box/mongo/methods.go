package mongo

import (
	"errors"
	"log"
	"time"

	"github.com/hoffme/boxmove/boxmove/box"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (b *MongoStorage) collection() *mongo.Collection {
	return b.conn.DB().Collection(b.collectionName)
}

func (b *MongoStorage) FindById(id string) (box.DTO, error) {
	idP, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	dto := &DTOMongo{}
	err = b.collection().FindOne(b.ctx, bson.M{ "_id": idP, "key": b.key }).Decode(dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (b *MongoStorage) FindAll(filter *box.DTOFilterParams) ([]box.DTO, error) {
	f := bsonFilter(filter)
	f["key"] = b.key

	cursor, err := b.collection().Find(b.ctx, f)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := cursor.Close(b.ctx)
		if err != nil {
			log.Print(err)
		}
	}()

	var result []box.DTO
	for cursor.Next(b.ctx) {
		dto := &DTOMongo{}

		err := cursor.Decode(dto)
		if err != nil {
			return nil, err
		}

		result = append(result, dto)
	}

	return result, nil
}

func (b *MongoStorage) Create(fields *box.DTOCreateParams) (box.DTO, error) {
	dto := dtoMongoFromCreateParams(fields, b.key)

	result, err := b.collection().InsertOne(b.ctx, dto)
	if err != nil {
		return nil, err
	}

	dto.Id = result.InsertedID.(primitive.ObjectID)

	return dto, nil
}

func (b *MongoStorage) Update(dto box.DTO, fields *box.DTOUpdateParams) error {
	dtoMongo, ok := dto.(*DTOMongo)
	if !ok {
		return errors.New("invalid DTO")
	}

	timeUpdate := time.Now()

	result, err := b.collection().UpdateOne(b.ctx, bson.M{ "_id": dtoMongo.Id, "key": b.key }, bson.D{
		{ "$set", bsonUpdate(fields, timeUpdate) },
	})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("not found Box")
	}

	dtoMongo.UpdatedAt = timeUpdate
	if fields.Route != nil {
		dtoMongo.Route = *fields.Route
	}
	if fields.Name != nil {
		dtoMongo.Name = *fields.Name
	}

	return nil
}

func (b *MongoStorage) Delete(dto box.DTO) error {
	dtoMongo, ok := dto.(*DTOMongo)
	if !ok {
		return errors.New("invalid DTO")
	}

	timeDeleted := time.Now()

	_, err := b.collection().UpdateOne(b.ctx, bson.M{ "_id": dtoMongo.Id, "key": b.key }, bson.D{
		{ "$set", bson.D{
			{ "deleted_at", timeDeleted },
		} },
	})
	if err != nil {
		return err
	}

	dtoMongo.DeletedAt = &timeDeleted

	return nil
}

func (b *MongoStorage) Restore(dto box.DTO) error {
	dtoMongo, ok := dto.(*DTOMongo)
	if !ok {
		return errors.New("invalid DTO")
	}

	_, err := b.collection().UpdateOne(b.ctx, bson.M{ "_id": dtoMongo.Id, "key": b.key }, bson.D{
		{ "$set", bson.D{
			{ "deleted_at", nil },
		} },
	})
	if err != nil {
		return err
	}

	dtoMongo.DeletedAt = nil

	return nil
}

func (b *MongoStorage) Remove(dto box.DTO) error {
	dtoMongo, ok := dto.(*DTOMongo)
	if !ok {
		return errors.New("invalid DTO")
	}

	result, err := b.collection().DeleteOne(b.ctx, bson.M{ "_id": dtoMongo.Id, "key": b.key })
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("box not deleted or not found")
	}

	return nil
}

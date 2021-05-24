package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"github.com/hoffme/boxmove/boxmove/move"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoStorage) collection() *mongo.Collection {
	return r.conn.DB().Collection(r.collectionName)
}

func (r *MongoStorage) FindById(id string) (move.DTO, error) {
	idP, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	dto := &DTOMongo{}
	err = r.collection().FindOne(r.ctx, bson.M{ "_id": idP, "key": r.key }).Decode(dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (r *MongoStorage) FindAll(filter *move.DTOFilterParams) ([]move.DTO, error) {
	f := bsonFilter(filter)
	f["key"] = r.key

	cursor, err := r.collection().Find(r.ctx, f)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := cursor.Close(r.ctx)
		if err != nil {
			log.Print(err)
		}
	}()

	var result []move.DTO
	for cursor.Next(r.ctx) {
		dto := &DTOMongo{}

		err := cursor.Decode(dto)
		if err != nil {
			return nil, err
		}

		result = append(result, dto)
	}

	return result, nil
}

func (r *MongoStorage) Create(params *move.DTOCreateParams) (move.DTO, error) {
	dto := dtoMongoFromCreateParams(params, r.key)

	result, err := r.collection().InsertOne(r.ctx, dto)
	if err != nil {
		return nil, err
	}

	dto.Id = result.InsertedID.(primitive.ObjectID)

	return dto, nil
}

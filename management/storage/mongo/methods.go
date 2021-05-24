package mongo

import (
	"errors"
	"time"

	"github.com/hoffme/boxmove/management/client"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (b *Storage) collection() *mongo.Collection {
	return b.conn.DB().Collection(b.collectionName)
}


func (b *Storage) New(params *client.DTOCreateParams) (client.DTO, error) {
	dto := &DTOMongo{
		Name: params.Name,
		CreatedAt: params.CreatedAt,
	}

	result, err := b.collection().InsertOne(b.ctx, dto)
	if err != nil {
		return nil, err
	}

	dto.Id = result.InsertedID.(primitive.ObjectID)

	return dto, nil
}

func (b *Storage) Get(id string) (client.DTO, error) {
	idP, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	dto := &DTOMongo{}
	err = b.collection().FindOne(b.ctx, bson.M{ "_id": idP, "deleted_at": nil }).Decode(dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (b *Storage) Delete(dto client.DTO) error {
	dtoMongo, ok := dto.(*DTOMongo)
	if !ok {
		return errors.New("invalid DTO")
	}

	timeDeleted := time.Now()

	_, err := b.collection().UpdateOne(b.ctx, bson.M{ "_id": dtoMongo.Id, "deleted_at": nil }, bson.D{
		{ "$set", bson.D{
			{ "deleted_at", timeDeleted },
		} },
	})
	if err != nil {
		return err
	}

	dtoMongo.DeletedAT = &timeDeleted

	return nil
}

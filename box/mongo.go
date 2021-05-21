package box

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/hoffme/boxmove/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type dtoMongo struct {
	Id        primitive.ObjectID `bson:"_id, omitempty"`
	Route     []string           `bson:"route"`
	Name      string             `bson:"name"`
	Type      string             `bson:"type"`
	CreatedAt time.Time			 `bson:"created_at"`
	UpdatedAt time.Time			 `bson:"updated_at"`
	DeletedAt *time.Time		 `bson:"deleted_at"`
}

func (d *dtoMongo) view() *View {
	return &View{
		ID:        d.Id.Hex(),
		Route:     d.Route,
		Name:      d.Name,
		Type:      d.Type,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		DeletedAt: d.DeletedAt,
	}
}

type mongoRepository struct {
	conn           *storage.Connection
	ctx            context.Context
	databaseName   string
	collectionName string
}

func NewMongoRepository(conn *storage.Connection, databaseName, collectionName string) (Repository, error) {
	repo := &mongoRepository{
		conn: 			conn,
		ctx:            conn.Ctx,
		databaseName:   databaseName,
		collectionName: collectionName,
	}

	_, err := repo.collection().Indexes().CreateOne(repo.ctx, mongo.IndexModel{
		Keys: bson.D{ { "name", "text" } }, Options: nil,
	})
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (b *mongoRepository) collection() *mongo.Collection {
	return b.conn.Client.Database(b.databaseName).Collection(b.collectionName)
}

func (b *mongoRepository) findById(id string) (dto, error) {
	idP, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	dto := &dtoMongo{}
	err = b.collection().FindOne(b.ctx, bson.M{ "_id": idP }).Decode(dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (b *mongoRepository) findAll(filter *Filter) ([]dto, error) {
	f := bsonFilter(filter)

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

	var result []dto
	for cursor.Next(b.ctx) {
		dto := &dtoMongo{}

		err := cursor.Decode(dto)
		if err != nil {
			return nil, err
		}

		result = append(result, dto)
	}

	return result, nil
}

func (b *mongoRepository) create(fields *createParams) (dto, error) {
	dto := dtoMongoFromCreateParams(fields)

	result, err := b.collection().InsertOne(b.ctx, dto)
	if err != nil {
		return nil, err
	}

	dto.Id = result.InsertedID.(primitive.ObjectID)

	return dto, nil
}

func (b *mongoRepository) update(dto dto, fields *updateParams) error {
	dtoMongo, ok := dto.(*dtoMongo)
	if !ok {
		return errors.New("invalid dto")
	}

	timeUpdate := time.Now()

	result, err := b.collection().UpdateByID(b.ctx, dtoMongo.Id, bson.D{
		{ "$set", bsonUpdate(fields, timeUpdate) },
	})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("not found Box")
	}

	dtoMongo.update(fields, timeUpdate)

	return nil
}

func (b *mongoRepository) delete(dto dto) error {
	dtoMongo, ok := dto.(*dtoMongo)
	if !ok {
		return errors.New("invalid dto")
	}

	timeDeleted := time.Now()

	_, err := b.collection().UpdateByID(b.ctx, dtoMongo.Id, bson.D{
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

func (b *mongoRepository) restore(dto dto) error {
	dtoMongo, ok := dto.(*dtoMongo)
	if !ok {
		return errors.New("invalid dto")
	}

	_, err := b.collection().UpdateByID(b.ctx, dtoMongo.Id, bson.D{
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

func (b *mongoRepository) remove(dto dto) error {
	dtoMongo, ok := dto.(*dtoMongo)
	if !ok {
		return errors.New("invalid dto")
	}

	searchParams := bson.M{ "_id": dtoMongo.Id }

	result, err := b.collection().DeleteOne(b.ctx, searchParams)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("box not deleted or not found")
	}

	return nil
}

// utils

func dtoMongoFromCreateParams(fields *createParams) *dtoMongo {
	return &dtoMongo{
		Route:     fields.Route,
		Name:      fields.Name,
		Type:      fields.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func bsonFilter(filter *Filter) bson.M {
	f := bson.M{}

	if filter != nil {
		if filter.ID != nil && len(filter.ID) > 0 {
			var ids []primitive.ObjectID
			for _, id := range filter.ID {
				idP, err := primitive.ObjectIDFromHex(id)
				if err == nil {
					ids = append(ids, idP)
				}
			}
			f["_id"] = bson.M{ "$in": ids }
		}
		if filter.ParentID != nil && len(*filter.ParentID) > 0 {
			f["route"] = bson.E{ Key: "$elemMatch", Value: *filter.ParentID }
		}
		if filter.AncestorID != nil && len(*filter.AncestorID) > 0 {
			f["route"] = bson.E{ Key: "$elemMatch", Value: *filter.AncestorID }
		}
		if len(filter.Name) > 0 {
			f["$text"] = bson.M{ "$search": filter.Name }
		}
		if len(filter.Type) > 0 {
			f["type"] = filter.Type
		}
	}

	if filter != nil && filter.Deleted {
		f["deleted_at"] = bson.D{bson.E{ Key: "$ne", Value: nil }}
	} else {
		f["deleted_at"] = nil
	}

	return f
}

func bsonUpdate(fields *updateParams, time time.Time) bson.D {
	result := bson.D{ { Key: "updated_at", Value: time } }

	if fields.Route != nil {
		result = append(result, bson.E{ Key: "route", Value: *fields.Route })
	}
	if fields.Name != nil {
		result = append(result, bson.E{ Key: "name", Value: *fields.Name })
	}

	return result
}

func (d *dtoMongo) update(fields *updateParams, time time.Time) {
	d.UpdatedAt = time

	if fields.Route != nil {
		d.Route = *fields.Route
	}
	if fields.Name != nil {
		d.Name = *fields.Name
	}
}
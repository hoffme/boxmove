package move

import (
	"context"
	"log"
	"time"

	"github.com/hoffme/boxmove/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type dtoMongo struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	FromID    string             `bson:"from_id"`
	ToID      string             `bson:"to_id"`
	Date      time.Time  		 `bson:"date"`
	Count     uint64             `bson:"count"`
	CreatedAt time.Time          `bson:"create_at"`
}

func (d *dtoMongo) view() *View {
	return &View{
		ID: d.Id.Hex(),
		FromID: d.FromID,
		ToID: d.ToID,
		Date: d.Date,
		Count: d.Count,
		CreatedAt: d.CreatedAt,
	}
}

type MongoRepository struct {
	conn           *storage.Connection
	ctx            context.Context
	databaseName   string
	collectionName string
}

func NewMongoRepository(conn *storage.Connection, databaseName, collectionName string) (Repository, error) {
	return &MongoRepository{
		conn:           conn,
		ctx:            conn.Ctx,
		databaseName:   databaseName,
		collectionName: collectionName,
	}, nil
}

func (r *MongoRepository) collection() *mongo.Collection {
	return r.conn.Client.Database(r.databaseName).Collection(r.collectionName)
}

func (r *MongoRepository) findById(id string) (dto, error) {
	idP, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	dto := &dtoMongo{}
	err = r.collection().FindOne(r.ctx, bson.M{ "_id": idP }).Decode(dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (r *MongoRepository) findAll(filter *Filter) ([]dto, error) {
	f := bsonFilter(filter)

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

	var result []dto
	for cursor.Next(r.ctx) {
		dto := &dtoMongo{}

		err := cursor.Decode(dto)
		if err != nil {
			return nil, err
		}

		result = append(result, dto)
	}

	return result, nil
}

func (r *MongoRepository) create(params *CreateParams) (dto, error) {
	dto := dtoMongoFromCreateParams(params)

	result, err := r.collection().InsertOne(r.ctx, dto)
	if err != nil {
		return nil, err
	}

	dto.Id = result.InsertedID.(primitive.ObjectID)

	return dto, nil
}

// utils

func dtoMongoFromCreateParams(params *CreateParams) *dtoMongo {
	return &dtoMongo{
		FromID:    params.FromID,
		ToID:      params.ToID,
		Date:      params.Date,
		Count:     params.Count,
		CreatedAt: time.Now(),
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
		if filter.FromID != nil && len(filter.FromID) > 0 {
			f["from_id"] = bson.M{ "$in": filter.FromID }
		}
		if filter.ToID != nil && len(filter.ToID) > 0 {
			f["to_id"] = bson.M{ "$in": filter.ToID }
		}
		if filter.CountMin > 0 || filter.CountMax > 0 {
			countFilter := bson.D{}
			if filter.CountMin > 0 {
				countFilter = append(countFilter, bson.E{ Key: "$gt", Value: int(filter.CountMin * 100) })
			}
			if filter.CountMax > 0 {
				countFilter = append(countFilter, bson.E{ Key: "$lt", Value: int(filter.CountMax * 100) })
			}
			f["count"] = countFilter
		}
		if !filter.DateMin.IsZero() || !filter.DateMax.IsZero() {
			dateFilter := bson.D{}
			if !filter.DateMin.IsZero() {
				dateFilter = append(dateFilter, bson.E{ Key: "$gt", Value: filter.DateMin })
			}
			if !filter.DateMax.IsZero() {
				dateFilter = append(dateFilter, bson.E{ Key: "$lt", Value: filter.DateMax })
			}
			f["date"] = dateFilter
		}
	}

	return f
}

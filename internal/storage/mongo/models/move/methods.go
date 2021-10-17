package move

import (
	"errors"
	"time"

	"github.com/hoffme/boxmove/internal/app/move"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Storage) collection() *mongo.Collection {
	return s.connection.DB().Collection(s.collectionName)
}

func (s *Storage) FindOne(client, id string) (move.DTO, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	dto := &DTO{}
	err = s.collection().FindOne(
		s.ctx,
		bson.M{"_id": oid, "client": client},
	).Decode(dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (s *Storage) FindAll(client string, params *move.Filter) ([]move.DTO, error) {
	filter, err := bsonFilter(client, params)
	if err != nil {
		return nil, err
	}

	cursor, err := s.collection().Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func() { _ = cursor.Close(s.ctx) }()

	var result []move.DTO
	for cursor.Next(s.ctx) {
		dto := &DTO{}

		err := cursor.Decode(dto)
		if err != nil {
			return nil, err
		}

		result = append(result, dto)
	}

	return result, nil
}

func (s *Storage) Create(client string, params *move.CreateParams) (move.DTO, error) {
	if params == nil {
		return nil, errors.New("invalid params")
	}

	dto := &DTO{
		Client:    client,
		Active:    params.Active,
		From:      params.From,
		To:        params.To,
		Date:      params.Date,
		Count:     params.Count,
		Meta:      params.Meta,
		CreatedAt: time.Now(),
	}

	result, err := s.collection().InsertOne(s.ctx, dto)
	if err != nil {
		return nil, err
	}

	dto.OID = result.InsertedID.(primitive.ObjectID)

	return dto, nil
}

func bsonFilter(client string, params *move.Filter) (bson.M, error) {
	if params == nil {
		return nil, errors.New("invalid filter")
	}
	if len(client) == 0 {
		return nil, errors.New("invalid client")
	}

	filter := bson.M{"client": client}

	if params.ID != nil && len(params.ID) > 0 {
		var ids []primitive.ObjectID
		for _, id := range params.ID {
			idP, err := primitive.ObjectIDFromHex(id)
			if err == nil {
				ids = append(ids, idP)
			}
		}
		filter["_id"] = bson.M{"$in": ids}
	}
	if len(params.Active) > 0 {
		filter["active"] = params.Active
	}
	if params.FromID != nil && len(params.FromID) > 0 {
		filter["from_id"] = bson.M{"$in": params.FromID}
	}
	if params.ToID != nil && len(params.ToID) > 0 {
		filter["to_id"] = bson.M{"$in": params.ToID}
	}
	if params.CountMin > 0 || params.CountMax > 0 {
		countFilter := bson.D{}
		if params.CountMin > 0 {
			countFilter = append(countFilter, bson.E{Key: "$gt", Value: params.CountMin})
		}
		if params.CountMax > 0 {
			countFilter = append(countFilter, bson.E{Key: "$lt", Value: params.CountMax})
		}
		filter["count"] = countFilter
	}
	if !params.DateMin.IsZero() || !params.DateMax.IsZero() {
		dateFilter := bson.D{}
		if !params.DateMin.IsZero() {
			dateFilter = append(dateFilter, bson.E{Key: "$gt", Value: params.DateMin})
		}
		if !params.DateMax.IsZero() {
			dateFilter = append(dateFilter, bson.E{Key: "$lt", Value: params.DateMax})
		}
		filter["date"] = dateFilter
	}

	return filter, nil
}

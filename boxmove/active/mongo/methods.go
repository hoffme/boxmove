package mongo

import (
	"errors"
	"time"

	"github.com/hoffme/boxmove/boxmove/active"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Storage) collection() *mongo.Collection {
	return s.connection.DB().Collection(s.collectionName)
}

func (s *Storage) FindOne(client, id string) (active.DTO, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	dto := &DTO{}
	err = s.collection().FindOne(
		s.ctx,
		bson.M{ "_id": oid, "client": client },
	).Decode(dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (s *Storage) FindAll(client string, params *active.Filter) ([]active.DTO, error) {
	if params == nil {
		return nil, errors.New("invalid filter")
	}

	filter, err := bsonFilter(client, params)
	if err != nil {
		return nil, err
	}

	cursor, err := s.collection().Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func() { _ = cursor.Close(s.ctx) }()

	var result []active.DTO
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

func (s *Storage) Create(client string, params *active.CreateParams) (active.DTO, error) {
	if params == nil {
		return nil, errors.New("invalid params")
	}

	now := time.Now()

	dto := &DTO{
		Client: client,
		Code: params.Code,
		Name: params.Name,
		Meta: params.Meta,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := s.collection().InsertOne(s.ctx, dto)
	if err != nil {
		return nil, err
	}

	dto.OID = result.InsertedID.(primitive.ObjectID)

	return dto, nil
}

func (s *Storage) Update(dtoI active.DTO, params *active.UpdateParams) error {
	dto, ok := dtoI.(*DTO)
	if !ok {
		return errors.New("invalid DTO")
	}

	timeUpdate := time.Now()
	update, err := bsonUpdate(params, timeUpdate)
	if err != nil {
		return err
	}

	result, err := s.collection().UpdateOne(
		s.ctx,
		bson.M{ "_id": dto.OID, "client": dto.Client },
		update,
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("not found")
	}

	dto.UpdatedAt = timeUpdate
	if params.Name != nil {
		dto.Name = *params.Name
	}
	if params.Meta != nil {
		dto.Meta = *params.Meta
	}

	return nil
}

func (s *Storage) Delete(dtoInterface active.DTO) error {
	dto, ok := dtoInterface.(*DTO)
	if !ok {
		return errors.New("invalid DTO")
	}

	timeDeleted := time.Now()

	result, err := s.collection().UpdateOne(
		s.ctx,
		bson.M{ "_id": dto.OID, "client": dto.Client },
		bson.D{{ "$set", bson.D{{ "deleted_at", timeDeleted }}}},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("not found")
	}

	dto.DeletedAt = &timeDeleted

	return nil
}

func (s *Storage) Restore(dtoInterface active.DTO) error {
	dto, ok := dtoInterface.(*DTO)
	if !ok {
		return errors.New("invalid DTO")
	}

	result, err := s.collection().UpdateOne(
		s.ctx,
		bson.M{ "_id": dto.OID, "client": dto.Client },
		bson.D{{ "$set", bson.D{{ "deleted_at", nil }}}},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("not found")
	}

	dto.DeletedAt = nil

	return nil
}

func (s *Storage) Remove(dtoInterface active.DTO) error {
	dto, ok := dtoInterface.(*DTO)
	if !ok {
		return errors.New("invalid DTO")
	}

	result, err := s.collection().DeleteOne(
		s.ctx,
		bson.M{ "_id": dto.OID, "client": dto.Client },
	)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("not found")
	}

	return nil
}

func bsonUpdate(params *active.UpdateParams, date time.Time) (bson.D, error) {
	fields := bson.D{{ Key: "updated_at", Value: date } }

	if params == nil || params.Name == nil || params.Meta == nil {
		return nil, errors.New("empty fields update")
	}

	if params.Name != nil {
		fields = append(fields, bson.E{ Key: "name", Value: *params.Name })
	}
	if params.Meta != nil {
		fields = append(fields, bson.E{ Key: "meta", Value: *params.Meta })
	}

	return bson.D{{ "$set", params }}, nil
}

func bsonFilter(client string, params *active.Filter) (bson.M, error) {
	filter := bson.M{ "client": client }

	if params.ID != nil && len(params.ID) > 0 {
		var ids []primitive.ObjectID
		for _, id := range params.ID {
			idP, err := primitive.ObjectIDFromHex(id)
			if err == nil {
				ids = append(ids, idP)
			}
		}
		filter["_id"] = bson.M{ "$in": ids }
	}
	if len(params.Name) > 0 {
		filter["name"] = bson.M{ "$search": params.Name }
	}
	if params.Codes != nil && len(params.Codes) > 0 {
		filter["code"] = bson.M{ "$in": params.Codes }
	}

	if params.Deleted {
		filter["deleted_at"] = bson.D{bson.E{ Key: "$ne", Value: nil }}
	} else {
		filter["deleted_at"] = nil
	}

	return filter, nil
}

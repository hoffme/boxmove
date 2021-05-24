package mongo

import (
	"time"

	"github.com/hoffme/boxmove/boxmove/box"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func dtoMongoFromCreateParams(fields *box.DTOCreateParams, key string) *DTOMongo {
	return &DTOMongo{
		Route:     fields.Route,
		Name:      fields.Name,
		Type:      fields.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Key:       key,
	}
}

func bsonFilter(filter *box.DTOFilterParams) bson.M {
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

func bsonUpdate(fields *box.DTOUpdateParams, time time.Time) bson.D {
	result := bson.D{ { Key: "updated_at", Value: time } }

	if fields.Route != nil {
		result = append(result, bson.E{ Key: "route", Value: *fields.Route })
	}
	if fields.Name != nil {
		result = append(result, bson.E{ Key: "name", Value: *fields.Name })
	}

	return result
}

package mongo

import (
	"time"

	"github.com/hoffme/boxmove/boxmove/move"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func dtoMongoFromCreateParams(params *move.DTOCreateParams, key string) *DTOMongo {
	return &DTOMongo{
		FromID:    params.FromID,
		ToID:      params.ToID,
		Date:      params.Date,
		Count:     params.Count,
		CreatedAt: time.Now(),
		Key: 	   key,
	}
}

func bsonFilter(filter *move.DTOFilterParams) bson.M {
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

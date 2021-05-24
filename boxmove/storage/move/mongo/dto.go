package mongo

import (
	"time"

	"github.com/hoffme/boxmove/boxmove/move"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DTOMongo struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	FromID    string             `bson:"from_id"`
	ToID      string             `bson:"to_id"`
	Date      time.Time  		 `bson:"date"`
	Count     uint64             `bson:"count"`
	CreatedAt time.Time          `bson:"create_at"`
	Key 	  string			 `bson:"key"`
}

func (d *DTOMongo) View() *move.View {
	return &move.View{
		ID: d.Id.Hex(),
		FromID: d.FromID,
		ToID: d.ToID,
		Date: d.Date,
		Count: d.Count,
		CreatedAt: d.CreatedAt,
	}
}

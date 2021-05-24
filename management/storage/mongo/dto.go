package mongo

import (
	"github.com/hoffme/boxmove/management/client"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DTOMongo struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	CreatedAt time.Time			 `bson:"created_at"`
	DeletedAT *time.Time 		 `bson:"deleted_at"`
}

func (d *DTOMongo) View() *client.View {
	return &client.View{
		ID:        d.Id.Hex(),
		Name:      d.Name,
		CreatedAt: d.CreatedAt,
		DeletedAt: d.DeletedAT,
	}
}

package mongo

import (
	"time"

	"github.com/hoffme/boxmove/boxmove/box"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DTOMongo struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Route     []string           `bson:"route"`
	Name      string             `bson:"name"`
	Type      string             `bson:"type"`
	CreatedAt time.Time			 `bson:"created_at"`
	UpdatedAt time.Time			 `bson:"updated_at"`
	DeletedAt *time.Time		 `bson:"deleted_at"`
	Key 	  string			 `bson:"key"`
}

func (d *DTOMongo) View() *box.View {
	return &box.View{
		ID:        d.Id.Hex(),
		Route:     d.Route,
		Name:      d.Name,
		Type:      d.Type,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		DeletedAt: d.DeletedAt,
	}
}

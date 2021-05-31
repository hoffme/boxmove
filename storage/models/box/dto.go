package box

import (
	"github.com/hoffme/boxmove/boxmove/box"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DTO struct {
	OID       primitive.ObjectID `bson:"_id,omitempty"`
	Client 	  string 			 `bson:"client"`
	Route 	  []string   		 `bson:"route"`
	Name      string     		 `bson:"name"`
	Actives   map[string]int64	 `bson:"actives"`
	Meta 	  interface{} 		 `bson:"meta"`
	CreatedAt time.Time  		 `bson:"created_at"`
	UpdatedAt time.Time  		 `bson:"updated_at"`
	DeletedAt *time.Time 		 `bson:"deleted_at,omitempty"`
}

func (dto *DTO) ID() string {
	return dto.OID.Hex()
}

func (dto *DTO) View() *box.View {
	return &box.View{
		ID: dto.ID(),
		Name: dto.Name,
		Meta: dto.Meta,
		Route: dto.Route,
		Actives: dto.Actives,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		DeletedAt: dto.DeletedAt,
	}
}
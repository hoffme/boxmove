package active

import (
	"time"

	"github.com/hoffme/boxmove/boxmove/active"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DTO struct {
	OID       primitive.ObjectID `bson:"_id,omitempty"`
	Client 	  string 			 `bson:"client"`
	Code 	  string             `bson:"code"`
	Name      string     		 `bson:"name"`
	Meta 	  interface{} 		 `bson:"meta"`
	CreatedAt time.Time  		 `bson:"created_at"`
	UpdatedAt time.Time  		 `bson:"updated_at"`
	DeletedAt *time.Time 		 `bson:"deleted_at,omitempty"`
}

func (dto *DTO) ID() string {
	return dto.OID.Hex()
}

func (dto *DTO) View() *active.View {
	return &active.View{
		ID: dto.ID(),
		Code: dto.Code,
		Name: dto.Name,
		Meta: dto.Meta,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		DeletedAt: dto.DeletedAt,
	}
}
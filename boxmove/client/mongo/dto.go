package mongo

import (
	"github.com/hoffme/boxmove/boxmove/client"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DTO struct {
	OID       primitive.ObjectID `bson:"_id,omitempty"`
	Name      string     		 `bson:"name"`
	CreatedAt time.Time  		 `bson:"created_at"`
	UpdatedAt time.Time  		 `bson:"updated_at"`
	DeletedAt *time.Time 		 `bson:"deleted_at,omitempty"`
}

func (dto *DTO) ID() string {
	return dto.OID.Hex()
}

func (dto *DTO) View() *client.View {
	return &client.View{
		ID: dto.ID(),
		Name: dto.Name,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		DeletedAt: dto.DeletedAt,
	}
}
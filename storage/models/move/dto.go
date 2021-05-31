package move

import (
	"github.com/hoffme/boxmove/boxmove/move"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DTO struct {
	OID       primitive.ObjectID `bson:"_id,omitempty"`
	Client 	  string 			 `bson:"client"`
	Active 	  string	  		 `bson:"active"`
	From   	  string 	  		 `bson:"from"`
	To     	  string	  		 `bson:"to"`
	Date   	  time.Time   		 `bson:"date"`
	Count  	  uint64  	  		 `bson:"count"`
	Meta 	  interface{} 		 `bson:"meta"`
	CreatedAt time.Time  		 `bson:"created_at"`
}

func (dto *DTO) ID() string {
	return dto.OID.Hex()
}

func (dto *DTO) View() *move.View {
	return &move.View{
		ID: dto.ID(),
		Active: dto.Active,
		From: dto.From,
		To: dto.To,
		Date: dto.Date,
		Count: dto.Count,
		Meta: dto.Meta,
		CreatedAt: dto.CreatedAt,
	}
}
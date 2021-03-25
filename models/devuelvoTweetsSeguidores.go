package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:_"id,omitempty"`
	UserId            string             `bson:"userid" json:"userId,omitempty"`
	UsuarioRelacionID string             `bson:"usuarioRelacionID" json:"usuarioRelacionID,omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}

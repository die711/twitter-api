package bd

import (
	"api/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func ConsultaRelacion(relacion models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         relacion.UsuarioId,
		"usuariorelacionid": relacion.UsuarioRelacionID,
	}

	var resultado models.Relacion
	err := collection.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil

}

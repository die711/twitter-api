package bd

import (
	"api/models"
	"context"
	"time"
)

func InsertoRelacion(relacion models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("relacion")

	_, err := collection.InsertOne(ctx, relacion)
	if err != nil {
		return false, err
	}

	return true, nil

}

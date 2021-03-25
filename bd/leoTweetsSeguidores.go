package bd

import (
	"api/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func LeoTweerSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("relacion")
	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":       "tweet",
			"localField": "usuariorelacionid",
			"foreignField": "userid",
			"as":         "tweet",
		},
	})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{
		"$sort": bson.M{"tweet.fecha": -1},
	})
	condiciones = append(condiciones, bson.M{
		"$skip": skip,
	})

	condiciones = append(condiciones, bson.M{
		"$limit": 20,
	})
	var result []models.DevuelvoTweetsSeguidores
	cursor, err := collection.Aggregate(ctx, condiciones)

	if err != nil {
		fmt.Println(err)
		return result, false
	}

	err = cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}
	return result, true
}

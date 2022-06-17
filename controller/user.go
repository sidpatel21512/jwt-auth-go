package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sidpatel21512/jwt-auth-go/collection"
	"github.com/sidpatel21512/jwt-auth-go/model"
	"go.mongodb.org/mongo-driver/bson"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	users, err := getUsers()

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(users)
}

func getUsers() ([]model.User, error) {
	curser, err := collection.UserCollection.Find(context.Background(), bson.D{{}})

	if err != nil {
		return nil, err
	}

	defer curser.Close(context.Background())

	var users []model.User

	for curser.Next(context.Background()) {
		var user model.User

		if err := curser.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

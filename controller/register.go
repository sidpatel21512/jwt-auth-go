package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sidpatel21512/jwt-auth-go/collection"
	"github.com/sidpatel21512/jwt-auth-go/model"
	"github.com/sidpatel21512/jwt-auth-go/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var newUser model.User

	json.NewDecoder(r.Body).Decode(&newUser)

	hashedPassword, _ := utils.HashPassword(newUser.Password)
	newUser.Password = string(hashedPassword)
	id := registerUser(newUser)

	fmt.Println("User added:", id)
	json.NewEncoder(w).Encode(id)
}

func registerUser(user model.User) primitive.ObjectID {
	inserted, err := collection.UserCollection.InsertOne(context.Background(), user, nil)

	if err != nil {
		log.Fatal(err)
	}

	return inserted.InsertedID.(primitive.ObjectID)
}

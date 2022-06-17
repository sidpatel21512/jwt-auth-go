package collection

import (
	"fmt"

	"github.com/sidpatel21512/jwt-auth-go/connection"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "users"

var UserCollection *mongo.Collection

func init() {
	UserCollection = connection.Database.Collection(collectionName)

	fmt.Println("users collection is ready")
}

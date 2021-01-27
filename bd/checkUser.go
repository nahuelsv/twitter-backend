package bd

import (
	"context"
	"time"

	"github.com/nahuelsv/twitter-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckUser function that check if the user exist*/
func CheckUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("twitter")
	col := db.Collection("user")
	cond := bson.M{"email": email}

	var result models.User
	err := col.FindOne(ctx, cond).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}

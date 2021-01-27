package bd

import (
	"context"
	"time"

	"github.com/nahuelsv/twitter-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertRegister is the function to insert the user into DB*/
func InsertRegister(user models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("twitter")
	col := db.Collection("user")

	user.Password, _ = PasswordEncrypt(user.Password)

	result, err := col.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

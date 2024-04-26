package poc

import (
	"context"
	"fmt"
	"pokemon-game-api/pkgs/di"
	stringutils "pokemon-game-api/pkgs/utils/string"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func init() {
// 	userCollection = database.NewMongoDbClient().Client.Database("learn").Collection("users")
// }

func CreateUser(c *gin.Context) {
	request := new(User)
	c.BindJSON(request)

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	userCollection := di.MongoDb.Client.Database("learn").Collection("users")
	r, _ := userCollection.InsertOne(ctx, request)

	c.JSON(200, r.InsertedID)
}

func GetUser(c *gin.Context) {
	name := c.Query("name")

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	// filters := bson.E{Key: "first_name", Value: name}
	filters := bson.D{
		{Key: "info.age", Value: bson.D{
			{Key: "$gte", Value: 15},
		}},
	}

	if !stringutils.IsNilOrEmpty(&name) {
		// filters = append(filters, bson.E{Key: "first_name", Value: bson.D{
		// 	{Key: "$eq", Value: name},
		// }})

		// short cut of syntax above
		filters = append(filters, bson.E{Key: "first_name", Value: name})
	}

	options := options.Find().SetSort(bson.D{
		{Key: "info.age", Value: -1},  // -1 is DESC ordering
		{Key: "first_name", Value: 1}, // 1 is ASC ordering
	})

	userCollection := di.MongoDb.Client.Database("learn").Collection("users")
	res, _ := userCollection.Find(ctx, filters, options)

	r := new([]User)
	res.All(ctx, r)

	for _, result := range *r {
		rr, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(rr))
	}

	c.JSON(200, r)
}

package poc

import (
	"context"
	"fmt"
	"net/http"
	"pokemon-game-api/pkgs/database"
	stringutils "pokemon-game-api/pkgs/utils/string"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	isOption = true

	userCollection = func() *mongo.Collection {
		if isOption {
			return NewMongoWithOption().Database("db-option").Collection("users")
		} else {
			// return database.NewMongoDbClient("mongodb://admin:password@localhost:27017/").Client.Database("learn").Collection("users")
			return database.NewMongoDbClient("mongodb://admin:password@localhost:27017/").Client.Database("db-option").Collection("users")
		}
	}()

	pokemonCollection = func() *mongo.Collection {
		if isOption {
			return NewMongoWithOption().Database("db-option").Collection("pokemons")
		} else {
			// return database.NewMongoDbClient("mongodb://admin:password@localhost:27017/").Client.Database("learn").Collection("pokemons")
			return database.NewMongoDbClient("mongodb://admin:password@localhost:27017/").Client.Database("db-option").Collection("pokemons")
		}
	}()
)

func CreateUser(c *gin.Context) {
	request := new(User)
	c.BindJSON(request)

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	r, _ := userCollection.InsertOne(ctx, request)

	c.JSON(200, r.InsertedID)
}

func GetUser(c *gin.Context) {
	name := c.Query("name")

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	// filters := bson.E{Key: "first_name", Value: name}

	// filter can not be nil
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

	res, _ := userCollection.Find(ctx, filters, options)

	r := new([]User)
	res.All(ctx, r)

	for _, result := range *r {
		rr, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(rr))
	}

	c.JSON(200, r)
}

func GetUserWithPokemon(c *gin.Context) {
	name := c.Query("name")

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	// filters := bson.E{Key: "first_name", Value: name}

	// filter can not be nil
	filters := bson.D{
		{Key: "info.age", Value: bson.D{
			{Key: "$gte", Value: 10},
		}},
	}

	if !stringutils.IsNilOrEmpty(&name) {
		// filters = append(filters, bson.E{Key: "first_name", Value: bson.D{
		// 	{Key: "$eq", Value: name},
		// }})

		// short cut of syntax above
		filters = append(filters, bson.E{Key: "first_name", Value: name})
	}

	pipeline := bson.A{
		bson.M{ // where in noSQL
			"$match": filters,
		},
		bson.M{
			"$lookup": bson.M{ // join in noSQL
				"from":         "pokemons",
				"localField":   "pokemon_ids",
				"foreignField": "_id",
				"as":           "pokemons_lookup",
				// lookup field technicaly it doesn't need to create property, It auto create for us in any type
				// BUT if I receive it by *Struct you still need to create prop manually
			},
		},
		bson.M{ // set sort in aggregate
			"$sort": bson.M{
				"info.age":   -1, // Sorting by age in descending order
				"first_name": 1,  // Secondary sorting by first name in ascending order
			},
		},
	}

	res, _ := userCollection.Aggregate(ctx, pipeline)

	r := new([]User)
	res.All(ctx, r)

	for _, result := range *r {
		rr, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(rr))
	}

	c.JSON(200, r)
}

func CreatePokemon(c *gin.Context) {
	request := new(Pokemon)
	c.BindJSON(request)

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	r, _ := pokemonCollection.InsertOne(ctx, request)

	c.JSON(200, r.InsertedID)
}

func GetPokemon(c *gin.Context) {
	name := c.Query("name")

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	// filter can not be nil
	filters := bson.D{}

	if !stringutils.IsNilOrEmpty(&name) {
		filters = bson.D{{Key: "name", Value: name}}
	}

	pipeline := bson.A{
		bson.M{ // where in noSQL
			"$match": filters,
		},
		bson.M{
			"$lookup": bson.M{ // join in noSQL
				"from":         "users",
				"localField":   "user",
				"foreignField": "_id",
				"as":           "user_data",
			},
		},
	}

	res, err := pokemonCollection.Aggregate(ctx, pipeline)

	if err != nil {
		c.AbortWithStatusJSON(500, err)
		return
	}

	r := make([]Pokemon, 0, 2)
	res.All(ctx, r)

	for _, result := range r {
		rr, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(rr))
	}

	c.JSON(200, r)
}

func UpdateUserPokemon(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	request := new([]primitive.ObjectID)
	c.BindJSON(request)

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{
		"pokemons": request,
	}}

	r, _ := userCollection.UpdateOne(ctx, filter, update)

	c.JSON(200, r.ModifiedCount)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	request := new(User)
	c.BindJSON(request)

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	filter := bson.M{"_id": objId}

	r, _ := userCollection.ReplaceOne(ctx, filter, request)

	c.JSON(200, r.ModifiedCount)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	filter := bson.M{"_id": objId}

	r, _ := userCollection.DeleteOne(ctx, filter)

	c.JSON(200, r.DeletedCount)
}

func DropUserCollection(c *gin.Context) {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	err := userCollection.Drop(ctx)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.Status(http.StatusOK)
}

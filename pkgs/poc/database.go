package poc

import (
	"context"
	customlog "pokemon-game-api/pkgs/logs"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once   sync.Once
	client *mongo.Client
)

func NewMongoWithOption() *mongo.Client {
	once.Do(func() {
		// Use the SetServerAPIOptions() method to set the Stable API version to 1
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		bsonOpts := &options.BSONOptions{
			UseJSONStructTags: true,
			NilSliceAsEmpty:   true,
		}

		opts := options.
			Client().
			ApplyURI("mongodb://admin:password@localhost:27017/").
			SetServerAPIOptions(serverAPI).
			SetBSONOptions(bsonOpts)

		// go driver use context to set timeout for this task only
		ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancle()

		// Create a new client and connect to the server
		mClient, err := mongo.Connect(ctx, opts)

		if err != nil {
			customlog.WriteMongoClientPanicLog(err.Error())
		}

		client = mClient
	})

	return client
}

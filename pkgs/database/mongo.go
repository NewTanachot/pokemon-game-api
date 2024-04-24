package database

import (
	"context"
	"fmt"
	"pokemon-game-api/pkgs/config"
	customlog "pokemon-game-api/pkgs/logs"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once   sync.Once
	client *MongoDb
)

type MongoDb struct {
	Client           *mongo.Client
	ConnectionString string
}

func NewMongoDbClient() *MongoDb {
	once.Do(func() {
		connectionString := fmt.Sprintf("mongodb://%v:%v@%v:%v/",
			config.MongoUser,
			config.MongoPassword,
			config.MongoHost,
			config.MongoPort)

		// Use the SetServerAPIOptions() method to set the Stable API version to 1
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.
			Client().
			ApplyURI(connectionString).
			SetServerAPIOptions(serverAPI)

		// go driver use context to set timeout for each excutetion task of mongo client
		ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancle()

		// Create a new client and connect to the server
		mClient, err := mongo.Connect(ctx, opts)

		if err != nil {
			customlog.WriteMongoClientPanicLog()
		}

		client = &MongoDb{
			Client:           mClient,
			ConnectionString: connectionString,
		}
	})

	if isSuccess := client.PingMongoDb(); !isSuccess {
		customlog.WriteMongoClientPanicLog()
	}

	return client
}

func (m *MongoDb) PingMongoDb() bool {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	if err := m.Client.Ping(ctx, nil); err != nil {
		return false
	}

	return true
}

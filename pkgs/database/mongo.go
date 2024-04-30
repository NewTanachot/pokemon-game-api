package database

import (
	"context"
	"fmt"
	"pokemon-game-api/pkgs/config"
	customerror "pokemon-game-api/pkgs/error"
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
	DbClient         *mongo.Database
	ConnectionString string
}

func NewMongoDbClient(connectionString ...string) *MongoDb {
	once.Do(func() {
		var cStr string

		if len(connectionString) == 0 {
			cStr = fmt.Sprintf("mongodb://%s:%s@%s:%s/", *config.MongoUser, *config.MongoPassword, *config.MongoHost, *config.MongoPort)
		} else {
			cStr = connectionString[0]
		}

		// Use the SetServerAPIOptions() method to set the Stable API version to 1
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.
			Client().
			ApplyURI(cStr).
			SetServerAPIOptions(serverAPI)

		// go driver use context to set timeout for this task only
		ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancle()

		// Create a new client and connect to the server
		mClient, err := mongo.Connect(ctx, opts)

		if err != nil {
			customlog.WriteMongoClientPanicLog(err.Error())
		}

		client = &MongoDb{
			Client:           mClient,
			DbClient:         mClient.Database(*config.MongoDbName),
			ConnectionString: cStr,
		}
	})

	if isSuccess := client.PingMongoDb(); !isSuccess {
		customlog.WriteMongoClientPanicLog(customerror.PingDbFail)
	}

	return client
}

func (m *MongoDb) CloseMongoDb() {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	m.Client.Disconnect(ctx)
}

func (m *MongoDb) PingMongoDb() bool {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	if err := m.Client.Ping(ctx, nil); err != nil {
		return false
	}

	return true
}

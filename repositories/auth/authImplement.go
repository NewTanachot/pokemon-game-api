package authrepo

import (
	"context"
	"fmt"
	"net/http"
	"pokemon-game-api/domains/entities"
	"pokemon-game-api/pkgs/constants"
	customerror "pokemon-game-api/pkgs/error"
	stringutils "pokemon-game-api/pkgs/utils/string"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	AuthDbCollection *mongo.Collection
	ExecutionTimeout time.Duration
}

func NewAuthGatway(db *mongo.Database) IAuthRepository {
	return AuthRepository{
		AuthDbCollection: db.Collection(constants.AuthColl),
		ExecutionTimeout: 5 * time.Second,
	}
}

func (a AuthRepository) CreateUser(user *entities.User) (*mongo.InsertOneResult, error) {
	ctx, cancle := context.WithTimeout(context.Background(), a.ExecutionTimeout)
	defer cancle()

	result, err := a.AuthDbCollection.InsertOne(ctx, user)

	if err != nil {
		return nil, customerror.NewCustomError(constants.AuthRepo,
			http.StatusBadRequest, fmt.Sprintf("%v_user", customerror.UnableToCreate))
	}

	return result, nil
}

func (a AuthRepository) ReadAllUser() (*[]entities.User, error) {
	ctx, cancle := context.WithTimeout(context.Background(), a.ExecutionTimeout)
	defer cancle()

	cursor, err := a.AuthDbCollection.Find(ctx, bson.D{})

	// TODO
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	result := []entities.User{}
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (a AuthRepository) ReadUserById(id string) (*entities.User, error) {
	ctx, cancle := context.WithTimeout(context.Background(), a.ExecutionTimeout)
	defer cancle()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	result := new(entities.User)

	err = a.AuthDbCollection.FindOne(ctx, filter).Decode(result)
	// TODO: mongo.ErrNoDocuments is error when not found
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a AuthRepository) UpdateUserById(id string, user *entities.User) (string, error) {
	return stringutils.Empty, nil
}

func (a AuthRepository) DeleteUserById(id string) (string, error) {
	return stringutils.Empty, nil
}

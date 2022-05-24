package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/the-bogdan/go-rest-api/app/internal/users"
	"github.com/the-bogdan/go-rest-api/app/pkg/logging"
)

type db struct {
	collection *mongo.Collection
	logger     logging.Logger
}

func NewStorage(database *mongo.Database, collection string, logger logging.Logger) users.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}

func (d *db) Create(ctx context.Context, user *users.User) (string, error) {
	d.logger.Debug("creating user")

	// insert user into mongodb
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user in mongodb. err %v", err)
	}

	d.logger.Debug("convert InsertedID to ObjectID")

	// convert result oid into hex string
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	return "", fmt.Errorf("failed to convert objectid to hex. oid : %s", oid)
}

func (d *db) FindOne(ctx context.Context, id string) (*users.User, error) {
	// convert user id into ObjectID obj
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert hex to objectid. hex: %s", id)
	}

	// create bson filter for mongo
	filter := bson.M{
		"_id": oid,
	}

	// search
	result := d.collection.FindOne(ctx, filter)
	options.FindOptions{}

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			// TODO ErrEntityNotFound
		}
		return nil, fmt.Errorf("failed to find user by id: %s due to error: %v", id, err)
	}

	// decode mongo response into golang User struct
	var u users.User
	if err = result.Decode(&u); err != nil {
		return nil, fmt.Errorf("failed to decode user (id:%s) due to error: %v", id, err)
	}
	return &u, nil
}

func (d *db) Update(ctx context.Context, user *users.User) error {
	// convert user id into ObjectID obj
	oid, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. hex: %s", user.Id)
	}

	// create bson filter for mongo
	filter := bson.M{
		"_id": oid,
	}

	// convert user model ingo bson object
	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to bson.Marshal user due to error: %v", err)
	}
	var updateUserObj bson.M
	if err = bson.Unmarshal(userBytes, &updateUserObj); err != nil {
		return fmt.Errorf("failed to bson.Unmarshal user bytes due to error: %v", err)
	}

	// TODO do we need thid check this
	//delete(updateUserObj, "_id")

	// make final bson query
	query := bson.M{
		"$set": updateUserObj,
	}

	// TODO check if i can update user without casting it to bson model
	// update user
	result, err := d.collection.UpdateOne(ctx, filter, query)

	if err != nil {
		// TODO check if not found creates error
		return fmt.Errorf("failed to execute update user query due to error: %v", err)
	}
	if result.ModifiedCount == 0 {
		// TODO ErrEntityNotFound
		return fmt.Errorf("failed to find user by id: %s", user.Id)
	}
	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. hex: %s", user.Id)
	}
	panic("implement me")
}

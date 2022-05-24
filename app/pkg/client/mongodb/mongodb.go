package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDb string) (*mongo.Database, error) {
	mongoDBURL := fmt.Sprintf("mongodb://%s:%s", host, port)

	credentials := options.Credential{
		Username: username,
		Password: password,
	}
	if authDb == "" {
		credentials.AuthSource = authDb
	}

	clientOpts := options.Client().ApplyURI(mongoDBURL).SetAuth(credentials)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("connection to mongoDB failed. err: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping mongodb failed. err: %v", err)
	}

	return client.Database(database), nil
}

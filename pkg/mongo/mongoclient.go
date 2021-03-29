package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type IMongoClient interface {
	RunCommand(cmd string) (bson.M, error)
	Close() error
}

type mongoClient struct {
	client mongo.Client
	ctx    context.Context
}

func NewMongoClient(addr string) (IMongoClient, error) {
	mc := &mongoClient{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(addr))
	if err != nil {
		return *mc, err
	}
	mc.client = *client
	mc.ctx = ctx
	return *mc, nil
}

func (m mongoClient) RunCommand(cmd string) (bson.M, error) {
	var db *mongo.Database
	db = m.client.Database("admin")
	command := bson.E{"isMaster", true}
	opts := options.RunCmd()
	var result bson.M
	if err := db.RunCommand(context.TODO(), command, opts).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (m mongoClient) Close() error {
	err := m.client.Disconnect(m.ctx)
	if err != nil {
		return err
	}
	return nil
}

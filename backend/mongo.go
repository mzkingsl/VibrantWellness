package main

import (
  "context"
  "os"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context) (*mongo.Client, error) {
  uri := os.Getenv("MONGO_URI") // e.g. "mongodb://mongo:27017/statesdb"
  return mongo.Connect(ctx, options.Client().ApplyURI(uri))
}

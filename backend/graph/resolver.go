package graph

import "go.mongodb.org/mongo-driver/mongo"

// Resolver holds your dependencies for injection into resolvers.
type Resolver struct {
    MongoClient *mongo.Client
}

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.72

import (
	"context"

	"example.com/states-typeahead/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

// States returns all states whose name or abbreviation matches `filter` (case-insensitive).
func (r *queryResolver) States(ctx context.Context, filter string) ([]*model.State, error) {
	col := r.MongoClient.
		Database("statesdb").
		Collection("states")

	// match name or abbreviation via a case-insensitive regex
	q := bson.M{"$or": []bson.M{
		{"name": bson.M{"$regex": filter, "$options": "i"}},
		{"abbreviation": bson.M{"$regex": filter, "$options": "i"}},
	}}

	cursor, err := col.Find(ctx, q)
	if err != nil {
		return nil, err
	}
	var results []*model.State
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

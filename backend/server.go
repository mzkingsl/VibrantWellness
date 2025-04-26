package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"example.com/states-typeahead/graph"
	"example.com/states-typeahead/graph/model"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultPort = "8080"

func main() {
	// setting up the port
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// create a context with 10 seconds before timing out
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoURI := os.Getenv("MONGO_URI") // e.g. "mongodb://mongo:27017/statesdb"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("cannot connect to mongo:", err)
	}

	// seed the states in our db
	seedStates(ctx, client)

	// connecting mongo to graphql resolver
	resolver := &graph.Resolver{MongoClient: client}

	// building graphql server, 
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](100))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// mitigating cors issues 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// setting CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.URL.Path == "/" {
			playground.Handler("GraphQL playground", "/query").ServeHTTP(w, r)
		} else if r.URL.Path == "/query" {
			srv.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// seedStates inserts all 50 states + territories if empty
func seedStates(ctx context.Context, client *mongo.Client) {
	col := client.Database("statesdb").Collection("states")
	count, err := col.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatal("count error:", err)
	}
	if count > 0 {
		return
	}
	// a simple in-memory list 
	states := []interface{}{
		model.State{Name: "Alabama", Abbreviation: "AL"},
		model.State{Name: "Alaska", Abbreviation: "AK"},
		model.State{Name: "Arizona", Abbreviation: "AZ"},
		model.State{Name: "Arkansas", Abbreviation: "AR"},
		model.State{Name: "California", Abbreviation: "CA"},
		model.State{Name: "Colorado", Abbreviation: "CO"},
		model.State{Name: "Connecticut", Abbreviation: "CT"},
		model.State{Name: "Delaware", Abbreviation: "DE"},
		model.State{Name: "Florida", Abbreviation: "FL"},
		model.State{Name: "Georgia", Abbreviation: "GA"},
		model.State{Name: "Hawaii", Abbreviation: "HI"},
		model.State{Name: "Idaho", Abbreviation: "ID"},
		model.State{Name: "Illinois", Abbreviation: "IL"},
		model.State{Name: "Indiana", Abbreviation: "IN"},
		model.State{Name: "Iowa", Abbreviation: "IA"},
		model.State{Name: "Kansas", Abbreviation: "KS"},
		model.State{Name: "Kentucky", Abbreviation: "KY"},
		model.State{Name: "Louisiana", Abbreviation: "LA"},
		model.State{Name: "Maine", Abbreviation: "ME"},
		model.State{Name: "Maryland", Abbreviation: "MD"},
		model.State{Name: "Massachusetts", Abbreviation: "MA"},
		model.State{Name: "Michigan", Abbreviation: "MI"},
		model.State{Name: "Minnesota", Abbreviation: "MN"},
		model.State{Name: "Mississippi", Abbreviation: "MS"},
		model.State{Name: "Missouri", Abbreviation: "MO"},
		model.State{Name: "Montana", Abbreviation: "MT"},
		model.State{Name: "Nebraska", Abbreviation: "NE"},
		model.State{Name: "Nevada", Abbreviation: "NV"},
		model.State{Name: "New Hampshire", Abbreviation: "NH"},
		model.State{Name: "New Jersey", Abbreviation: "NJ"},
		model.State{Name: "New Mexico", Abbreviation: "NM"},
		model.State{Name: "New York", Abbreviation: "NY"},
		model.State{Name: "North Carolina", Abbreviation: "NC"},
		model.State{Name: "North Dakota", Abbreviation: "ND"},
		model.State{Name: "Ohio", Abbreviation: "OH"},
		model.State{Name: "Oklahoma", Abbreviation: "OK"},
		model.State{Name: "Oregon", Abbreviation: "OR"},
		model.State{Name: "Pennsylvania", Abbreviation: "PA"},
		model.State{Name: "Rhode Island", Abbreviation: "RI"},
		model.State{Name: "South Carolina", Abbreviation: "SC"},
		model.State{Name: "South Dakota", Abbreviation: "SD"},
		model.State{Name: "Tennessee", Abbreviation: "TN"},
		model.State{Name: "Texas", Abbreviation: "TX"},
		model.State{Name: "Utah", Abbreviation: "UT"},
		model.State{Name: "Vermont", Abbreviation: "VT"},
		model.State{Name: "Virginia", Abbreviation: "VA"},
		model.State{Name: "Washington", Abbreviation: "WA"},
		model.State{Name: "West Virginia", Abbreviation: "WV"},
		model.State{Name: "Wisconsin", Abbreviation: "WI"},
		model.State{Name: "Wyoming", Abbreviation: "WY"},
	}
	_, err = col.InsertMany(ctx, states)
	if err != nil {
		log.Fatal("seed error:", err)
	}
}

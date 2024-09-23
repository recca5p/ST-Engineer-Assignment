package main

import (
	"database/sql"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"
)

type Beer struct {
	ID            int     `json:"id"`
	Price         string  `json:"price"`
	Name          string  `json:"name"`
	AverageRating float64 `json:"average_rating"`
	Reviews       int     `json:"reviews"`
	Image         string  `json:"image"`
}

var beerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Beer",
	Fields: graphql.Fields{
		"price": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"rating": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "Rating",
				Fields: graphql.Fields{
					"average": &graphql.Field{
						Type: graphql.Float,
					},
					"reviews": &graphql.Field{
						Type: graphql.Int,
					},
				},
			}),
		},
		"image": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"beer": &graphql.Field{
			Type: beerType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(int)
				beer := getBeerByID(id)
				return beer, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})

// Database connection
var db *sql.DB

func init() {
	var err error
	connStr := "user=koffee dbname=mockdata password=koffee host=103.166.182.81 port=5432 sslmode=disable" // Adjust this string
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}

func getBeerByID(id int) *Beer {
	beer := &Beer{}
	row := db.QueryRow("SELECT id, price, name, average_rating, reviews, image FROM beers WHERE id = $1", id)
	err := row.Scan(&beer.ID, &beer.Price, &beer.Name, &beer.AverageRating, &beer.Reviews, &beer.Image)
	if err != nil {
		return nil
	}
	return beer
}

func main() {
	http.Handle("/graphql", handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}))
	http.ListenAndServe(":8080", nil)
}

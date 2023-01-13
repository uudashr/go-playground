package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func runMain() error {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return err
	}

	// Query
	query := `
		{
			hello
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if r.HasErrors() {
		return fmt.Errorf("failed to execute graphql operation, errors: %+v", r.Errors)
	}

	rJSON, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", rJSON)
	return nil
}

func main() {
	if err := runMain(); err != nil {
		log.Println("runMain error:", err)
	}
	log.Println("Done")
}

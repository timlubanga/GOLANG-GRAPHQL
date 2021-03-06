package main

import (
	"fmt"
	"gqlgen-todos/graph"
	"gqlgen-todos/graph/generated"
	"gqlgen-todos/graph/postgres"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := postgres.InitDB()
	if err != nil {
		fmt.Printf("this is the error:%v", err)

	}

	defer postgres.Close()
	c := generated.Config{
		Resolvers: &graph.Resolver{TodosRepo: postgres.TodosRepo{DB: db},
			UserRepo: postgres.UserRepo{DB: db}},
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ipreferwater/graphql-theory/config"
	"github.com/ipreferwater/graphql-theory/db"
	"github.com/ipreferwater/graphql-theory/graph"
	"github.com/ipreferwater/graphql-theory/graph/generated"
	"github.com/ipreferwater/graphql-theory/logs"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.InitConfig()
	logs.InitLogs()
	db.InitTODORepo()
	db.InitMysqlPostRepository()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}

package main

import (
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/ipreferwater/graphql-theory/auth"
	"github.com/ipreferwater/graphql-theory/config"
	"github.com/ipreferwater/graphql-theory/db"
	"github.com/ipreferwater/graphql-theory/graph"
	"github.com/ipreferwater/graphql-theory/graph/generated"
	"github.com/ipreferwater/graphql-theory/logs"
	log "github.com/sirupsen/logrus"
)

func main() {

	arg := os.Args[1]
	if arg == "local" {
		config.SetEnvLocal()
	}

	config.InitConfig()

	logs.InitLogs()
	db.InitTODORepo()
	db.InitMysqlPostRepository()

	r := gin.Default()
	r.Use(auth.AuthMiddleware())
	
	r.GET("/sandbox", playgroundHandler())
	r.POST("/query", graphqlHandler())
	r.POST("/login", auth.Login)

	r.Run(":8000")
	log.Info("graphql ready")
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

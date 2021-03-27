package main

import (
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
	config.InitConfig()
	logs.InitLogs()
	db.InitTODORepo()
	db.InitMysqlPostRepository()

	//srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	r := gin.Default()
	r.Use(auth.AuthMiddleware())
	//authorized := r.Group("/")
	//authorized.Use(auth.AuthMiddleware())
	//{
		r.POST("/query", graphqlHandler())
		r.GET("/", playgroundHandler())
	//}

	r.Run(":8000")
	log.Info("graphql ready")

	/*http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))*/
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

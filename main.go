package main

import (
	"log"
	"playfolio-backend/graph"
	resolver "playfolio-backend/graph/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func graphqlHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := playground.Handler("GraphQL playground", "/query")
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	app := gin.Default()

	app.Any("/query", graphqlHandler())
	app.GET("/", playgroundHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "3000")
	app.Run(":" + "3000")
}

package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/SimranBhagwandasani/EventPlanner/database"
	"github.com/SimranBhagwandasani/EventPlanner/routes"
	"github.com/SimranBhagwandasani/EventPlanner/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)
var eventCollection *mongo.collection = database.OpenCollection(database.Client,"food")
func main() {
	fmt.Print("Welcome to the Event planner.")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.EventRoutes(router)
	routes.CatalogRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":"+port)

}

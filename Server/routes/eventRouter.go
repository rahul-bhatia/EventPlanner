package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/SimranBhagwandasani/EventPlanner/controllers"
)


func EventRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.GET("/events",controller.GetEvents())
	incomingRoutes.GET("/events/:event_id",controller.GetEvent())
	incomingRoutes.POST("/events",controller.CreateEvent())
	incomingRoutes.PATCH("/events/:event_id",controller.UpdateeEvent())

}
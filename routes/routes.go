package routes

import (
	"github.com/aungsannphyo/go-restapi/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// event
	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:eventId", updateEvents)
	authenticated.DELETE("/events/:eventId", deleteEvent)

	// authentication
	server.POST("/signup", signup)
	server.POST("/login", login)
}

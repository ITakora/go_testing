package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {

	authenticated := route.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("events/:id", updateEvent)
	authenticated.DELETE("events/:id", deleteEvent)

	route.GET("/events", getEvents)
	route.GET("/events/:id", getEvent)
	route.POST("/signup", signUp)
	route.POST("/login", login)
}

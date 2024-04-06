package routes

import "github.com/gin-gonic/gin"

func Route(route *gin.Engine) {
	route.GET("/events", getEvents)
	route.GET("/events/:id", getEvent)
	route.POST("/events", createEvents)
	route.PUT("events/:id", updateEvent)
	route.DELETE("events/:id", deleteEvent)
	route.POST("/signup", signUp)
	route.POST("/login", login)
}

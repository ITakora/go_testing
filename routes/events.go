package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvent(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the id"})
	}

	event, err := models.GetEventById(eventID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the id"})
	}

	ctx.JSON(http.StatusOK, event)

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not connect to database :("})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Errors"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not connect to database :("})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

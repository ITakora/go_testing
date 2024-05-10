package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func createRegisterEvent(context *gin.Context) {
	userid := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the id"})
		return
	}

	err = event.Register(userid)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register users "})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered"})

}

func deleteRegisterEvent(context *gin.Context) {

}

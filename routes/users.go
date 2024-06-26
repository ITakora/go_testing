package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Errors"})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the users"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created succesfully"})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateTokens(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generating token."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successfully", "token": token})
}

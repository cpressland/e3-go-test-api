package controllers

import (
	"net/http"

	"github.com/cpressland/e3-go-test-api/contracts"
	"github.com/cpressland/e3-go-test-api/models"

	"github.com/gin-gonic/gin"
)

func (c *controller) Create(ctx *gin.Context) {
	var request contracts.CreateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := c.userRepository.CheckUsername(request.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	users, err := c.userRepository.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//get highest id from users

	id := 0
	for _, user := range users {
		if user.ID > id {
			id = user.ID
		}
	}

	user := models.User{
		ID:        id + 1,
		Username:  request.Username,
		Email:     request.Email,
		Telephone: request.Telephone,
	}

	if exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
		return
	}

	if err := c.userRepository.Create(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := contracts.CreateUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Telephone: user.Telephone,
	}

	ctx.JSON(http.StatusCreated, response)

}

package controllers

import (
	"go-test-api/contracts"
	"go-test-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *controller) Update(ctx *gin.Context) {
	var request contracts.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//get id from path
	id, _ := ctx.Params.Get("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := c.userRepository.CheckUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	user := models.User{
		ID:        userID,
		Username:  request.Username,
		Email:     request.Email,
		Telephone: request.Telephone,
	}
	err = c.userRepository.Update(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := contracts.UpdateUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Telephone: user.Telephone,
	}

	ctx.JSON(http.StatusOK, response)
}

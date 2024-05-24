package controllers

import (
	"go-test-api/contracts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *controller) List(ctx *gin.Context) {
	users, err := c.userRepository.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []contracts.GetUserResponse = []contracts.GetUserResponse{}
	for _, user := range users {
		response = append(response, contracts.GetUserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Telephone: user.Telephone,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

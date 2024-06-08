package controllers

import (
	"net/http"
	"strconv"

	"github.com/cpressland/e3-go-test-api/contracts"

	"github.com/gin-gonic/gin"
)

func (c *controller) Get(ctx *gin.Context) {
	userID := ctx.Param("id")
	id, err := strconv.Atoi(userID)

	user, err := c.userRepository.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response contracts.GetUserResponse
	response.ID = user.ID
	response.Username = user.Username
	response.Email = user.Email
	response.Telephone = user.Telephone

	ctx.JSON(http.StatusOK, response)
}

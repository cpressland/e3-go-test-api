package controllers

import (
	"go-test-api/contracts"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *controller) Delete(ctx *gin.Context) {
	userID := ctx.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := c.userRepository.CheckUserID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	err = c.userRepository.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response contracts.DeleteUserResponse
	response.ID = id
	response.Message = "user deleted"
	response.Deleted = true

	ctx.JSON(http.StatusOK, response)
}

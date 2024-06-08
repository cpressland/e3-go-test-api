package main

import (
	"github.com/cpressland/e3-go-test-api/controllers"
	"github.com/cpressland/e3-go-test-api/repositories"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func main() {
	//log api start
	log.Info("API started")

	//setup gin server
	router := gin.New()

	//setup repositories
	userRepository := repositories.NewUserRepository()

	//setup conrollers
	_ = controllers.NewController(router, userRepository)

	//run server
	router.Run(":8080")
}

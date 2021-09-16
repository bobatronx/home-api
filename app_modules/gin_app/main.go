package main

import (
	"github.com/bobatronx/gin-app/v2/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	controllers.InitializeControllers(router)
	router.Run(":9000")
}

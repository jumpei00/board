package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/infrastructure"
	"github.com/jumpei00/board/backend/app/interfaces"
)

func main() {
	router := gin.Default()

	group := router.Group("/api")

	// DB
	threadDB := infrastructure.NewThreadDB()

	// application
	threadApp := application.NewThreadApplication(threadDB)

	// handler
	threadHandler := interfaces.NewThreadHandler(threadApp)

	// router setup
	threadHandler.SetupRouter(group)

	router.Run(":8080")
}

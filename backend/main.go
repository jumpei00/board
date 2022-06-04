package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/infrastructure"
	"github.com/jumpei00/board/backend/app/interfaces"
)

func main() {
	router := gin.Default()

	threadGroup := router.Group("/api/thread")
	commentGroup := router.Group("/api/comment")

	// DB
	threadDB := infrastructure.NewThreadDB()
	commentDB := infrastructure.NewCommentDB()

	// application
	threadApp := application.NewThreadApplication(threadDB)
	commentApp := application.NewCommentApplication(threadDB, commentDB)

	// handler
	threadHandler := interfaces.NewThreadHandler(threadApp)
	commentHandler := interfaces.NewCommentHandler(threadApp, commentApp)

	// router setup
	threadHandler.SetupRouter(threadGroup)
	commentHandler.SetupRouter(commentGroup)

	log.Fatal(router.Run(":8080"))
}

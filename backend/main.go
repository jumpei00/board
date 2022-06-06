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

	apigroup := router.Group("/api")
	threadGroup := router.Group("/api/thread")
	commentGroup := router.Group("/api/comment")

	// DB
	userDB := infrastructure.NewUserDB()
	threadDB := infrastructure.NewThreadDB()
	commentDB := infrastructure.NewCommentDB()

	// application
	userApp := application.NewUserApplication(userDB)
	threadApp := application.NewThreadApplication(threadDB)
	commentApp := application.NewCommentApplication(threadDB, commentDB)

	// handler
	userHandler := interfaces.NewUserHandler(userApp)
	threadHandler := interfaces.NewThreadHandler(threadApp)
	commentHandler := interfaces.NewCommentHandler(threadApp, commentApp)

	// router setup
	userHandler.SetupRouter(apigroup)
	threadHandler.SetupRouter(threadGroup)
	commentHandler.SetupRouter(commentGroup)

	log.Fatal(router.Run(":8080"))
}

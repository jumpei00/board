package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/config"
	"github.com/jumpei00/board/backend/app/infrastructure"
	"github.com/jumpei00/board/backend/app/interfaces"
	"github.com/jumpei00/board/backend/app/library/logger"
	_ "github.com/jumpei00/board/backend/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Board API
// @version 1.0
// @description This is api server of board application
// @host localhost.api
// license.name jumpei00
// @BasePath /
func main() {
	// DB
	dbPool, err := infrastructure.GenerateDBPool()
	if err != nil {
		logger.Fatal("db session open error", "error", err)
	}

	userDB := infrastructure.NewUserDB(dbPool)
	visitorDB := infrastructure.NewVisitorDB(dbPool)
	threadDB := infrastructure.NewThreadRepository(dbPool)
	commentDB := infrastructure.NewCommentDB(dbPool)

	// application
	userApp := application.NewUserApplication(userDB)
	visitApp := application.NewVisitorApplication(visitorDB)
	threadApp := application.NewThreadApplication(threadDB, commentDB)
	commentApp := application.NewCommentApplication(threadDB, commentDB)

	// handler
	userHandler := interfaces.NewUserHandler(userApp)
	visitorHandler := interfaces.NewVisitorsHandler(visitApp)
	threadHandler := interfaces.NewThreadHandler(threadApp)
	commentHandler := interfaces.NewCommentHandler(threadApp, commentApp)

	// router setup
	router := gin.Default()

	apigroup := router.Group("/api")
	visitorGroup := router.Group("/api/visitor")
	threadGroup := router.Group("/api/thread")
	commentGroup := router.Group("/api/comment")

	userHandler.SetupRouter(apigroup)
	visitorHandler.SetupRouter(visitorGroup)
	threadHandler.SetupRouter(threadGroup)
	commentHandler.SetupRouter(commentGroup)

	// swaggerの設定
	if config.IsDevelopment() {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	if err := router.Run(":8080"); err != nil {
		logger.Fatal("API server run failed", "error", err)
	}
}

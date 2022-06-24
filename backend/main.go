package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"github.com/gin-contrib/sessions"
	ginRedis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/config"
	"github.com/jumpei00/board/backend/app/infrastructure"
	"github.com/jumpei00/board/backend/app/interfaces"
	"github.com/jumpei00/board/backend/app/interfaces/session"
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
	// mysql
	dbPool, err := infrastructure.GenerateDBPool()
	if err != nil {
		logger.Fatal("db session open error", "error", err)
	}

	// redis
	redisPool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.GetRedisHost())
			if err != nil {
				logger.Fatal("redis session open error", "error", err)
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	redisStore, err := ginRedis.NewStoreWithPool(redisPool, []byte(config.GetSessionSecret()))
	if err != nil {
		logger.Fatal("redis store open error", "error", err)
	}

	if err := ginRedis.SetKeyPrefix(redisStore, "user:"); err != nil {
		logger.Fatal("redis store key prefix set error", "error", err)
	}

	// session
	sessionMiddleware := sessions.Sessions(config.SessionName, redisStore)

	// security
	secureConfig := secure.Config{
		SSLRedirect: true,
		STSSeconds:  315360000,
		STSIncludeSubdomains: true,
		FrameDeny: true,
		ContentTypeNosniff: true,
		BrowserXssFilter: true,
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		IsDevelopment: config.IsDevelopment(),
	}
	secureMiddleware := secure.New(secureConfig)

	// cross origin
	crossOriginConfig := cors.Config{
		AllowOrigins: []string{config.GetFrontURL()},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}
	crossOriginMiddleware := cors.New(crossOriginConfig)

	// repository
	session := session.NewSessionManager()
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
	userHandler := interfaces.NewUserHandler(session, userApp)
	visitorHandler := interfaces.NewVisitorsHandler(visitApp)
	threadHandler := interfaces.NewThreadHandler(session, threadApp)
	commentHandler := interfaces.NewCommentHandler(session, threadApp, commentApp)

	// router setup
	router := gin.Default()

	// middleware
	router.Use(sessionMiddleware, secureMiddleware, crossOriginMiddleware)

	// routing setup
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

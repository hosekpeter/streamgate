package api

import (
	"net/http"
	"streamgate/internal/api/middleware"
	"streamgate/internal/dependencies"

	"github.com/gin-gonic/gin"
)

func New(deps dependencies.Service) http.Handler {
	cfg := deps.Config()

	gin.SetMode(cfg.GinMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestLogger(deps.Logger()))

	stream := router.Group("/stream", cors)
	stream.GET("/:stream/file/:file", upstream(deps))

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return router
}

func cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}

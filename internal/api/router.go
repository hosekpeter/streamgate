package api

import (
	"net/http"
	"streamgate/internal/dependencies"

	"github.com/gin-gonic/gin"
)

func New(deps dependencies.Service) http.Handler {
	cfg := deps.Config()

	gin.SetMode(cfg.GinMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	stream := router.Group("/stream", cors)
	stream.GET("/:stream/file/:file", upstream(deps))

	router.GET("/", func(c *gin.Context) {
		c.JSONP(http.StatusOK, "Hello, World!")
	})

	return router
}

func cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}

package api

import (
	"net/http"
	"streamgate/internal/dependencies"
	"streamgate/internal/httpclient"

	"github.com/gin-gonic/gin"
)

func upstream(deps dependencies.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		stream := c.Param("stream")
		file := c.Param("file")

		resp, err := deps.HTTPClient().MakeRequest(c.Request.Context(), httpclient.Request{
			Method:  http.MethodGet,
			BaseURL: deps.Config().MediatedHLSURL,
			Path:    stream + "/" + file,
			Query:   c.Request.URL.Query(),
		})
		if err != nil {
			deps.Logger().ErrorContext(c.Request.Context(), err.Error())
			c.AbortWithError(http.StatusBadGateway, err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			deps.Logger().InfoContext(c.Request.Context(), "upstream request completed", "body", resp.Body)
			c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
			return
		}

		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	}
}

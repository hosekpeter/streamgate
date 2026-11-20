package middleware

import (
	"log/slog"
	"streamgate/internal/logging"
	"time"
	"uuid"

	"github.com/gin-gonic/gin"
)

func RequestLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := c.GetHeader("X-Request-ID")

		if requestID == "" || len(requestID) > 64 {
			requestID = uuid.NewV7().String()
		}

		ctx := logging.WithRequestID(c.Request.Context(), requestID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		status := c.Writer.Status()

		defer func() {
			end := time.Now()
			elapsed := end.Sub(start)

			level := slog.LevelInfo
			if status >= 500 {
				level = slog.LevelError
			} else if status >= 400 {
				level = slog.LevelWarn
			}

			logger.Log(ctx,
				level,
				"Request",
				"method", c.Request.Method,
				"status", status,
				"path", c.Request.URL.Path,
				"duration", elapsed.String(),
			)
		}()
	}
}

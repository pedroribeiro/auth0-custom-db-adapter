package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoggerMiddleware struct {
	UTC       bool
	SkipPaths []string
}

// Use provides implementation for middleware
func (m LoggerMiddleware) UseLogger(r *gin.Engine) {
	// Skip paths
	skip := make(map[string]bool)

	// Convert list of path to map
	for _, p := range m.SkipPaths {
		skip[p] = true
	}

	middleware := func(c *gin.Context) {
		// At receive request, define some vars
		start := time.Now()
		path := c.Request.URL.Path
		url := c.Request.URL.RequestURI()

		// Move forward middlewares and wait finish execution
		c.Next()

		// Verify if should log this request
		track := true

		if _, ok := skip[path]; ok {
			track = false
		}

		if _, ok := skip[url]; ok {
			track = false
		}

		if track {
			// Prepare to log
			end := time.Now()
			latency := end.Sub(start)

			// Convert time to UTC
			if m.UTC {
				end = end.UTC()
			}

			msg := "Request"
			if len(c.Errors) > 0 {
				msg = c.Errors.String()
			}

			log := logrus.WithFields(logrus.Fields{
				"module":  "http",
				"method":  c.Request.Method,
				"path":    url,
				"status":  c.Writer.Status(),
				"latency": latency,
				"ip":      c.ClientIP(),
			})

			switch {
			case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
				log.Warn(msg)

			case c.Writer.Status() >= http.StatusInternalServerError:
				log.Error(msg)

			default:
				log.Info(msg)
			}
		}
	}

	// Install middleware
	r.Use(middleware)
}

// NewLogger return instance of LoggerMiddleware
func NewLogger(utc bool, skipPaths []string) *LoggerMiddleware {
	return &LoggerMiddleware{
		UTC:       utc,
		SkipPaths: skipPaths,
	}
}

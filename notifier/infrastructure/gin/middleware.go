package gin

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func MiddlewarePing(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet && strings.HasSuffix(c.Request.URL.Path, path) {
			c.String(http.StatusOK, "pong")
			c.Abort()
			return
		}
		c.Next()
	}
}

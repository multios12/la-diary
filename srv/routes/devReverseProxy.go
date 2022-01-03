package routes

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// 開発環境向けリバースプロキシ
func devReverseProxy(c *gin.Context) {
	remote, _ := url.Parse("http://localhost:8080")
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(c.Writer, c.Request)
}

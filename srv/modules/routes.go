package routes

import (
	"embed"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// gin.Engineインスタンスにルーティングを設定して返す
func Initial(proxyMode bool, static embed.FS) *gin.Engine {
	router := gin.Default()

	var f = func(c *gin.Context) {
		if proxyMode {
			remote, _ := url.Parse("http://localhost:8080")
			proxy := httputil.NewSingleHostReverseProxy(remote)
			proxy.ServeHTTP(c.Writer, c.Request)
		} else {
			c.FileFromFS("static"+c.Request.URL.Path, http.FS(static))
		}
	}

	router.GET("/", f)
	router.GET("/favicon.ico", f)
	router.GET("/css/:file", f)
	router.GET("/js/:file", f)

	router.GET("/api/:year/:month", getMonth)
	router.GET("/api/:year/:month/:day", getDay)
	router.POST("/api/:year/:month/:day", postDay)
	return router
}

func getDay(c *gin.Context) {
	day := c.Param("year") + "-" + c.Param("month") + "-" + c.Param("day")
	l := ReadLine(day)
	if len(l.Day) == 0 {
		c.Status(404)
		return
	}

	c.JSON(200, l)
}

func getMonth(c *gin.Context) {
	day := c.Param("year") + c.Param("month")
	l := ReadMonthFile(day)
	if len(l) == 0 {
		c.Status(404)
		return
	}

	c.JSON(200, l)
}

func postDay(c *gin.Context) {
	day := c.Param("year") + "-" + c.Param("month") + "-" + c.Param("day")

	var line lineModel
	if err := c.ShouldBindJSON(&line); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	WriteLine(day, line.Memo)
	c.Status(200)
}

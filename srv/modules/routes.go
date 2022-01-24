package routes

import (
	"embed"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// gin.Engineインスタンスにルーティングを設定して返す
func Initial(static embed.FS) *gin.Engine {
	router := gin.Default()

	var f = func(c *gin.Context) {
		c.FileFromFS("static"+c.Request.URL.Path, http.FS(static))
	}

	router.GET("/", f)
	router.GET("/favicon.ico", f)
	router.GET("/static/:dir/:file", f)

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
	line.Memo = strings.TrimSpace(line.Memo)
	if line.Memo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "memo is not found."})
	}

	WriteLine(day, line.Memo)
	c.Status(200)
}

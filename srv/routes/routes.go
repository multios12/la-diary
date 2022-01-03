package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin.Engineインスタンスにルーティングを設定して返す
func Initial(proxyMode bool) *gin.Engine {
	router := gin.Default()

	if proxyMode {
		router.GET("/", devReverseProxy)
		router.GET("/js/:file", devReverseProxy)
	} else {

	}

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
	l := ReadLineFile(day)
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

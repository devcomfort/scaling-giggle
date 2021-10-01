package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html") // 해당 경로의 *.html 파일에 해당하는 모든 파일을 메모리에 올려둡니다.
	router.Static("/static", "static")      // 정적 파일 정보를 모두 메모리에 올려둡니다.

	/* 메인 페이지 */
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	/* 설명서 */
	// router.GET("/explanation", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "description.html", nil)
	// })

	router.Run(":" + port)
}

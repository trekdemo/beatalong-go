package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/trekdemo/beatalong-go/music_url_parser"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/assets", "assets")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/j", func(c *gin.Context) {
		musicURL := c.Query("u")
		resourceInfo, _ := musicURLParser.Parse(musicURL)

		result := map[string]interface{}{
			"url":  musicURL,
			"info": resourceInfo,
		}
		c.JSON(http.StatusOK, result)
	})

	router.Run(":" + port)
}

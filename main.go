package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	engine := gin.Default()
	engine.POST("/upload", func(c *gin.Context) {
		file, handler, err := c.Request.FormFile("image")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad Request")
			return
		}

		fileName := handler.Filename
		dir, _ := os.Getwd()
		out, err := os.Create(dir + "\\images\\" + fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	engine.Run(":3000")
}

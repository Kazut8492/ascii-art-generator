package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/common-nighthawk/go-figure"

	"bufio"
	"io/ioutil"
	"os"
)

func main() {

	postAscii := func(c *gin.Context) {
		var ascii string
		if err := c.BindJSON(&ascii); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			// Create a text file
			f, err := os.Create("result.txt")
			if err != nil {
				fmt.Println(err.Error())
			}
			defer f.Close()

			// Generate ascii and write it into the file
			myFigure := figure.NewFigure(ascii, "", true)
			datawriter := bufio.NewWriter(f)
			for _, data := range myFigure.Slicify() {
				_, _ = datawriter.WriteString(data + "\n")
			}
			datawriter.Flush()

			// Print a content of result.txt to the console
			content, err := ioutil.ReadFile("result.txt")
			if err != nil {
				fmt.Println(err.Error())
			}
			// fmt.Println(string(content))

			// Return a content of result.txt to the frontend
			c.JSON(http.StatusOK, string(content))
		}
	}

	// Server with gin
	r := gin.Default()
	r.Use(CORSMiddleware())
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 		"content": string(content),
	// 	})
	// })
	r.POST("/ascii", postAscii)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

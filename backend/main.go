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
	// Create a text file
	f, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	// Printing ascii art
	myFigure := figure.NewFigure("Hello hello", "", true)
	myFigure.Print()

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
	fmt.Println(string(content))

	// TODO: Send a string to the frontend??

	// Server with gin
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

package main

import (
	"fmt"
	"io"
	"os"

	index "github.com/milijan-mosic/code-compete/api/controllers"
	lib "github.com/milijan-mosic/code-compete/api/lib"

	gin "github.com/gin-gonic/gin"
)

// GLOBAL VARIABLE
var env map[string]string

func main() {
	gin.SetMode(gin.ReleaseMode)
	api := gin.Default()
	URL_PREFIX := "/api/1.0"

	logFile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile)

	api.ForwardedByClientIP = true
	api.SetTrustedProxies(nil)
	api.Use(gin.Logger())
	api.Use(gin.Recovery())

	env, _ = lib.LoadEnvVariables(".env")

	test := api.Group(URL_PREFIX)
	{
		test.GET("/", index.IndexHandler)
		test.POST("/hello", index.HelloHandler)
	}

	fmt.Println("Server listening on 4500...")
	api.Run(":4500")
}

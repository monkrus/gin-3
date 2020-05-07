package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/monkrus/gin-2/middlewares"
	"github.com/monkrus/gin/controller"
	"github.com/monkrus/gin/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// changing default to new
func main() {

	setupLogOutput()
	server := gin.New()
	// customizing the middleware
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	//server.Use(gin.Logger())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}

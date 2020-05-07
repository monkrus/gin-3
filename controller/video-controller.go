package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/monkrus/gin/entity"
	"github.com/monkrus/gin/service"
)

type VideoController interface {
	FindAll() []entity.Video 
	Save(ctx *gin.Context) error
}

// implement that interface
type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return service.New().FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video 
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	c.service.Save(video) 
	return nil
}

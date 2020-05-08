package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/monkrus/gin-3/validators"
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

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
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
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

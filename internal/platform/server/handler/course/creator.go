package course

import (
	"github.com/dasalgadoc/clean-architecture-go/internal/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type courseDto struct {
	Name string `json:"name" binding:"required"`
}

type PostCourseCreator struct {
	useCase application.CourseCreator
}

func (cc *PostCourseCreator) Do(c *gin.Context) {
	var (
		err     error
		request courseDto
	)
	if err = c.BindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err = cc.useCase.Invoke(c, request.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "OK",
	})
}

func NewPostCourseCreator(creator application.CourseCreator) PostCourseCreator {
	return PostCourseCreator{
		useCase: creator,
	}
}

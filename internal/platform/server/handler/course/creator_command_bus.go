package course

import (
	"github.com/dasalgadoc/clean-architecture-go/internal/domain"
	"github.com/dasalgadoc/clean-architecture-go/shared/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostCourseCreatorAsync struct {
	commandBus command.CommandBus
}

func (cc *PostCourseCreatorAsync) Do(c *gin.Context) {
	var request courseDto
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	err := cc.commandBus.Dispatch(c, domain.NewCourseCommand(request.Name))

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

func NewPostCourseCreatorAsync(commandBus command.CommandBus) PostCourseCreatorAsync {
	return PostCourseCreatorAsync{
		commandBus: commandBus,
	}
}

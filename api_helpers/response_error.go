package api_helpers

import (
	"fmt"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Error       string
	Description string
}

func RespondError(c *gin.Context, status int, custom_description string) {
	fmt.Println("status:", status)
	var response Error

	response.Error = utils.StatusMessage(status)

	if custom_description == "default" {
		response.Description = utils.StatusDescription(status)
	} else {
		response.Description = custom_description
	}

	c.JSON(status, response)
}

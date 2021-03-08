package api_helpers

import (
	"fmt"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Error       string
	Description string
}

func RespondError(c *gin.Context, status int, custom_description string) {
	fmt.Println("status:", status)
	var res ResponseError

	res.Error = utils.StatusMessage(status)

	if custom_description == "default" {
		res.Description = utils.StatusDescription(status)
	} else {
		res.Description = custom_description
	}

	c.JSON(status, res)
}

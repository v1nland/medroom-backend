package api_helpers

import (
	"fmt"
	"medroom-backend/app/Utils"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Error       string
	Description string
}

func RespondError(c *gin.Context, status int, custom_description string) {
	fmt.Println("status:", status)
	var res ResponseError

	res.Error = Utils.StatusMessage(status)

	if custom_description == "default" {
		res.Description = Utils.StatusDescription(status)
	} else {
		res.Description = custom_description
	}

	c.JSON(status, res)
}

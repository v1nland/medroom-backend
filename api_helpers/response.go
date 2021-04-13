package api_helpers

import (
	"fmt"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

type Json struct {
	Status bool        `json:"status"`
	Meta   string      `json:"meta"`
	Data   interface{} `json:"data"`
}

func RespondJson(c *gin.Context, status int, payload interface{}) {
	fmt.Println("status ", status)

	var res Json

	res.Status = (status == 200)
	res.Meta = utils.StatusMessage(status)
	res.Data = payload

	c.JSON(200, res)
}

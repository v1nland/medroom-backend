package api_helpers

import (
	"fmt"
	"medroom-backend/app/Utils"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status bool        `json:"status"`
	Meta   string      `json:"meta"`
	Data   interface{} `json:"data"`
}

func RespondJSON(c *gin.Context, status int, payload interface{}) {
	fmt.Println("status ", status)

	var res Response

	res.Status = (status == 200)
	res.Meta = Utils.StatusMessage(status)
	res.Data = payload

	c.JSON(200, res)
}

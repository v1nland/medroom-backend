package ApiHelpers

import (
	"fmt"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status bool        `json:"status"`
	Meta   string      `json:"meta"`
	Data   interface{} `json:"data"`
}

func RespondJSON(w *gin.Context, status int, payload interface{}) {
	fmt.Println("status ", status)

	var res ResponseData

	res.Status = (status == 200)
	res.Meta = Utils.StatusMessage(status)
	res.Data = payload

	w.JSON(200, res)
}

type ResponseError struct {
	Error       string
	Description string
}

func RespondError(w *gin.Context, status int, custom_description string) {
	fmt.Println("status:", status)
	var res ResponseError

	res.Error = Utils.StatusMessage(status)

	if custom_description == "default" {
		res.Description = Utils.StatusDescription(status)
	} else {
		res.Description = custom_description
	}

	w.JSON(status, res)
}

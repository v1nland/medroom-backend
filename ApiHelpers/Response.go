package ApiHelpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"medroom-backend/Utils"
)

/*
*
*
*
*	RESPONSE: JSON
*
*
*
 */
type ResponseData struct {
	Status bool
	Meta   string
	Data   interface{}
}

func RespondJSON(w *gin.Context, status int, payload interface{}) {
	// http status code
	fmt.Println("status ", status)

	// response container
	var res ResponseData

	// set values
	res.Status = (status == 200)
	res.Meta = Utils.StatusMessage(status)
	res.Data = payload

	// response
	w.JSON(200, res)
}

/*
*
*
*
*	RESPONSE: ERROR
*
*
*
 */
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

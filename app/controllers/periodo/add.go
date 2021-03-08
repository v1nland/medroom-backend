package periodo

import (
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/formats/f_input"
	"medroom-backend/app/formats/f_output"
	"medroom-backend/app/messages/Request"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Agrega un nuevo periodo
// @Description Genera un nuevo periodo con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_periodo     body    Request.AddNewPeriodo     true        "Periodo a agregar"
// @Success 200 {object} Swagger.AddNewPeriodoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/periodos [post]
func AddNewPeriodo(c *gin.Context) {
	var input Request.AddNewPeriodo
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.AddNewPeriodo(&input)

	periodo := models.Periodo{
		Nombre_periodo: *input.Nombre_periodo,
	}

	if err := repositories.AddNewPeriodo(&periodo); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, f_output.AddNewPeriodo(periodo))
}

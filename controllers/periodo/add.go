package periodo

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/api_helpers"
	"medroom-backend/formats/Output"
	"medroom-backend/formats/f_input"
	"medroom-backend/models"
	"medroom-backend/repositories"

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

	api_helpers.RespondJSON(c, 200, Output.AddNewPeriodoOutput(periodo))
}

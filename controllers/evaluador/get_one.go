package evaluador

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene un evaluador
// @Description Obtiene un evaluador según su UUID
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a buscar"
// @Success 200 {object} Swagger.GetOneEvaluadorSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores/{uuid_evaluador} [get]
func Get(c *gin.Context) {
	id := c.Params.ByName("id")

	var evaluador models.Evaluador
	if err := repositories.GetOneEvaluador(&evaluador, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.GetOneEvaluador(container))
	api_helpers.RespondJSON(c, 200, evaluador)
}

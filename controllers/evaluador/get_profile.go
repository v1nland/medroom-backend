package evaluador

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene el perfil del evaluador
// @Description Obtiene el perfil del evaluador según su token
// @Tags Evaluadores
// @Accept  json
// @Produce  json
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /evaluadores/me [get]
func Profile(c *gin.Context) {
	id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var evaluador models.Evaluador
	if err := repositories.GetOneEvaluador(&evaluador, id_evaluador); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.GetMyEvaluador(container))
	api_helpers.RespondJson(c, 200, evaluador)
}

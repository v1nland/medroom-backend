package evaluador

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de evaluadores
// @Description Lista todos los evaluadores
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/evaluadores [get]
func List(c *gin.Context) {
	var evaluadores []models.Evaluador
	if err := repositories.GetAllEvaluadores(&evaluadores); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEvaluadores(container))
	api_helpers.RespondJson(c, 200, evaluadores)
}

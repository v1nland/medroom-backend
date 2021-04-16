package evaluador

import (
	"errors"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Lista de evaluadores
// @Description Lista todos los evaluadores
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListEvaluadoresSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores [get]
func ListEvaluadores(c *gin.Context) {
	var evaluadores []models.Evaluador
	if err := repositories.GetAllEvaluadores(&evaluadores); err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			api_helpers.RespondJSON(c, 200, evaluadores)
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEvaluadores(container))
	api_helpers.RespondJSON(c, 200, evaluadores)
}

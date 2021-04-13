package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de grupos
// @Description Lista todos los grupos
// @Tags Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-academica/cursos/{id_periodo}/{sigla_curso}/grupos [get]
func ListGrupos(c *gin.Context) {
	var grupos []models.Grupo

	if err := repositories.GetAllGrupos(&grupos); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, grupos)
}

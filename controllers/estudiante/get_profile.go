package estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

type profileResponse struct {
	Rut_estudiante                string `json:"rut_estudiante"`
	Nombres_estudiante            string `json:"nombres_estudiante"`
	Apellidos_estudiante          string `json:"apellidos_estudiante"`
	Correo_electronico_estudiante string `json:"correo_electronico_estudiante"`
	Telefono_celular_estudiante   string `json:"telefono_celular_estudiante"`
}

// @Summary Obtiene el perfil del estudiante
// @Description Obtiene el perfil del estudiante según su token
// @Tags Estudiantes
// @Accept  json
// @Produce  json
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /estudiantes/me [get]
func Profile(c *gin.Context) {
	id_estudiante := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var estudiante models.Estudiante
	if err := repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	response := profileResponse{
		Rut_estudiante:                estudiante.Rut_estudiante,
		Nombres_estudiante:            estudiante.Nombres_estudiante,
		Apellidos_estudiante:          estudiante.Apellidos_estudiante,
		Correo_electronico_estudiante: estudiante.Correo_electronico_estudiante,
		Telefono_celular_estudiante:   estudiante.Telefono_celular_estudiante,
	}

	api_helpers.RespondJson(c, 200, response)
}

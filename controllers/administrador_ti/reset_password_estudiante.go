package administrador_ti

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

type resetPasswordEstudianteRequest struct {
}

func resetPasswordEstudianteRequestParse(u *resetPasswordEstudianteRequest) {
}

// @Summary Reestablecer password estudiante
// @Description Reestablecer password estudiante por defecto
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a modificar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/estudiantes/{uuid_estudiante}/reestablecer [put]
func ResetPasswordEstudiante(c *gin.Context) {
	id_estudiante := c.Params.ByName("id_estudiante")

	var estudiante models.Estudiante
	if err := repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	// newPass := strings.ReplaceAll(strings.ReplaceAll(estudiante.Rut_estudiante, ".", ""), "-", "")
	// newPassBytes := sha256.Sum256([]byte(newPass))
	// estudiante.Hash_contrasena_estudiante = hex.EncodeToString(newPassBytes[:])
	estudiante.Hash_contrasena_estudiante = utils.GeneratePassword(estudiante.Rut_estudiante)

	if err := repositories.PutOneEstudiante(&estudiante, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, estudiante)
}

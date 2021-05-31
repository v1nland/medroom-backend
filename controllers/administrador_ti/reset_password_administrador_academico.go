package administrador_ti

import (
	"encoding/hex"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"strings"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
)

type resetPasswordAdministradorAcademicoRequest struct {
}

func resetPasswordAdministradorAcademicoRequestParse(u *resetPasswordAdministradorAcademicoRequest) {
}

// @Summary Reestablecer password administrador_academico
// @Description Reestablecer password administrador_academico por defecto
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a modificar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/administradores-academicos/{uuid_administrador_academico}/reestablecer [put]
func ResetPasswordAdministradorAcademico(c *gin.Context) {
	id_administrador_academico := c.Params.ByName("id_administrador_academico")

	var administrador_academico models.AdministradorAcademico
	if err := repositories.GetAdministradorAcademico(&administrador_academico, id_administrador_academico); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	newPass := strings.ReplaceAll(strings.ReplaceAll(administrador_academico.Rut_administrador_academico, ".", ""), "-", "")
	newPassBytes := sha256.Sum256([]byte(newPass))
	administrador_academico.Hash_contrasena_administrador_academico = hex.EncodeToString(newPassBytes[:])

	if err := repositories.PutAdministradorAcademico(&administrador_academico, id_administrador_academico); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, administrador_academico)
}

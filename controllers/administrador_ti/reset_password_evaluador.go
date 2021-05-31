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

type resetPasswordEvaluadorRequest struct {
}

func resetPasswordEvaluadorRequestParse(u *resetPasswordEvaluadorRequest) {
}

// @Summary Reestablecer password evaluador
// @Description Reestablecer password evaluador por defecto
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a modificar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/evaluadores/{uuid_evaluador}/reestablecer [put]
func ResetPasswordEvaluador(c *gin.Context) {
	id_evaluador := c.Params.ByName("id_evaluador")

	var evaluador models.Evaluador
	if err := repositories.GetOneEvaluador(&evaluador, id_evaluador); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	newPass := strings.ReplaceAll(strings.ReplaceAll(evaluador.Rut_evaluador, ".", ""), "-", "")
	newPassBytes := sha256.Sum256([]byte(newPass))
	evaluador.Hash_contrasena_evaluador = hex.EncodeToString(newPassBytes[:])

	if err := repositories.PutOneEvaluador(&evaluador, id_evaluador); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, evaluador)
}
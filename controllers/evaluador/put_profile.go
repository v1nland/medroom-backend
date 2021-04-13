package evaluador

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type putProfileRequest struct {
	Hash_contrasena_evaluador       *string `json:"hash_contrasena_evaluador"`
	Hash_nueva_contrasena_evaluador *string `json:"hash_nueva_contrasena_evaluador"`
	Telefono_fijo_evaluador         *string `json:"telefono_fijo_evaluador"`
	Telefono_celular_evaluador      *string `json:"telefono_celular_evaluador"`
	Cargo_evaluador                 *string `json:"cargo_evaluador"`
}

func putProfileRequestParse(u *putProfileRequest) {
	if u.Telefono_fijo_evaluador != nil {
		*u.Telefono_fijo_evaluador = strings.TrimSpace(*u.Telefono_fijo_evaluador)
		*u.Telefono_fijo_evaluador = strings.ToUpper(*u.Telefono_fijo_evaluador)
		*u.Telefono_fijo_evaluador = utils.RemoveAccents(*u.Telefono_fijo_evaluador)
	}
	if u.Telefono_celular_evaluador != nil {
		*u.Telefono_celular_evaluador = strings.TrimSpace(*u.Telefono_celular_evaluador)
		*u.Telefono_celular_evaluador = strings.ToUpper(*u.Telefono_celular_evaluador)
		*u.Telefono_celular_evaluador = utils.RemoveAccents(*u.Telefono_celular_evaluador)
	}
	if u.Cargo_evaluador != nil {
		*u.Cargo_evaluador = strings.TrimSpace(*u.Cargo_evaluador)
		*u.Cargo_evaluador = strings.ToUpper(*u.Cargo_evaluador)
		*u.Cargo_evaluador = utils.RemoveAccents(*u.Cargo_evaluador)
	}
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil del propio evaluador con los datos entregados
// @Tags Evaluadores
// @Accept  json
// @Produce  json
// @Param   input_actualiza_evaluador     body    putProfileRequest     true        "Evaluador a modificar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /evaluadores/me [put]
func PutProfile(c *gin.Context) {
	id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var input putProfileRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	putProfileRequestParse(&input)

	var evaluador models.Evaluador
	if err := repositories.GetOneEvaluador(&evaluador, id_evaluador); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	if evaluador.Hash_contrasena_evaluador != *input.Hash_contrasena_evaluador {
		api_helpers.RespondJson(c, 403, "Wrong current password")
		return
	}

	evaluador = models.Evaluador{
		Id:                           evaluador.Id,
		Id_rol:                       evaluador.Id_rol,
		Rol_evaluador:                evaluador.Rol_evaluador,
		Rut_evaluador:                evaluador.Rut_evaluador,
		Nombres_evaluador:            evaluador.Nombres_evaluador,
		Apellidos_evaluador:          evaluador.Apellidos_evaluador,
		Hash_contrasena_evaluador:    utils.CheckNullString(input.Hash_nueva_contrasena_evaluador, evaluador.Hash_contrasena_evaluador),
		Correo_electronico_evaluador: evaluador.Correo_electronico_evaluador,
		Telefono_fijo_evaluador:      utils.CheckNullString(input.Telefono_fijo_evaluador, evaluador.Telefono_fijo_evaluador),
		Telefono_celular_evaluador:   utils.CheckNullString(input.Telefono_celular_evaluador, evaluador.Telefono_celular_evaluador),
		Recinto_evaluador:            evaluador.Recinto_evaluador,
		Cargo_evaluador:              utils.CheckNullString(input.Cargo_evaluador, evaluador.Cargo_evaluador),
	}

	if err := repositories.PutOneEvaluador(&evaluador, id_evaluador); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.PutMyEvaluador(model))
	api_helpers.RespondJson(c, 200, evaluador)
}

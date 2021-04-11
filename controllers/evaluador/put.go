package evaluador

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type putRequest struct {
	Rut_evaluador                *string `json:"rut_evaluador"`
	Nombres_evaluador            *string `json:"nombres_evaluador"`
	Apellidos_evaluador          *string `json:"apellidos_evaluador"`
	Hash_contrasena_evaluador    *string `json:"hash_contrasena_evaluador"`
	Correo_electronico_evaluador *string `json:"correo_electronico_evaluador"`
	Telefono_fijo_evaluador      *string `json:"telefono_fijo_evaluador"`
	Telefono_celular_evaluador   *string `json:"telefono_celular_evaluador"`
	Recinto_evaluador            *string `json:"recinto_evaluador"`
	Cargo_evaluador              *string `json:"cargo_evaluador"`
}

func putRequestParse(u *putRequest) {
	if u.Rut_evaluador != nil {
		*u.Rut_evaluador = strings.TrimSpace(*u.Rut_evaluador)
		*u.Rut_evaluador = strings.ToUpper(*u.Rut_evaluador)
		*u.Rut_evaluador = utils.RemoveAccents(*u.Rut_evaluador)
	}
	if u.Nombres_evaluador != nil {
		*u.Nombres_evaluador = strings.TrimSpace(*u.Nombres_evaluador)
		*u.Nombres_evaluador = strings.ToUpper(*u.Nombres_evaluador)
		*u.Nombres_evaluador = utils.RemoveAccents(*u.Nombres_evaluador)
	}
	if u.Apellidos_evaluador != nil {
		*u.Apellidos_evaluador = strings.TrimSpace(*u.Apellidos_evaluador)
		*u.Apellidos_evaluador = strings.ToUpper(*u.Apellidos_evaluador)
		*u.Apellidos_evaluador = utils.RemoveAccents(*u.Apellidos_evaluador)
	}
	if u.Correo_electronico_evaluador != nil {
		*u.Correo_electronico_evaluador = strings.TrimSpace(*u.Correo_electronico_evaluador)
		*u.Correo_electronico_evaluador = strings.ToUpper(*u.Correo_electronico_evaluador)
		*u.Correo_electronico_evaluador = utils.RemoveAccents(*u.Correo_electronico_evaluador)
	}
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
	if u.Recinto_evaluador != nil {
		*u.Recinto_evaluador = strings.TrimSpace(*u.Recinto_evaluador)
		*u.Recinto_evaluador = strings.ToUpper(*u.Recinto_evaluador)
		*u.Recinto_evaluador = utils.RemoveAccents(*u.Recinto_evaluador)
	}
	if u.Cargo_evaluador != nil {
		*u.Cargo_evaluador = strings.TrimSpace(*u.Cargo_evaluador)
		*u.Cargo_evaluador = strings.ToUpper(*u.Cargo_evaluador)
		*u.Cargo_evaluador = utils.RemoveAccents(*u.Cargo_evaluador)
	}
}

// @Summary Modifica un evaluador
// @Description Modifica un evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a modificar"
// @Param   input_actualiza_evaluador     body    Request.Put     true        "Evaluador a modificar"
// @Success 200 {object} Swagger.PutOneEvaluadorSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores/{uuid_evaluador} [put]
func Put(c *gin.Context) {
	id := c.Params.ByName("id")

	var input putRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	putRequestParse(&input)

	var evaluador models.Evaluador
	if err := repositories.GetOneEvaluador(&evaluador, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	evaluador = models.Evaluador{
		Id:                           evaluador.Id,
		Id_rol:                       evaluador.Id_rol,
		Rol_evaluador:                evaluador.Rol_evaluador,
		Rut_evaluador:                utils.CheckNullString(input.Rut_evaluador, evaluador.Rut_evaluador),
		Nombres_evaluador:            utils.CheckNullString(input.Nombres_evaluador, evaluador.Nombres_evaluador),
		Apellidos_evaluador:          utils.CheckNullString(input.Apellidos_evaluador, evaluador.Apellidos_evaluador),
		Hash_contrasena_evaluador:    utils.CheckNullString(input.Hash_contrasena_evaluador, evaluador.Hash_contrasena_evaluador),
		Correo_electronico_evaluador: utils.CheckNullString(input.Correo_electronico_evaluador, evaluador.Correo_electronico_evaluador),
		Telefono_fijo_evaluador:      utils.CheckNullString(input.Telefono_fijo_evaluador, evaluador.Telefono_fijo_evaluador),
		Telefono_celular_evaluador:   utils.CheckNullString(input.Telefono_celular_evaluador, evaluador.Telefono_celular_evaluador),
		Recinto_evaluador:            utils.CheckNullString(input.Recinto_evaluador, evaluador.Recinto_evaluador),
		Cargo_evaluador:              utils.CheckNullString(input.Cargo_evaluador, evaluador.Cargo_evaluador),
	}

	if err := repositories.PutOneEvaluador(&evaluador, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.PutOneEvaluador(model))
	api_helpers.RespondJSON(c, 200, evaluador)
}

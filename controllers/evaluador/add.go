package evaluador

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type addRequest struct {
	Id_rol                       *int    `json:"id_rol"`
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

func addRequestParse(u *addRequest) {
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

// @Summary Agrega un nuevo evaluador
// @Description Genera un nuevo evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_evaluador     body    Request.Add     true        "Evaluador a agregar"
// @Success 200 {object} Swagger.AddNewEvaluadorSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores [post]
func Add(c *gin.Context) {
	var input addRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	addRequestParse(&input)

	evaluador := models.Evaluador{
		Id_rol:                       *input.Id_rol,
		Rut_evaluador:                *input.Rut_evaluador,
		Nombres_evaluador:            *input.Nombres_evaluador,
		Apellidos_evaluador:          *input.Apellidos_evaluador,
		Hash_contrasena_evaluador:    *input.Hash_contrasena_evaluador,
		Correo_electronico_evaluador: *input.Correo_electronico_evaluador,
		Telefono_fijo_evaluador:      *input.Telefono_fijo_evaluador,
		Telefono_celular_evaluador:   *input.Telefono_celular_evaluador,
		Recinto_evaluador:            *input.Recinto_evaluador,
		Cargo_evaluador:              *input.Cargo_evaluador,
	}

	if err := repositories.AddNewEvaluador(&evaluador); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.AddNewEvaluador(model))
	api_helpers.RespondJSON(c, 200, evaluador)
}

package evaluador

import (
	"errors"
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_input"
	"medroom-backend/messages/Request"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Modifica un evaluador
// @Description Modifica un evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a modificar"
// @Param   input_actualiza_evaluador     body    Request.PutOneEvaluador     true        "Evaluador a modificar"
// @Success 200 {object} Swagger.PutOneEvaluadorSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores/{uuid_evaluador} [put]
func PutOneEvaluador(c *gin.Context) {
	id := c.Params.ByName("id")

	var input Request.PutOneEvaluador
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.PutOneEvaluador(&input)

	var evaluador models.Evaluador
	if err := repositories.GetOneEvaluador(&evaluador, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Evaluador not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.PutOneEvaluador(model))
	api_helpers.RespondJSON(c, 200, evaluador)
}

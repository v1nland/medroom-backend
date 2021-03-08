package evaluador

import (
	"errors"
	"medroom-backend/app/Messages/Request"
	"medroom-backend/app/Utils"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/formats/f_input"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Modifica mi perfil
// @Description Modifica el perfil del propio evaluador con los datos entregados
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   input_actualiza_evaluador     body    Request.PutMyEvaluador     true        "Evaluador a modificar"
// @Success 200 {object} Swagger.PutMyEvaluadorSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /evaluadores/me [put]
func PutMyEvaluador(c *gin.Context) {
	id := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var input Request.PutMyEvaluador
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.PutMyEvaluador(&input)

	var evaluador models.Evaluador
	if err := repositories.GetOneEvaluador(&evaluador, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	evaluador = models.Evaluador{
		Id:                           evaluador.Id,
		Id_rol:                       evaluador.Id_rol,
		Rol_evaluador:                evaluador.Rol_evaluador,
		Rut_evaluador:                evaluador.Rut_evaluador,
		Nombres_evaluador:            evaluador.Nombres_evaluador,
		Apellidos_evaluador:          evaluador.Apellidos_evaluador,
		Hash_contrasena_evaluador:    Utils.CheckNullString(input.Hash_contrasena_evaluador, evaluador.Hash_contrasena_evaluador),
		Correo_electronico_evaluador: evaluador.Correo_electronico_evaluador,
		Telefono_fijo_evaluador:      Utils.CheckNullString(input.Telefono_fijo_evaluador, evaluador.Telefono_fijo_evaluador),
		Telefono_celular_evaluador:   Utils.CheckNullString(input.Telefono_celular_evaluador, evaluador.Telefono_celular_evaluador),
		Recinto_evaluador:            evaluador.Recinto_evaluador,
		Cargo_evaluador:              Utils.CheckNullString(input.Cargo_evaluador, evaluador.Cargo_evaluador),
	}

	if err := repositories.PutOneEvaluador(&evaluador, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.PutMyEvaluador(model))
	api_helpers.RespondJSON(c, 200, evaluador)
}

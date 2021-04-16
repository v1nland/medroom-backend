package evaluador

import (
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_input"
	"medroom-backend/messages/Request"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Agrega un nuevo evaluador
// @Description Genera un nuevo evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_evaluador     body    Request.AddNewEvaluador     true        "Evaluador a agregar"
// @Success 200 {object} Swagger.AddNewEvaluadorSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores [post]
func AddNewEvaluador(c *gin.Context) {
	var input Request.AddNewEvaluador
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.AddNewEvaluador(&input)

	model := models.Evaluador{
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

	err := repositories.AddNewEvaluador(&model)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.AddNewEvaluador(model))
	api_helpers.RespondJSON(c, 200, model)
}

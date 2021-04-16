package estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_input"
	"medroom-backend/messages/Request"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Agrega un nuevo estudiante
// @Description Genera un nuevo estudiante con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_estudiante     body    Request.AddNewEstudiante     true        "Estudiante a agregar"
// @Success 200 {object} Swagger.AddNewEstudianteSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes [post]
func AddNewEstudiante(c *gin.Context) {
	var input Request.AddNewEstudiante
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.AddNewEstudiante(&input)

	estudiante := models.Estudiante{
		Id_rol:                        *input.Id_rol,
		Rut_estudiante:                *input.Rut_estudiante,
		Nombres_estudiante:            *input.Nombres_estudiante,
		Apellidos_estudiante:          *input.Apellidos_estudiante,
		Hash_contrasena_estudiante:    *input.Hash_contrasena_estudiante,
		Correo_electronico_estudiante: *input.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      *input.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   *input.Telefono_celular_estudiante,
	}

	if err := repositories.AddNewEstudiante(&estudiante); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.AddNewEstudiante(model))
	api_helpers.RespondJSON(c, 200, estudiante)
}

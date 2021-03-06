package estudiante

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

// @Summary Modifica mi perfil
// @Description Modifica el perfil del propio estudiante con los datos entregados
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   input_actualiza_estudiante     body    Request.PutMyEstudiante     true        "Nuevos datos del estudiante a modificar"
// @Success 200 {object} Swagger.PutMyEstudianteSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /estudiantes/me [put]
func PutMyEstudiante(c *gin.Context) {
	id := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var input Request.PutMyEstudiante
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.PutMyEstudiante(&input)

	var estudiante models.Estudiante
	if err := repositories.GetOneEstudiante(&estudiante, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	if estudiante.Hash_contrasena_estudiante != *input.Hash_contrasena_estudiante {
		api_helpers.RespondJSON(c, 403, "Current password mismatch")
		return
	}

	estudiante = models.Estudiante{
		Id:                            estudiante.Id,
		Id_rol:                        estudiante.Id_rol,
		Rol_estudiante:                estudiante.Rol_estudiante,
		Rut_estudiante:                estudiante.Rut_estudiante,
		Nombres_estudiante:            estudiante.Nombres_estudiante,
		Apellidos_estudiante:          estudiante.Apellidos_estudiante,
		Hash_contrasena_estudiante:    utils.CheckNullString(input.Hash_nueva_contrasena_estudiante, estudiante.Hash_contrasena_estudiante),
		Correo_electronico_estudiante: estudiante.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      utils.CheckNullString(input.Telefono_fijo_estudiante, estudiante.Telefono_fijo_estudiante),
		Telefono_celular_estudiante:   utils.CheckNullString(input.Telefono_celular_estudiante, estudiante.Telefono_celular_estudiante),
	}

	if err := repositories.PutOneEstudiante(&estudiante, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.PutMyEstudiante(model))
	api_helpers.RespondJSON(c, 200, estudiante)
}

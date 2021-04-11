package estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type putProfileRequest struct {
	Hash_contrasena_estudiante       *string `json:"hash_contrasena_estudiante"`
	Hash_nueva_contrasena_estudiante *string `json:"hash_nueva_contrasena_estudiante"`
	Telefono_fijo_estudiante         *string `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante      *string `json:"telefono_celular_estudiante"`
}

func putProfileRequestParse(u *putProfileRequest) {
	if u.Telefono_fijo_estudiante != nil {
		*u.Telefono_fijo_estudiante = strings.TrimSpace(*u.Telefono_fijo_estudiante)
		*u.Telefono_fijo_estudiante = strings.ToUpper(*u.Telefono_fijo_estudiante)
		*u.Telefono_fijo_estudiante = utils.RemoveAccents(*u.Telefono_fijo_estudiante)
	}
	if u.Telefono_celular_estudiante != nil {
		*u.Telefono_celular_estudiante = strings.TrimSpace(*u.Telefono_celular_estudiante)
		*u.Telefono_celular_estudiante = strings.ToUpper(*u.Telefono_celular_estudiante)
		*u.Telefono_celular_estudiante = utils.RemoveAccents(*u.Telefono_celular_estudiante)
	}
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil del propio estudiante con los datos entregados
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   input_actualiza_estudiante     body    Request.PutProfile     true        "Nuevos datos del estudiante a modificar"
// @Success 200 {object} Swagger.PutMyEstudianteSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /estudiantes/me [put]
func PutProfile(c *gin.Context) {
	id := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var input putProfileRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	putProfileRequestParse(&input)

	var estudiante models.Estudiante
	if err := repositories.GetOneEstudiante(&estudiante, id); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	if estudiante.Hash_contrasena_estudiante != *input.Hash_contrasena_estudiante {
		api_helpers.RespondJSON(c, 403, "Wrong current password")
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
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.PutMyEstudiante(model))
	api_helpers.RespondJSON(c, 200, estudiante)
}

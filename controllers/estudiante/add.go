package estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type addRequest struct {
	Id_rol                        *int    `json:"id_rol"`
	Rut_estudiante                *string `json:"rut_estudiante"`
	Nombres_estudiante            *string `json:"nombres_estudiante"`
	Apellidos_estudiante          *string `json:"apellidos_estudiante"`
	Hash_contrasena_estudiante    *string `json:"hash_contrasena_estudiante"`
	Correo_electronico_estudiante *string `json:"correo_electronico_estudiante"`
	Telefono_fijo_estudiante      *string `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   *string `json:"telefono_celular_estudiante"`
}

func addRequestParse(u *addRequest) {
	if u.Rut_estudiante != nil {
		*u.Rut_estudiante = strings.TrimSpace(*u.Rut_estudiante)
		*u.Rut_estudiante = strings.ToUpper(*u.Rut_estudiante)
		*u.Rut_estudiante = utils.RemoveAccents(*u.Rut_estudiante)
	}
	if u.Nombres_estudiante != nil {
		*u.Nombres_estudiante = strings.TrimSpace(*u.Nombres_estudiante)
		*u.Nombres_estudiante = strings.ToUpper(*u.Nombres_estudiante)
		*u.Nombres_estudiante = utils.RemoveAccents(*u.Nombres_estudiante)
	}
	if u.Apellidos_estudiante != nil {
		*u.Apellidos_estudiante = strings.TrimSpace(*u.Apellidos_estudiante)
		*u.Apellidos_estudiante = strings.ToUpper(*u.Apellidos_estudiante)
		*u.Apellidos_estudiante = utils.RemoveAccents(*u.Apellidos_estudiante)
	}
	if u.Correo_electronico_estudiante != nil {
		*u.Correo_electronico_estudiante = strings.TrimSpace(*u.Correo_electronico_estudiante)
		*u.Correo_electronico_estudiante = strings.ToUpper(*u.Correo_electronico_estudiante)
		*u.Correo_electronico_estudiante = utils.RemoveAccents(*u.Correo_electronico_estudiante)
	}
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

// @Summary Agrega un nuevo estudiante
// @Description Genera un nuevo estudiante con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_estudiante     body    Request.Add     true        "Estudiante a agregar"
// @Success 200 {object} Swagger.AddNewEstudianteSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes [post]
func Add(c *gin.Context) {
	var input addRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	addRequestParse(&input)

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
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.AddNewEstudiante(model))
	api_helpers.RespondJSON(c, 200, estudiante)
}

package estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type putRequest struct {
	Rut_estudiante                *string `json:"rut_estudiante"`
	Nombres_estudiante            *string `json:"nombres_estudiante"`
	Apellidos_estudiante          *string `json:"apellidos_estudiante"`
	Hash_contrasena_estudiante    *string `json:"hash_contrasena_estudiante"`
	Correo_electronico_estudiante *string `json:"correo_electronico_estudiante"`
	Telefono_fijo_estudiante      *string `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   *string `json:"telefono_celular_estudiante"`
}

func putRequestParse(u *putRequest) {
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

// @Summary Modifica un estudiante
// @Description Modifica un estudiante con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a modificar"
// @Param   input_actualiza_estudiante     body    putRequest     true        "Estudiante a modificar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/estudiantes/{uuid_estudiante} [put]
func Put(c *gin.Context) {
	id_estudiante := c.Params.ByName("id_estudiante")

	var input putRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	putRequestParse(&input)

	var estudiante models.Estudiante
	if err := repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	estudiante = models.Estudiante{
		Id:                            estudiante.Id,
		Id_rol:                        estudiante.Id_rol,
		Rol_estudiante:                estudiante.Rol_estudiante,
		Rut_estudiante:                utils.CheckNullString(input.Rut_estudiante, estudiante.Rut_estudiante),
		Nombres_estudiante:            utils.CheckNullString(input.Nombres_estudiante, estudiante.Nombres_estudiante),
		Apellidos_estudiante:          utils.CheckNullString(input.Apellidos_estudiante, estudiante.Apellidos_estudiante),
		Hash_contrasena_estudiante:    utils.CheckNullString(input.Hash_contrasena_estudiante, estudiante.Hash_contrasena_estudiante),
		Correo_electronico_estudiante: utils.CheckNullString(input.Correo_electronico_estudiante, estudiante.Correo_electronico_estudiante),
		Telefono_fijo_estudiante:      utils.CheckNullString(input.Telefono_fijo_estudiante, estudiante.Telefono_fijo_estudiante),
		Telefono_celular_estudiante:   utils.CheckNullString(input.Telefono_celular_estudiante, estudiante.Telefono_celular_estudiante),
	}

	if err := repositories.PutOneEstudiante(&estudiante, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.PutOneEstudiante(model))
	api_helpers.RespondJson(c, 200, estudiante)
}

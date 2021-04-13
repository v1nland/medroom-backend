package administrador_academico

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type putRequest struct {
	Id_rol                                     *int    `json:"id_rol"`
	Rut_administrador_academico                *string `json:"rut_administrador_academico"`
	Nombres_administrador_academico            *string `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          *string `json:"apellidos_administrador_academico"`
	Hash_contrasena_administrador_academico    *string `json:"hash_contrasena_administrador_academico"`
	Correo_electronico_administrador_academico *string `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      *string `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   *string `json:"telefono_celular_administrador_academico"`
}

func putRequestFormat(u *putRequest) {
	if u.Rut_administrador_academico != nil {
		*u.Rut_administrador_academico = strings.TrimSpace(*u.Rut_administrador_academico)
		*u.Rut_administrador_academico = strings.ToUpper(*u.Rut_administrador_academico)
		*u.Rut_administrador_academico = utils.RemoveAccents(*u.Rut_administrador_academico)
	}

	if u.Nombres_administrador_academico != nil {
		*u.Nombres_administrador_academico = strings.TrimSpace(*u.Nombres_administrador_academico)
		*u.Nombres_administrador_academico = strings.ToUpper(*u.Nombres_administrador_academico)
		*u.Nombres_administrador_academico = utils.RemoveAccents(*u.Nombres_administrador_academico)
	}

	if u.Apellidos_administrador_academico != nil {
		*u.Apellidos_administrador_academico = strings.TrimSpace(*u.Apellidos_administrador_academico)
		*u.Apellidos_administrador_academico = strings.ToUpper(*u.Apellidos_administrador_academico)
		*u.Apellidos_administrador_academico = utils.RemoveAccents(*u.Apellidos_administrador_academico)
	}

	if u.Correo_electronico_administrador_academico != nil {
		*u.Correo_electronico_administrador_academico = strings.TrimSpace(*u.Correo_electronico_administrador_academico)
		*u.Correo_electronico_administrador_academico = strings.ToUpper(*u.Correo_electronico_administrador_academico)
		*u.Correo_electronico_administrador_academico = utils.RemoveAccents(*u.Correo_electronico_administrador_academico)
	}

	if u.Telefono_fijo_administrador_academico != nil {
		*u.Telefono_fijo_administrador_academico = strings.TrimSpace(*u.Telefono_fijo_administrador_academico)
		*u.Telefono_fijo_administrador_academico = strings.ToUpper(*u.Telefono_fijo_administrador_academico)
		*u.Telefono_fijo_administrador_academico = utils.RemoveAccents(*u.Telefono_fijo_administrador_academico)
	}

	if u.Telefono_celular_administrador_academico != nil {
		*u.Telefono_celular_administrador_academico = strings.TrimSpace(*u.Telefono_celular_administrador_academico)
		*u.Telefono_celular_administrador_academico = strings.ToUpper(*u.Telefono_celular_administrador_academico)
		*u.Telefono_celular_administrador_academico = utils.RemoveAccents(*u.Telefono_celular_administrador_academico)
	}
}

// @Summary Modifica un administrador_academico
// @Description Modifica un administrador_academico con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a modificar"
// @Param   input_actualiza_administrador_academico     body    putRequest     true        "Administrador académico a modificar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/administradores-academicos/{uuid_administrador_academico} [put]
func Put(c *gin.Context) {
	id_administrador_academico := c.Params.ByName("id_administrador_academico")

	var input putRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	putRequestFormat(&input)

	var admin models.AdministradorAcademico
	if err := repositories.GetAdministradorAcademico(&admin, id_administrador_academico); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	admin = models.AdministradorAcademico{
		Id:                                      admin.Id,
		Id_rol:                                  admin.Id_rol,
		Rut_administrador_academico:             utils.CheckNullString(input.Rut_administrador_academico, admin.Rut_administrador_academico),
		Nombres_administrador_academico:         utils.CheckNullString(input.Nombres_administrador_academico, admin.Nombres_administrador_academico),
		Apellidos_administrador_academico:       utils.CheckNullString(input.Apellidos_administrador_academico, admin.Apellidos_administrador_academico),
		Hash_contrasena_administrador_academico: utils.CheckNullString(input.Hash_contrasena_administrador_academico, admin.Hash_contrasena_administrador_academico),
		Correo_electronico_administrador_academico: utils.CheckNullString(input.Correo_electronico_administrador_academico, admin.Correo_electronico_administrador_academico),
		Telefono_fijo_administrador_academico:      utils.CheckNullString(input.Telefono_fijo_administrador_academico, admin.Telefono_fijo_administrador_academico),
		Telefono_celular_administrador_academico:   utils.CheckNullString(input.Telefono_celular_administrador_academico, admin.Telefono_celular_administrador_academico),
	}

	if err := repositories.PutAdministradorAcademico(&admin, id_administrador_academico); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, admin)
}

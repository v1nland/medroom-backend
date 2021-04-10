package administrador_academico

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type putProfileRequest struct {
	Id_rol                                        *int    `json:"id_rol"`
	Rut_administrador_academico                   *string `json:"rut_administrador_academico"`
	Nombres_administrador_academico               *string `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico             *string `json:"apellidos_administrador_academico"`
	Hash_contrasena_administrador_academico       *string `json:"hash_contrasena_administrador_academico"`
	Hash_nueva_contrasena_administrador_academico *string `json:"hash_nueva_contrasena_administrador_academico"`
	Correo_electronico_administrador_academico    *string `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico         *string `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico      *string `json:"telefono_celular_administrador_academico"`
}

func putProfileRequestFormat(u *putProfileRequest) {
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

// @Summary Modifica mi perfil
// @Description Modifica el perfil de un administrador_academico con los datos entregados
// @Tags 04 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   input_actualiza_administrador_academico     body    Request.PutProfile     true        "AdministradorAcademico a modificar"
// @Success 200 {object} Swagger.PutMyAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/me [put]
func PutProfile(c *gin.Context) {
	id := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	var input putProfileRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	// format input
	putProfileRequestFormat(&input)

	// generate admin entity
	var admin models.AdministradorAcademico
	if err := repositories.GetAdministradorAcademico(&admin, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	if admin.Hash_contrasena_administrador_academico != *input.Hash_contrasena_administrador_academico {
		api_helpers.RespondJSON(c, 403, "Wrong current password")
		return
	}

	admin = models.AdministradorAcademico{
		Id:                                      admin.Id,
		Id_rol:                                  admin.Id_rol,
		Rut_administrador_academico:             admin.Rut_administrador_academico,
		Nombres_administrador_academico:         admin.Nombres_administrador_academico,
		Apellidos_administrador_academico:       admin.Apellidos_administrador_academico,
		Hash_contrasena_administrador_academico: utils.CheckNullString(input.Hash_nueva_contrasena_administrador_academico, admin.Hash_contrasena_administrador_academico),
		Correo_electronico_administrador_academico: utils.CheckNullString(input.Correo_electronico_administrador_academico, admin.Correo_electronico_administrador_academico),
		Telefono_fijo_administrador_academico:      utils.CheckNullString(input.Telefono_fijo_administrador_academico, admin.Telefono_fijo_administrador_academico),
		Telefono_celular_administrador_academico:   utils.CheckNullString(input.Telefono_celular_administrador_academico, admin.Telefono_celular_administrador_academico),
	}

	if err := repositories.PutAdministradorAcademico(&admin, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, admin)
}

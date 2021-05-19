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
	Hash_contrasena_administrador_academico       *string `json:"hash_contrasena_administrador_academico"`
	Hash_nueva_contrasena_administrador_academico *string `json:"hash_nueva_contrasena_administrador_academico"`
	Telefono_celular_administrador_academico      *string `json:"telefono_celular_administrador_academico"`
}

func putProfileRequestFormat(u *putProfileRequest) {
	if u.Telefono_celular_administrador_academico != nil {
		*u.Telefono_celular_administrador_academico = strings.TrimSpace(*u.Telefono_celular_administrador_academico)
		*u.Telefono_celular_administrador_academico = strings.ToUpper(*u.Telefono_celular_administrador_academico)
		*u.Telefono_celular_administrador_academico = utils.RemoveAccents(*u.Telefono_celular_administrador_academico)
	}
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil de un administrador_academico con los datos entregados
// @Tags Administración Académica
// @Accept  json
// @Produce  json
// @Param   input_actualiza_administrador_academico     body    putProfileRequest     true        "Administrador académico a modificar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-academica/me [put]
func PutProfile(c *gin.Context) {
	id_administrador_academico := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	var input putProfileRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	// format input
	putProfileRequestFormat(&input)

	// generate admin entity
	var admin models.AdministradorAcademico
	if err := repositories.GetAdministradorAcademico(&admin, id_administrador_academico); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	if admin.Hash_contrasena_administrador_academico != *input.Hash_contrasena_administrador_academico {
		api_helpers.RespondJson(c, 403, "Wrong current password")
		return
	}

	admin = models.AdministradorAcademico{
		Id:                                      admin.Id,
		Id_rol:                                  admin.Id_rol,
		Rut_administrador_academico:             admin.Rut_administrador_academico,
		Nombres_administrador_academico:         admin.Nombres_administrador_academico,
		Apellidos_administrador_academico:       admin.Apellidos_administrador_academico,
		Hash_contrasena_administrador_academico: utils.CheckNullString(input.Hash_nueva_contrasena_administrador_academico, admin.Hash_contrasena_administrador_academico),
		Correo_electronico_administrador_academico: admin.Correo_electronico_administrador_academico,
		Telefono_fijo_administrador_academico:      admin.Telefono_fijo_administrador_academico,
		Telefono_celular_administrador_academico:   utils.CheckNullString(input.Telefono_celular_administrador_academico, admin.Telefono_celular_administrador_academico),
	}

	if err := repositories.PutAdministradorAcademico(&admin, id_administrador_academico); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, admin)
}

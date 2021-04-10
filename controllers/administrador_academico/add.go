package administrador_academico

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type addRequest struct {
	Id_rol                                     *int    `json:"id_rol"`
	Rut_administrador_academico                *string `json:"rut_administrador_academico"`
	Nombres_administrador_academico            *string `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          *string `json:"apellidos_administrador_academico"`
	Hash_contrasena_administrador_academico    *string `json:"hash_contrasena_administrador_academico"`
	Correo_electronico_administrador_academico *string `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      *string `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   *string `json:"telefono_celular_administrador_academico"`
}

func addRequestParse(u *addRequest) {
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

// @Summary Agrega un nuevo administrador_academico
// @Description Genera un nuevo administrador_academico con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_administrador_academico     body    Request.Add     true        "AdministradorAcademico a agregar"
// @Success 200 {object} Swagger.AddNewAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos [post]
func Add(c *gin.Context) {

	// read payload
	var input addRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	addRequestParse(&input)

	admin := models.AdministradorAcademico{
		Id_rol:                                     *input.Id_rol,
		Rut_administrador_academico:                *input.Rut_administrador_academico,
		Nombres_administrador_academico:            *input.Nombres_administrador_academico,
		Apellidos_administrador_academico:          *input.Apellidos_administrador_academico,
		Hash_contrasena_administrador_academico:    *input.Hash_contrasena_administrador_academico,
		Correo_electronico_administrador_academico: *input.Correo_electronico_administrador_academico,
		Telefono_fijo_administrador_academico:      *input.Telefono_fijo_administrador_academico,
		Telefono_celular_administrador_academico:   *input.Telefono_celular_administrador_academico,
	}

	if err := repositories.AddAdministradorAcademico(&admin); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, admin.Id)
}

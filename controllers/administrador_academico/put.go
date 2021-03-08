package administrador_academico

import (
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_input"
	"medroom-backend/formats/f_output"
	"medroom-backend/messages/Request"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Modifica un administrador_academico
// @Description Modifica un administrador_academico con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a modificar"
// @Param   input_actualiza_administrador_academico     body    Request.PutOneAdministradorAcademico     true        "AdministradorAcademico a modificar"
// @Success 200 {object} Swagger.PutOneAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos/{uuid_administrador_academico} [put]
func PutOneAdministradorAcademico(c *gin.Context) {
	id := c.Params.ByName("id")

	var input Request.PutOneAdministradorAcademico
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.PutOneAdministradorAcademico(&input)

	var administrador_academico models.AdministradorAcademico
	if err := repositories.GetOneAdministradorAcademico(&administrador_academico, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	administrador_academico = models.AdministradorAcademico{
		Id:                                      administrador_academico.Id,
		Id_rol:                                  administrador_academico.Id_rol,
		Rut_administrador_academico:             utils.CheckNullString(input.Rut_administrador_academico, administrador_academico.Rut_administrador_academico),
		Nombres_administrador_academico:         utils.CheckNullString(input.Nombres_administrador_academico, administrador_academico.Nombres_administrador_academico),
		Apellidos_administrador_academico:       utils.CheckNullString(input.Apellidos_administrador_academico, administrador_academico.Apellidos_administrador_academico),
		Hash_contrasena_administrador_academico: utils.CheckNullString(input.Hash_contrasena_administrador_academico, administrador_academico.Hash_contrasena_administrador_academico),
		Correo_electronico_administrador_academico: utils.CheckNullString(input.Correo_electronico_administrador_academico, administrador_academico.Correo_electronico_administrador_academico),
		Telefono_fijo_administrador_academico:      utils.CheckNullString(input.Telefono_fijo_administrador_academico, administrador_academico.Telefono_fijo_administrador_academico),
		Telefono_celular_administrador_academico:   utils.CheckNullString(input.Telefono_celular_administrador_academico, administrador_academico.Telefono_celular_administrador_academico),
	}

	if err := repositories.PutOneAdministradorAcademico(&administrador_academico, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// output
	api_helpers.RespondJSON(c, 200, f_output.PutOneAdministradorAcademico(administrador_academico))
}

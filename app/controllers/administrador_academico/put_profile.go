package administrador_academico

import (
	"errors"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/formats/f_input"
	"medroom-backend/app/formats/f_output"
	"medroom-backend/app/messages/Request"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
	"medroom-backend/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Modifica mi perfil
// @Description Modifica el perfil de un administrador_academico con los datos entregados
// @Tags 04 - Administración Academica
// @Accept  json
// @Produce  json
// @Param   input_actualiza_administrador_academico     body    Request.PutMyAdministradorAcademico     true        "AdministradorAcademico a modificar"
// @Success 200 {object} Swagger.PutMyAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/me [put]
func PutMyAdministradorAcademico(c *gin.Context) {
	// params
	id_administrador_academico := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	// input input
	var input Request.PutMyAdministradorAcademico
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	// format input
	f_input.PutMyAdministradorAcademico(&input)

	// generate model entity
	var model models.AdministradorAcademico
	if err := repositories.GetOneAdministradorAcademico(&model, id_administrador_academico); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Administrador académico not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// replace data in model entity
	model = models.AdministradorAcademico{
		Id:                                      model.Id,
		Id_rol:                                  model.Id_rol,
		Rut_administrador_academico:             model.Rut_administrador_academico,
		Nombres_administrador_academico:         model.Nombres_administrador_academico,
		Apellidos_administrador_academico:       model.Apellidos_administrador_academico,
		Hash_contrasena_administrador_academico: utils.CheckNullString(input.Hash_contrasena_administrador_academico, model.Hash_contrasena_administrador_academico),
		Correo_electronico_administrador_academico: utils.CheckNullString(input.Correo_electronico_administrador_academico, model.Correo_electronico_administrador_academico),
		Telefono_fijo_administrador_academico:      utils.CheckNullString(input.Telefono_fijo_administrador_academico, model.Telefono_fijo_administrador_academico),
		Telefono_celular_administrador_academico:   utils.CheckNullString(input.Telefono_celular_administrador_academico, model.Telefono_celular_administrador_academico),
	}

	// put query
	if err := repositories.PutOneAdministradorAcademico(&model, id_administrador_academico); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Administrador académico not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// output
	api_helpers.RespondJSON(c, 200, f_output.PutMyAdministradorAcademico(model))
}

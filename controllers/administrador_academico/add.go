package administrador_academico

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/api_helpers"
	"medroom-backend/formats/Output"
	"medroom-backend/formats/f_input"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Agrega un nuevo administrador_academico
// @Description Genera un nuevo administrador_academico con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_administrador_academico     body    Request.AddNewAdministradorAcademico     true        "AdministradorAcademico a agregar"
// @Success 200 {object} Swagger.AddNewAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos [post]
func AddNewAdministradorAcademico(c *gin.Context) {
	var input Request.AddNewAdministradorAcademico

	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.AddNewAdministradorAcademico(&input)

	model_container := models.AdministradorAcademico{
		Id_rol:                                     *input.Id_rol,
		Rut_administrador_academico:                *input.Rut_administrador_academico,
		Nombres_administrador_academico:            *input.Nombres_administrador_academico,
		Apellidos_administrador_academico:          *input.Apellidos_administrador_academico,
		Hash_contrasena_administrador_academico:    *input.Hash_contrasena_administrador_academico,
		Correo_electronico_administrador_academico: *input.Correo_electronico_administrador_academico,
		Telefono_fijo_administrador_academico:      *input.Telefono_fijo_administrador_academico,
		Telefono_celular_administrador_academico:   *input.Telefono_celular_administrador_academico,
	}

	// query
	err := repositories.AddNewAdministradorAcademico(&model_container)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// output
	api_helpers.RespondJSON(c, 200, Output.AddNewAdministradorAcademicoOutput(model_container))
}

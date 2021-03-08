package administrador_ti

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_input"
	"medroom-backend/formats/f_output"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Agrega un nuevo administrador_ti
// @Description Genera un nuevo administrador_ti con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_administrador_ti     body    Request.AddNewAdministradorTi     true        "AdministradorTi a agregar"
// @Success 200 {object} Swagger.AddNewAdministradorTiSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti [post]
func AddNewAdministradorTi(c *gin.Context) {
	var input Request.AddNewAdministradorTi
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.AddNewAdministradorTi(&input)

	model_container := models.AdministradorTi{
		Id_rol:                              *input.Id_rol,
		Rut_administrador_ti:                *input.Rut_administrador_ti,
		Nombres_administrador_ti:            *input.Nombres_administrador_ti,
		Apellidos_administrador_ti:          *input.Apellidos_administrador_ti,
		Hash_contrasena_administrador_ti:    *input.Hash_contrasena_administrador_ti,
		Correo_electronico_administrador_ti: *input.Correo_electronico_administrador_ti,
		Telefono_fijo_administrador_ti:      *input.Telefono_fijo_administrador_ti,
		Telefono_celular_administrador_ti:   *input.Telefono_celular_administrador_ti,
	}

	if err := repositories.AddNewAdministradorTi(&model_container); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, f_output.AddNewAdministradorTi(model_container))
}

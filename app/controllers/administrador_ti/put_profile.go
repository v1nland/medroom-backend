package administrador_ti

import (
	"medroom-backend/app/Messages/Request"
	"medroom-backend/app/Utils"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/formats/f_input"
	"medroom-backend/app/formats/f_output"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Modifica mi perfil
// @Description Modifica el perfil de un administrador_ti con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_actualiza_administrador_ti     body    Request.PutMyAdministradorTi     true        "AdministradorTi a modificar"
// @Success 200 {object} Swagger.PutMyAdministradorTiSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/me [put]
func PutMyAdministradorTi(c *gin.Context) {
	id := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_TI")

	var input Request.PutMyAdministradorTi
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.PutMyAdministradorTi(&input)

	var administrador_ti models.AdministradorTi
	if err := repositories.GetOneAdministradorTi(&administrador_ti, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	administrador_ti = models.AdministradorTi{
		Id:                                  administrador_ti.Id,
		Id_rol:                              administrador_ti.Id_rol,
		Rut_administrador_ti:                Utils.CheckNullString(input.Rut_administrador_ti, administrador_ti.Rut_administrador_ti),
		Nombres_administrador_ti:            Utils.CheckNullString(input.Nombres_administrador_ti, administrador_ti.Nombres_administrador_ti),
		Apellidos_administrador_ti:          Utils.CheckNullString(input.Apellidos_administrador_ti, administrador_ti.Apellidos_administrador_ti),
		Hash_contrasena_administrador_ti:    Utils.CheckNullString(input.Hash_contrasena_administrador_ti, administrador_ti.Hash_contrasena_administrador_ti),
		Correo_electronico_administrador_ti: Utils.CheckNullString(input.Correo_electronico_administrador_ti, administrador_ti.Correo_electronico_administrador_ti),
		Telefono_fijo_administrador_ti:      Utils.CheckNullString(input.Telefono_fijo_administrador_ti, administrador_ti.Telefono_fijo_administrador_ti),
		Telefono_celular_administrador_ti:   Utils.CheckNullString(input.Telefono_celular_administrador_ti, administrador_ti.Telefono_celular_administrador_ti),
	}

	if err := repositories.PutOneAdministradorTi(&administrador_ti, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, f_output.PutMyAdministradorTi(administrador_ti))
}

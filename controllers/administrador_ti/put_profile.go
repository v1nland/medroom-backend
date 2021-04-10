package administrador_ti

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type putProfileRequest struct {
	Id_rol                                 *int    `json:"id_rol"`
	Id_grupo                               *int    `json:"id_grupo"`
	Rut_administrador_ti                   *string `json:"rut_administrador_ti"`
	Nombres_administrador_ti               *string `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti             *string `json:"apellidos_administrador_ti"`
	Hash_contrasena_administrador_ti       *string `json:"hash_contrasena_administrador_ti"`
	Hash_nueva_contrasena_administrador_ti *string `json:"hash_nueva_contrasena_administrador_ti"`
	Correo_electronico_administrador_ti    *string `json:"correo_electronico_administrador_ti"`
	Telefono_fijo_administrador_ti         *string `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti      *string `json:"telefono_celular_administrador_ti"`
}

func putProfileRequestParse(u *putProfileRequest) {
	if u.Rut_administrador_ti != nil {
		*u.Rut_administrador_ti = strings.TrimSpace(*u.Rut_administrador_ti)
		*u.Rut_administrador_ti = strings.ToUpper(*u.Rut_administrador_ti)
		*u.Rut_administrador_ti = utils.RemoveAccents(*u.Rut_administrador_ti)
	}

	if u.Nombres_administrador_ti != nil {
		*u.Nombres_administrador_ti = strings.TrimSpace(*u.Nombres_administrador_ti)
		*u.Nombres_administrador_ti = strings.ToUpper(*u.Nombres_administrador_ti)
		*u.Nombres_administrador_ti = utils.RemoveAccents(*u.Nombres_administrador_ti)
	}

	if u.Apellidos_administrador_ti != nil {
		*u.Apellidos_administrador_ti = strings.TrimSpace(*u.Apellidos_administrador_ti)
		*u.Apellidos_administrador_ti = strings.ToUpper(*u.Apellidos_administrador_ti)
		*u.Apellidos_administrador_ti = utils.RemoveAccents(*u.Apellidos_administrador_ti)
	}

	if u.Correo_electronico_administrador_ti != nil {
		*u.Correo_electronico_administrador_ti = strings.TrimSpace(*u.Correo_electronico_administrador_ti)
		*u.Correo_electronico_administrador_ti = strings.ToUpper(*u.Correo_electronico_administrador_ti)
		*u.Correo_electronico_administrador_ti = utils.RemoveAccents(*u.Correo_electronico_administrador_ti)
	}

	if u.Telefono_fijo_administrador_ti != nil {
		*u.Telefono_fijo_administrador_ti = strings.TrimSpace(*u.Telefono_fijo_administrador_ti)
		*u.Telefono_fijo_administrador_ti = strings.ToUpper(*u.Telefono_fijo_administrador_ti)
		*u.Telefono_fijo_administrador_ti = utils.RemoveAccents(*u.Telefono_fijo_administrador_ti)
	}

	if u.Telefono_celular_administrador_ti != nil {
		*u.Telefono_celular_administrador_ti = strings.TrimSpace(*u.Telefono_celular_administrador_ti)
		*u.Telefono_celular_administrador_ti = strings.ToUpper(*u.Telefono_celular_administrador_ti)
		*u.Telefono_celular_administrador_ti = utils.RemoveAccents(*u.Telefono_celular_administrador_ti)
	}
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil de un administrador_ti con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_actualiza_administrador_ti     body    Request.PutProfile     true        "AdministradorTi a modificar"
// @Success 200 {object} Swagger.PutMyAdministradorTiSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/me [put]
func PutProfile(c *gin.Context) {
	id := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_TI")

	var input putProfileRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	putProfileRequestParse(&input)

	var administrador_ti models.AdministradorTi
	if err := repositories.GetOneAdministradorTi(&administrador_ti, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	if administrador_ti.Hash_contrasena_administrador_ti != *input.Hash_contrasena_administrador_ti {
		api_helpers.RespondJSON(c, 403, "Wrong current password")
		return
	}

	// replace data in model entity
	administrador_ti = models.AdministradorTi{
		Id:                                  administrador_ti.Id,
		Id_rol:                              administrador_ti.Id_rol,
		Rut_administrador_ti:                utils.CheckNullString(input.Rut_administrador_ti, administrador_ti.Rut_administrador_ti),
		Nombres_administrador_ti:            utils.CheckNullString(input.Nombres_administrador_ti, administrador_ti.Nombres_administrador_ti),
		Apellidos_administrador_ti:          utils.CheckNullString(input.Apellidos_administrador_ti, administrador_ti.Apellidos_administrador_ti),
		Hash_contrasena_administrador_ti:    utils.CheckNullString(input.Hash_nueva_contrasena_administrador_ti, administrador_ti.Hash_contrasena_administrador_ti),
		Correo_electronico_administrador_ti: utils.CheckNullString(input.Correo_electronico_administrador_ti, administrador_ti.Correo_electronico_administrador_ti),
		Telefono_fijo_administrador_ti:      utils.CheckNullString(input.Telefono_fijo_administrador_ti, administrador_ti.Telefono_fijo_administrador_ti),
		Telefono_celular_administrador_ti:   utils.CheckNullString(input.Telefono_celular_administrador_ti, administrador_ti.Telefono_celular_administrador_ti),
	}

	if err := repositories.PutOneAdministradorTi(&administrador_ti, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, administrador_ti)
}

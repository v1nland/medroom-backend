package administrador_ti

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type putRequest struct {
	Id_rol                              *int    `json:"id_rol"`
	Id_grupo                            *int    `json:"id_grupo"`
	Rut_administrador_ti                *string `json:"rut_administrador_ti"`
	Nombres_administrador_ti            *string `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti          *string `json:"apellidos_administrador_ti"`
	Hash_contrasena_administrador_ti    *string `json:"hash_contrasena_administrador_ti"`
	Correo_electronico_administrador_ti *string `json:"correo_electronico_administrador_ti"`
	Telefono_fijo_administrador_ti      *string `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti   *string `json:"telefono_celular_administrador_ti"`
}

func putRequestParse(u *putRequest) {
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

// @Summary Modifica un administrador_ti
// @Description Modifica un administrador_ti con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_ti     path    string     true        "UUID del administrador_ti a modificar"
// @Param   input_actualiza_administrador_ti     body    Request.Put     true        "AdministradorTi a modificar"
// @Success 200 {object} Swagger.PutOneAdministradorTiSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti/{uuid_administrador_ti} [put]
func Put(c *gin.Context) {
	id := c.Params.ByName("id")

	var input putRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	putRequestParse(&input)

	var admin models.AdministradorTi
	if err := repositories.GetOneAdministradorTi(&admin, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	admin = models.AdministradorTi{
		Id:                                  admin.Id,
		Id_rol:                              admin.Id_rol,
		Rut_administrador_ti:                utils.CheckNullString(input.Rut_administrador_ti, admin.Rut_administrador_ti),
		Nombres_administrador_ti:            utils.CheckNullString(input.Nombres_administrador_ti, admin.Nombres_administrador_ti),
		Apellidos_administrador_ti:          utils.CheckNullString(input.Apellidos_administrador_ti, admin.Apellidos_administrador_ti),
		Hash_contrasena_administrador_ti:    utils.CheckNullString(input.Hash_contrasena_administrador_ti, admin.Hash_contrasena_administrador_ti),
		Correo_electronico_administrador_ti: utils.CheckNullString(input.Correo_electronico_administrador_ti, admin.Correo_electronico_administrador_ti),
		Telefono_fijo_administrador_ti:      utils.CheckNullString(input.Telefono_fijo_administrador_ti, admin.Telefono_fijo_administrador_ti),
		Telefono_celular_administrador_ti:   utils.CheckNullString(input.Telefono_celular_administrador_ti, admin.Telefono_celular_administrador_ti),
	}

	if err := repositories.PutOneAdministradorTi(&admin, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, admin)
}

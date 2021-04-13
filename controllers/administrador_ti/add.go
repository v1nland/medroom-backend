package administrador_ti

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type addRequest struct {
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

func addRequestParse(u *addRequest) {
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

// @Summary Agrega un nuevo administrador_ti
// @Description Genera un nuevo administrador_ti con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_administrador_ti     body    addRequest     true        "Administrador Ti a agregar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/administradores-ti [post]
func Add(c *gin.Context) {
	var input addRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	addRequestParse(&input)

	admin := models.AdministradorTi{
		Id_rol:                              *input.Id_rol,
		Rut_administrador_ti:                *input.Rut_administrador_ti,
		Nombres_administrador_ti:            *input.Nombres_administrador_ti,
		Apellidos_administrador_ti:          *input.Apellidos_administrador_ti,
		Hash_contrasena_administrador_ti:    *input.Hash_contrasena_administrador_ti,
		Correo_electronico_administrador_ti: *input.Correo_electronico_administrador_ti,
		Telefono_fijo_administrador_ti:      *input.Telefono_fijo_administrador_ti,
		Telefono_celular_administrador_ti:   *input.Telefono_celular_administrador_ti,
	}

	if err := repositories.AddNewAdministradorTi(&admin); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, admin)
}

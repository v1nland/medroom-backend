package administrador_ti

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type massiveAddRequest struct {
	AdministradoresTi []struct {
		Rut_administrador_ti                *string `json:"rut_administrador_ti"`
		Nombres_administrador_ti            *string `json:"nombres_administrador_ti"`
		Apellidos_administrador_ti          *string `json:"apellidos_administrador_ti"`
		Hash_contrasena_administrador_ti    *string `json:"hash_contrasena_administrador_ti"`
		Correo_electronico_administrador_ti *string `json:"correo_electronico_administrador_ti"`
		Telefono_fijo_administrador_ti      *string `json:"telefono_fijo_administrador_ti"`
		Telefono_celular_administrador_ti   *string `json:"telefono_celular_administrador_ti"`
	} `json:"administradores_ti"`
}

func massiveAddFormat(u *massiveAddRequest) {
	for i := 0; i < len(u.AdministradoresTi); i++ {
		if u.AdministradoresTi[i].Rut_administrador_ti != nil {
			*u.AdministradoresTi[i].Rut_administrador_ti = strings.TrimSpace(*u.AdministradoresTi[i].Rut_administrador_ti)
			*u.AdministradoresTi[i].Rut_administrador_ti = strings.ToUpper(*u.AdministradoresTi[i].Rut_administrador_ti)
			*u.AdministradoresTi[i].Rut_administrador_ti = utils.RemoveAccents(*u.AdministradoresTi[i].Rut_administrador_ti)
		}
		if u.AdministradoresTi[i].Nombres_administrador_ti != nil {
			*u.AdministradoresTi[i].Nombres_administrador_ti = strings.TrimSpace(*u.AdministradoresTi[i].Nombres_administrador_ti)
			*u.AdministradoresTi[i].Nombres_administrador_ti = strings.ToUpper(*u.AdministradoresTi[i].Nombres_administrador_ti)
			*u.AdministradoresTi[i].Nombres_administrador_ti = utils.RemoveAccents(*u.AdministradoresTi[i].Nombres_administrador_ti)
		}
		if u.AdministradoresTi[i].Apellidos_administrador_ti != nil {
			*u.AdministradoresTi[i].Apellidos_administrador_ti = strings.TrimSpace(*u.AdministradoresTi[i].Apellidos_administrador_ti)
			*u.AdministradoresTi[i].Apellidos_administrador_ti = strings.ToUpper(*u.AdministradoresTi[i].Apellidos_administrador_ti)
			*u.AdministradoresTi[i].Apellidos_administrador_ti = utils.RemoveAccents(*u.AdministradoresTi[i].Apellidos_administrador_ti)
		}
		if u.AdministradoresTi[i].Correo_electronico_administrador_ti != nil {
			*u.AdministradoresTi[i].Correo_electronico_administrador_ti = strings.TrimSpace(*u.AdministradoresTi[i].Correo_electronico_administrador_ti)
			*u.AdministradoresTi[i].Correo_electronico_administrador_ti = strings.ToUpper(*u.AdministradoresTi[i].Correo_electronico_administrador_ti)
			*u.AdministradoresTi[i].Correo_electronico_administrador_ti = utils.RemoveAccents(*u.AdministradoresTi[i].Correo_electronico_administrador_ti)
		}
		if u.AdministradoresTi[i].Telefono_fijo_administrador_ti != nil {
			*u.AdministradoresTi[i].Telefono_fijo_administrador_ti = strings.TrimSpace(*u.AdministradoresTi[i].Telefono_fijo_administrador_ti)
			*u.AdministradoresTi[i].Telefono_fijo_administrador_ti = strings.ToUpper(*u.AdministradoresTi[i].Telefono_fijo_administrador_ti)
			*u.AdministradoresTi[i].Telefono_fijo_administrador_ti = utils.RemoveAccents(*u.AdministradoresTi[i].Telefono_fijo_administrador_ti)
		}
		if u.AdministradoresTi[i].Telefono_celular_administrador_ti != nil {
			*u.AdministradoresTi[i].Telefono_celular_administrador_ti = strings.TrimSpace(*u.AdministradoresTi[i].Telefono_celular_administrador_ti)
			*u.AdministradoresTi[i].Telefono_celular_administrador_ti = strings.ToUpper(*u.AdministradoresTi[i].Telefono_celular_administrador_ti)
			*u.AdministradoresTi[i].Telefono_celular_administrador_ti = utils.RemoveAccents(*u.AdministradoresTi[i].Telefono_celular_administrador_ti)
		}
	}
}

// @Summary Agrega nuevos administrador_ti de forma masiva
// @Description Genera nuevos administrador_ti con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_administrador_ti     body    massiveAddRequest     true        "Administradores Ti a agregar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/administradores-ti/carga-masiva [post]
func MassiveAdd(c *gin.Context) {
	var payload massiveAddRequest
	if err := c.ShouldBind(&payload); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	massiveAddFormat(&payload)

	var administradores_tis_error []string
	for i := 0; i < len(payload.AdministradoresTi); i++ {
		administrador_ti := models.AdministradorTi{
			Id_rol:                              4,
			Rut_administrador_ti:                *payload.AdministradoresTi[i].Rut_administrador_ti,
			Nombres_administrador_ti:            *payload.AdministradoresTi[i].Nombres_administrador_ti,
			Apellidos_administrador_ti:          *payload.AdministradoresTi[i].Apellidos_administrador_ti,
			Hash_contrasena_administrador_ti:    *payload.AdministradoresTi[i].Hash_contrasena_administrador_ti,
			Correo_electronico_administrador_ti: *payload.AdministradoresTi[i].Correo_electronico_administrador_ti,
			Telefono_fijo_administrador_ti:      *payload.AdministradoresTi[i].Telefono_fijo_administrador_ti,
			Telefono_celular_administrador_ti:   *payload.AdministradoresTi[i].Telefono_celular_administrador_ti,
		}

		if err := repositories.AddNewAdministradorTi(&administrador_ti); err != nil {
			administradores_tis_error = append(administradores_tis_error, "["+*payload.AdministradoresTi[i].Rut_administrador_ti+"] "+*payload.AdministradoresTi[i].Nombres_administrador_ti+" "+*payload.AdministradoresTi[i].Apellidos_administrador_ti)
		}
	}

	if len(administradores_tis_error) > 0 {
		api_helpers.RespondJson(c, 201, administradores_tis_error)
	} else {
		api_helpers.RespondJson(c, 200, "ok")
	}
}

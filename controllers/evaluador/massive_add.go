package evaluador

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type massiveAddRequest struct {
	Evaluadores []struct {
		Id_grupos                    []*int  `json:"id_grupos"`
		Rut_evaluador                *string `json:"rut_evaluador"`
		Nombres_evaluador            *string `json:"nombres_evaluador"`
		Apellidos_evaluador          *string `json:"apellidos_evaluador"`
		Hash_contrasena_evaluador    *string `json:"hash_contrasena_evaluador"`
		Correo_electronico_evaluador *string `json:"correo_electronico_evaluador"`
		Telefono_fijo_evaluador      *string `json:"telefono_fijo_evaluador"`
		Telefono_celular_evaluador   *string `json:"telefono_celular_evaluador"`
		Recinto_evaluador            *string `json:"recinto_evaluador"`
		Cargo_evaluador              *string `json:"cargo_evaluador"`
	} `json:"evaluadores"`
}

func massiveAddRequestParse(u *massiveAddRequest) {
	for i := 0; i < len(u.Evaluadores); i++ {
		if u.Evaluadores[i].Rut_evaluador != nil {
			*u.Evaluadores[i].Rut_evaluador = strings.TrimSpace(*u.Evaluadores[i].Rut_evaluador)
			*u.Evaluadores[i].Rut_evaluador = strings.ToUpper(*u.Evaluadores[i].Rut_evaluador)
			*u.Evaluadores[i].Rut_evaluador = utils.RemoveAccents(*u.Evaluadores[i].Rut_evaluador)
		}
		if u.Evaluadores[i].Nombres_evaluador != nil {
			*u.Evaluadores[i].Nombres_evaluador = strings.TrimSpace(*u.Evaluadores[i].Nombres_evaluador)
			*u.Evaluadores[i].Nombres_evaluador = strings.ToUpper(*u.Evaluadores[i].Nombres_evaluador)
			*u.Evaluadores[i].Nombres_evaluador = utils.RemoveAccents(*u.Evaluadores[i].Nombres_evaluador)
		}
		if u.Evaluadores[i].Apellidos_evaluador != nil {
			*u.Evaluadores[i].Apellidos_evaluador = strings.TrimSpace(*u.Evaluadores[i].Apellidos_evaluador)
			*u.Evaluadores[i].Apellidos_evaluador = strings.ToUpper(*u.Evaluadores[i].Apellidos_evaluador)
			*u.Evaluadores[i].Apellidos_evaluador = utils.RemoveAccents(*u.Evaluadores[i].Apellidos_evaluador)
		}
		if u.Evaluadores[i].Correo_electronico_evaluador != nil {
			*u.Evaluadores[i].Correo_electronico_evaluador = strings.TrimSpace(*u.Evaluadores[i].Correo_electronico_evaluador)
			*u.Evaluadores[i].Correo_electronico_evaluador = strings.ToUpper(*u.Evaluadores[i].Correo_electronico_evaluador)
			*u.Evaluadores[i].Correo_electronico_evaluador = utils.RemoveAccents(*u.Evaluadores[i].Correo_electronico_evaluador)
		}
		if u.Evaluadores[i].Telefono_fijo_evaluador != nil {
			*u.Evaluadores[i].Telefono_fijo_evaluador = strings.TrimSpace(*u.Evaluadores[i].Telefono_fijo_evaluador)
			*u.Evaluadores[i].Telefono_fijo_evaluador = strings.ToUpper(*u.Evaluadores[i].Telefono_fijo_evaluador)
			*u.Evaluadores[i].Telefono_fijo_evaluador = utils.RemoveAccents(*u.Evaluadores[i].Telefono_fijo_evaluador)
		}
		if u.Evaluadores[i].Telefono_celular_evaluador != nil {
			*u.Evaluadores[i].Telefono_celular_evaluador = strings.TrimSpace(*u.Evaluadores[i].Telefono_celular_evaluador)
			*u.Evaluadores[i].Telefono_celular_evaluador = strings.ToUpper(*u.Evaluadores[i].Telefono_celular_evaluador)
			*u.Evaluadores[i].Telefono_celular_evaluador = utils.RemoveAccents(*u.Evaluadores[i].Telefono_celular_evaluador)
		}
		if u.Evaluadores[i].Recinto_evaluador != nil {
			*u.Evaluadores[i].Recinto_evaluador = strings.TrimSpace(*u.Evaluadores[i].Recinto_evaluador)
			*u.Evaluadores[i].Recinto_evaluador = strings.ToUpper(*u.Evaluadores[i].Recinto_evaluador)
			*u.Evaluadores[i].Recinto_evaluador = utils.RemoveAccents(*u.Evaluadores[i].Recinto_evaluador)
		}
		if u.Evaluadores[i].Cargo_evaluador != nil {
			*u.Evaluadores[i].Cargo_evaluador = strings.TrimSpace(*u.Evaluadores[i].Cargo_evaluador)
			*u.Evaluadores[i].Cargo_evaluador = strings.ToUpper(*u.Evaluadores[i].Cargo_evaluador)
			*u.Evaluadores[i].Cargo_evaluador = utils.RemoveAccents(*u.Evaluadores[i].Cargo_evaluador)
		}
	}
}

// @Summary Agrega nuevos evaluadores de forma masiva
// @Description Genera nuevos evaluadores con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_evaluador     body    massive_add_input     true        "Evaluador a agregar"
// @Success 200 {object} Swagger.AddNewEvaluadoresSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores/carga-masiva [post]
func AddNewEvaluadores(c *gin.Context) {
	var payload massiveAddRequest
	if err := c.ShouldBind(&payload); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	massiveAddRequestParse(&payload)

	var evaluadores_error []string
	for i := 0; i < len(payload.Evaluadores); i++ {
		evaluador := models.Evaluador{
			Id_rol:                       2,
			Grupos_evaluador:             []models.Grupo{},
			Rut_evaluador:                *payload.Evaluadores[i].Rut_evaluador,
			Nombres_evaluador:            *payload.Evaluadores[i].Nombres_evaluador,
			Apellidos_evaluador:          *payload.Evaluadores[i].Apellidos_evaluador,
			Hash_contrasena_evaluador:    *payload.Evaluadores[i].Hash_contrasena_evaluador,
			Correo_electronico_evaluador: *payload.Evaluadores[i].Correo_electronico_evaluador,
			Telefono_fijo_evaluador:      *payload.Evaluadores[i].Telefono_fijo_evaluador,
			Telefono_celular_evaluador:   *payload.Evaluadores[i].Telefono_celular_evaluador,
			Recinto_evaluador:            *payload.Evaluadores[i].Recinto_evaluador,
			Cargo_evaluador:              *payload.Evaluadores[i].Cargo_evaluador,
		}

		for j := 0; j < len(payload.Evaluadores[i].Id_grupos); j++ {
			var gp models.Grupo
			if err := repositories.GetOneGrupo(&gp, utils.IntToString(*payload.Evaluadores[i].Id_grupos[j])); err == nil {
				evaluador.Grupos_evaluador = append(evaluador.Grupos_evaluador, gp)
			}
		}

		if err := repositories.AddNewEvaluador(&evaluador); err != nil {
			evaluadores_error = append(evaluadores_error, "["+*payload.Evaluadores[i].Rut_evaluador+"] "+*payload.Evaluadores[i].Nombres_evaluador+" "+*payload.Evaluadores[i].Apellidos_evaluador)
		}
	}

	if len(evaluadores_error) > 0 {
		api_helpers.RespondJSON(c, 201, evaluadores_error)
	} else {
		api_helpers.RespondJSON(c, 200, "ok")
	}
}

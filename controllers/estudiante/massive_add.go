package estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type grupo struct {
	Sigla_curso *string `json:"sigla_curso"`
	Periodo     *string `json:"periodo"`
	Sigla_grupo *string `json:"sigla_grupo"`
}

type massiveAddRequest struct {
	Estudiantes []struct {
		Grupos                        []grupo `json:"grupos"`
		Rut_estudiante                *string `json:"rut_estudiante"`
		Nombres_estudiante            *string `json:"nombres_estudiante"`
		Apellidos_estudiante          *string `json:"apellidos_estudiante"`
		Hash_contrasena_estudiante    *string `json:"hash_contrasena_estudiante"`
		Correo_electronico_estudiante *string `json:"correo_electronico_estudiante"`
		Telefono_fijo_estudiante      *string `json:"telefono_fijo_estudiante"`
		Telefono_celular_estudiante   *string `json:"telefono_celular_estudiante"`
	} `json:"estudiantes"`
}

func massiveAddRequestFormat(u *massiveAddRequest) {
	for i := 0; i < len(u.Estudiantes); i++ {
		if u.Estudiantes[i].Rut_estudiante != nil {
			*u.Estudiantes[i].Rut_estudiante = strings.TrimSpace(*u.Estudiantes[i].Rut_estudiante)
			*u.Estudiantes[i].Rut_estudiante = strings.ToUpper(*u.Estudiantes[i].Rut_estudiante)
			*u.Estudiantes[i].Rut_estudiante = utils.RemoveAccents(*u.Estudiantes[i].Rut_estudiante)
		}
		if u.Estudiantes[i].Nombres_estudiante != nil {
			*u.Estudiantes[i].Nombres_estudiante = strings.TrimSpace(*u.Estudiantes[i].Nombres_estudiante)
			*u.Estudiantes[i].Nombres_estudiante = strings.ToUpper(*u.Estudiantes[i].Nombres_estudiante)
			*u.Estudiantes[i].Nombres_estudiante = utils.RemoveAccents(*u.Estudiantes[i].Nombres_estudiante)
		}
		if u.Estudiantes[i].Apellidos_estudiante != nil {
			*u.Estudiantes[i].Apellidos_estudiante = strings.TrimSpace(*u.Estudiantes[i].Apellidos_estudiante)
			*u.Estudiantes[i].Apellidos_estudiante = strings.ToUpper(*u.Estudiantes[i].Apellidos_estudiante)
			*u.Estudiantes[i].Apellidos_estudiante = utils.RemoveAccents(*u.Estudiantes[i].Apellidos_estudiante)
		}
		if u.Estudiantes[i].Correo_electronico_estudiante != nil {
			*u.Estudiantes[i].Correo_electronico_estudiante = strings.TrimSpace(*u.Estudiantes[i].Correo_electronico_estudiante)
			*u.Estudiantes[i].Correo_electronico_estudiante = strings.ToUpper(*u.Estudiantes[i].Correo_electronico_estudiante)
			*u.Estudiantes[i].Correo_electronico_estudiante = utils.RemoveAccents(*u.Estudiantes[i].Correo_electronico_estudiante)
		}
		if u.Estudiantes[i].Telefono_fijo_estudiante != nil {
			*u.Estudiantes[i].Telefono_fijo_estudiante = strings.TrimSpace(*u.Estudiantes[i].Telefono_fijo_estudiante)
			*u.Estudiantes[i].Telefono_fijo_estudiante = strings.ToUpper(*u.Estudiantes[i].Telefono_fijo_estudiante)
			*u.Estudiantes[i].Telefono_fijo_estudiante = utils.RemoveAccents(*u.Estudiantes[i].Telefono_fijo_estudiante)
		}
		if u.Estudiantes[i].Telefono_celular_estudiante != nil {
			*u.Estudiantes[i].Telefono_celular_estudiante = strings.TrimSpace(*u.Estudiantes[i].Telefono_celular_estudiante)
			*u.Estudiantes[i].Telefono_celular_estudiante = strings.ToUpper(*u.Estudiantes[i].Telefono_celular_estudiante)
			*u.Estudiantes[i].Telefono_celular_estudiante = utils.RemoveAccents(*u.Estudiantes[i].Telefono_celular_estudiante)
		}
	}
}

// @Summary Agrega nuevos estudiantes de forma masiva
// @Description Genera nuevos estudiantes con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_estudiante     body    massiveAddRequest     true        "Estudiante a agregar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/estudiantes/carga-masiva [post]
func MassiveAdd(c *gin.Context) {
	var payload massiveAddRequest
	if err := c.ShouldBind(&payload); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	massiveAddRequestFormat(&payload)

	var estudiantes_error []string
	for i := 0; i < len(payload.Estudiantes); i++ {
		estudiante := models.Estudiante{
			Id_rol:                        1,
			Grupos_estudiante:             []models.Grupo{},
			Rut_estudiante:                *payload.Estudiantes[i].Rut_estudiante,
			Nombres_estudiante:            *payload.Estudiantes[i].Nombres_estudiante,
			Apellidos_estudiante:          *payload.Estudiantes[i].Apellidos_estudiante,
			Hash_contrasena_estudiante:    *payload.Estudiantes[i].Hash_contrasena_estudiante,
			Correo_electronico_estudiante: *payload.Estudiantes[i].Correo_electronico_estudiante,
			Telefono_fijo_estudiante:      *payload.Estudiantes[i].Telefono_fijo_estudiante,
			Telefono_celular_estudiante:   *payload.Estudiantes[i].Telefono_celular_estudiante,
		}

		for j := 0; j < len(payload.Estudiantes[i].Grupos); j++ {
			var gp models.Grupo
			if err := repositories.GetOneGrupo(&gp, *payload.Estudiantes[i].Grupos[j].Sigla_curso, *payload.Estudiantes[i].Grupos[j].Periodo, *payload.Estudiantes[i].Grupos[j].Sigla_grupo); err == nil {
				estudiante.Grupos_estudiante = append(estudiante.Grupos_estudiante, gp)
			}
		}

		if err := repositories.AddNewEstudiante(&estudiante); err != nil {
			estudiantes_error = append(estudiantes_error, "["+*payload.Estudiantes[i].Rut_estudiante+"] "+*payload.Estudiantes[i].Nombres_estudiante+" "+*payload.Estudiantes[i].Apellidos_estudiante)
		}
	}

	if len(estudiantes_error) > 0 {
		api_helpers.RespondJson(c, 201, estudiantes_error)
	} else {
		api_helpers.RespondJson(c, 200, "OK")
	}
}

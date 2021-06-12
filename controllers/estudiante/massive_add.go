package estudiante

import (
	"fmt"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type Grupo struct {
	Id_periodo  *string `json:"id_periodo"`
	Sigla_curso *string `json:"sigla_curso"`
	Sigla_grupo *string `json:"sigla_grupo"`
}

type massiveAddRequest struct {
	Estudiantes []struct {
		Grupo                         Grupo   `json:"grupo"`
		Rut_estudiante                *string `json:"rut_estudiante"`
		Nombres_estudiante            *string `json:"nombres_estudiante"`
		Apellidos_estudiante          *string `json:"apellidos_estudiante"`
		Correo_electronico_estudiante *string `json:"correo_electronico_estudiante"`
		Telefono_fijo_estudiante      *string `json:"telefono_fijo_estudiante"`
		Telefono_celular_estudiante   *string `json:"telefono_celular_estudiante"`
	} `json:"estudiantes"`
}

func massiveAddRequestFormat(u *massiveAddRequest) {
	for i := 0; i < len(u.Estudiantes); i++ {
		if u.Estudiantes[i].Grupo.Sigla_curso != nil {
			*u.Estudiantes[i].Grupo.Sigla_curso = strings.TrimSpace(*u.Estudiantes[i].Grupo.Sigla_curso)
			*u.Estudiantes[i].Grupo.Sigla_curso = strings.ToUpper(*u.Estudiantes[i].Grupo.Sigla_curso)
			*u.Estudiantes[i].Grupo.Sigla_curso = utils.RemoveAccents(*u.Estudiantes[i].Grupo.Sigla_curso)
		}
		if u.Estudiantes[i].Grupo.Sigla_grupo != nil {
			*u.Estudiantes[i].Grupo.Sigla_grupo = strings.TrimSpace(*u.Estudiantes[i].Grupo.Sigla_grupo)
			*u.Estudiantes[i].Grupo.Sigla_grupo = strings.ToUpper(*u.Estudiantes[i].Grupo.Sigla_grupo)
			*u.Estudiantes[i].Grupo.Sigla_grupo = utils.RemoveAccents(*u.Estudiantes[i].Grupo.Sigla_grupo)
		}
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
		var estudiante models.Estudiante

		if err := repositories.GetOneEstudianteByRut(&estudiante, *payload.Estudiantes[i].Rut_estudiante); err == nil {
			// estudiante ya existe

			// buscamos el grupo y lo asociamos al estudiante
			var gp models.Grupo
			if err := repositories.GetOneGrupo(&gp, *payload.Estudiantes[i].Grupo.Sigla_curso, *payload.Estudiantes[i].Grupo.Id_periodo, *payload.Estudiantes[i].Grupo.Sigla_grupo); err == nil {
				estudiante.Grupos_estudiante = append(estudiante.Grupos_estudiante, gp)
			}

			// actualizamos al estudiante con la nueva lista de grupos
			if err := repositories.PutOneEstudiante(&estudiante, estudiante.Id.String()); err != nil {
				estudiantes_error = append(estudiantes_error, fmt.Sprintf("[%s] %s %s", *payload.Estudiantes[i].Rut_estudiante, *payload.Estudiantes[i].Nombres_estudiante, *payload.Estudiantes[i].Apellidos_estudiante))
			}

		} else {
			// estudiante no existe
			estudiante = models.Estudiante{
				Id_rol:                        1,
				Grupos_estudiante:             []models.Grupo{},
				Rut_estudiante:                *payload.Estudiantes[i].Rut_estudiante,
				Nombres_estudiante:            *payload.Estudiantes[i].Nombres_estudiante,
				Apellidos_estudiante:          *payload.Estudiantes[i].Apellidos_estudiante,
				Hash_contrasena_estudiante:    utils.GeneratePassword(*payload.Estudiantes[i].Rut_estudiante),
				Correo_electronico_estudiante: *payload.Estudiantes[i].Correo_electronico_estudiante,
				Telefono_fijo_estudiante:      *payload.Estudiantes[i].Telefono_fijo_estudiante,
				Telefono_celular_estudiante:   *payload.Estudiantes[i].Telefono_celular_estudiante,
			}

			// buscamos el grupo y lo asociamos al nuevo estudiante
			var gp models.Grupo
			if e_err := repositories.GetOneGrupo(&gp, *payload.Estudiantes[i].Grupo.Sigla_curso, *payload.Estudiantes[i].Grupo.Id_periodo, *payload.Estudiantes[i].Grupo.Sigla_grupo); e_err != nil {
				estudiantes_error = append(estudiantes_error, fmt.Sprintf("[%s] %s %s", *payload.Estudiantes[i].Rut_estudiante, *payload.Estudiantes[i].Nombres_estudiante, *payload.Estudiantes[i].Apellidos_estudiante))
				continue
			}

			estudiante.Grupos_estudiante = append(estudiante.Grupos_estudiante, gp)

			// agregamos el estudiante a la db
			if e_err := repositories.AddNewEstudiante(&estudiante); e_err != nil {
				estudiantes_error = append(estudiantes_error, fmt.Sprintf("[%s] %s %s", *payload.Estudiantes[i].Rut_estudiante, *payload.Estudiantes[i].Nombres_estudiante, *payload.Estudiantes[i].Apellidos_estudiante))
			}
		}
	}

	if len(estudiantes_error) > 0 {
		api_helpers.RespondJson(c, 201, estudiantes_error)
	} else {
		api_helpers.RespondJson(c, 200, "OK")
	}
}

package administrador_academico

import (
	"fmt"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type Curso struct {
	Id_periodo  *string `json:"id_periodo"`
	Sigla_curso *string `json:"sigla_curso"`
}

type massiveAddRequest struct {
	AdministradoresAcademicos []struct {
		Curso                                      Curso   `json:"curso"`
		Rut_administrador_academico                *string `json:"rut_administrador_academico"`
		Nombres_administrador_academico            *string `json:"nombres_administrador_academico"`
		Apellidos_administrador_academico          *string `json:"apellidos_administrador_academico"`
		Hash_contrasena_administrador_academico    *string `json:"hash_contrasena_administrador_academico"`
		Correo_electronico_administrador_academico *string `json:"correo_electronico_administrador_academico"`
		Telefono_fijo_administrador_academico      *string `json:"telefono_fijo_administrador_academico"`
		Telefono_celular_administrador_academico   *string `json:"telefono_celular_administrador_academico"`
	} `json:"administradores_academicos"`
}

func massiveAddRequestParse(u *massiveAddRequest) {
	for i := 0; i < len(u.AdministradoresAcademicos); i++ {
		if u.AdministradoresAcademicos[i].Rut_administrador_academico != nil {
			*u.AdministradoresAcademicos[i].Rut_administrador_academico = strings.TrimSpace(*u.AdministradoresAcademicos[i].Rut_administrador_academico)
			*u.AdministradoresAcademicos[i].Rut_administrador_academico = strings.ToUpper(*u.AdministradoresAcademicos[i].Rut_administrador_academico)
			*u.AdministradoresAcademicos[i].Rut_administrador_academico = utils.RemoveAccents(*u.AdministradoresAcademicos[i].Rut_administrador_academico)
		}
		if u.AdministradoresAcademicos[i].Nombres_administrador_academico != nil {
			*u.AdministradoresAcademicos[i].Nombres_administrador_academico = strings.TrimSpace(*u.AdministradoresAcademicos[i].Nombres_administrador_academico)
			*u.AdministradoresAcademicos[i].Nombres_administrador_academico = strings.ToUpper(*u.AdministradoresAcademicos[i].Nombres_administrador_academico)
			*u.AdministradoresAcademicos[i].Nombres_administrador_academico = utils.RemoveAccents(*u.AdministradoresAcademicos[i].Nombres_administrador_academico)
		}
		if u.AdministradoresAcademicos[i].Apellidos_administrador_academico != nil {
			*u.AdministradoresAcademicos[i].Apellidos_administrador_academico = strings.TrimSpace(*u.AdministradoresAcademicos[i].Apellidos_administrador_academico)
			*u.AdministradoresAcademicos[i].Apellidos_administrador_academico = strings.ToUpper(*u.AdministradoresAcademicos[i].Apellidos_administrador_academico)
			*u.AdministradoresAcademicos[i].Apellidos_administrador_academico = utils.RemoveAccents(*u.AdministradoresAcademicos[i].Apellidos_administrador_academico)
		}
		if u.AdministradoresAcademicos[i].Correo_electronico_administrador_academico != nil {
			*u.AdministradoresAcademicos[i].Correo_electronico_administrador_academico = strings.TrimSpace(*u.AdministradoresAcademicos[i].Correo_electronico_administrador_academico)
			*u.AdministradoresAcademicos[i].Correo_electronico_administrador_academico = strings.ToUpper(*u.AdministradoresAcademicos[i].Correo_electronico_administrador_academico)
			*u.AdministradoresAcademicos[i].Correo_electronico_administrador_academico = utils.RemoveAccents(*u.AdministradoresAcademicos[i].Correo_electronico_administrador_academico)
		}
		if u.AdministradoresAcademicos[i].Telefono_fijo_administrador_academico != nil {
			*u.AdministradoresAcademicos[i].Telefono_fijo_administrador_academico = strings.TrimSpace(*u.AdministradoresAcademicos[i].Telefono_fijo_administrador_academico)
			*u.AdministradoresAcademicos[i].Telefono_fijo_administrador_academico = strings.ToUpper(*u.AdministradoresAcademicos[i].Telefono_fijo_administrador_academico)
			*u.AdministradoresAcademicos[i].Telefono_fijo_administrador_academico = utils.RemoveAccents(*u.AdministradoresAcademicos[i].Telefono_fijo_administrador_academico)
		}
		if u.AdministradoresAcademicos[i].Telefono_celular_administrador_academico != nil {
			*u.AdministradoresAcademicos[i].Telefono_celular_administrador_academico = strings.TrimSpace(*u.AdministradoresAcademicos[i].Telefono_celular_administrador_academico)
			*u.AdministradoresAcademicos[i].Telefono_celular_administrador_academico = strings.ToUpper(*u.AdministradoresAcademicos[i].Telefono_celular_administrador_academico)
			*u.AdministradoresAcademicos[i].Telefono_celular_administrador_academico = utils.RemoveAccents(*u.AdministradoresAcademicos[i].Telefono_celular_administrador_academico)
		}
	}
}

// @Summary Agrega nuevos administrador_academicos de forma masiva
// @Description Genera nuevos administrador_academicos con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_administrador_academico     body    massiveAddRequest     true        "AdministradoresAcademico a agregar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/administradores-academicos/carga-masiva [post]
func MassiveAdd(c *gin.Context) {
	var payload massiveAddRequest
	if err := c.ShouldBind(&payload); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	massiveAddRequestParse(&payload)

	var administradores_academicos_error []string
	for i := 0; i < len(payload.AdministradoresAcademicos); i++ {
		var administrador_academico models.AdministradorAcademico

		if err := repositories.GetOneAdministradorAcademicoByRut(&administrador_academico, *payload.AdministradoresAcademicos[i].Rut_administrador_academico); err == nil {
			// administrador_academico ya existe

			// buscamos el grupo y lo asociamos al administrador_academico
			var curso models.Curso
			if err := repositories.GetOneCurso(&curso, *payload.AdministradoresAcademicos[i].Curso.Sigla_curso, *payload.AdministradoresAcademicos[i].Curso.Id_periodo); err == nil {
				administrador_academico.Cursos_administrador_academico = append(administrador_academico.Cursos_administrador_academico, curso)
			}

			// actualizamos al administrador_academico con la nueva lista de grupos
			if err := repositories.PutAdministradorAcademico(&administrador_academico, administrador_academico.Id.String()); err != nil {
				administradores_academicos_error = append(administradores_academicos_error, fmt.Sprintf("[%s] %s %s", *payload.AdministradoresAcademicos[i].Rut_administrador_academico, *payload.AdministradoresAcademicos[i].Nombres_administrador_academico, *payload.AdministradoresAcademicos[i].Apellidos_administrador_academico))
			}

		} else {
			// administrador_academico no existe
			administrador_academico = models.AdministradorAcademico{
				Id_rol:                                     1,
				Cursos_administrador_academico:             []models.Curso{},
				Rut_administrador_academico:                *payload.AdministradoresAcademicos[i].Rut_administrador_academico,
				Nombres_administrador_academico:            *payload.AdministradoresAcademicos[i].Nombres_administrador_academico,
				Apellidos_administrador_academico:          *payload.AdministradoresAcademicos[i].Apellidos_administrador_academico,
				Hash_contrasena_administrador_academico:    *payload.AdministradoresAcademicos[i].Hash_contrasena_administrador_academico,
				Correo_electronico_administrador_academico: *payload.AdministradoresAcademicos[i].Correo_electronico_administrador_academico,
				Telefono_fijo_administrador_academico:      *payload.AdministradoresAcademicos[i].Telefono_fijo_administrador_academico,
				Telefono_celular_administrador_academico:   *payload.AdministradoresAcademicos[i].Telefono_celular_administrador_academico,
			}

			// buscamos el grupo y lo asociamos al nuevo administrador_academico
			var curso models.Curso
			if err := repositories.GetOneCurso(&curso, *payload.AdministradoresAcademicos[i].Curso.Sigla_curso, *payload.AdministradoresAcademicos[i].Curso.Id_periodo); err == nil {
				administrador_academico.Cursos_administrador_academico = append(administrador_academico.Cursos_administrador_academico, curso)
			}

			// agregamos el administrador_academico a la db
			if err := repositories.AddAdministradorAcademico(&administrador_academico); err != nil {
				administradores_academicos_error = append(administradores_academicos_error, fmt.Sprintf("[%s] %s %s", *payload.AdministradoresAcademicos[i].Rut_administrador_academico, *payload.AdministradoresAcademicos[i].Nombres_administrador_academico, *payload.AdministradoresAcademicos[i].Apellidos_administrador_academico))
			}
		}
	}

	// for i := 0; i < len(payload.AdministradoresAcademicos); i++ {
	// 	administrador_academico := models.AdministradorAcademico{
	// 		Id_rol:                                     3,
	// 		Cursos_administrador_academico:             []models.Curso{},
	// 		Rut_administrador_academico:                *payload.AdministradoresAcademicos[i].Rut_administrador_academico,
	// 		Nombres_administrador_academico:            *payload.AdministradoresAcademicos[i].Nombres_administrador_academico,
	// 		Apellidos_administrador_academico:          *payload.AdministradoresAcademicos[i].Apellidos_administrador_academico,
	// 		Hash_contrasena_administrador_academico:    *payload.AdministradoresAcademicos[i].Hash_contrasena_administrador_academico,
	// 		Correo_electronico_administrador_academico: *payload.AdministradoresAcademicos[i].Correo_electronico_administrador_academico,
	// 		Telefono_fijo_administrador_academico:      *payload.AdministradoresAcademicos[i].Telefono_fijo_administrador_academico,
	// 		Telefono_celular_administrador_academico:   *payload.AdministradoresAcademicos[i].Telefono_celular_administrador_academico,
	// 	}

	// 	for j := 0; j < len(payload.AdministradoresAcademicos[i].Cursos); j++ {
	// 		var crs models.Curso

	// 		if err := repositories.GetOneCurso(&crs, *payload.AdministradoresAcademicos[i].Cursos[j].Sigla, *payload.AdministradoresAcademicos[i].Cursos[j].Periodo); err == nil {
	// 			administrador_academico.Cursos_administrador_academico = append(administrador_academico.Cursos_administrador_academico, crs)
	// 		}
	// 	}

	// 	if err := repositories.AddAdministradorAcademico(&administrador_academico); err != nil {
	// 		administradores_academicos_error = append(administradores_academicos_error, "["+*payload.AdministradoresAcademicos[i].Rut_administrador_academico+"] "+*payload.AdministradoresAcademicos[i].Nombres_administrador_academico+" "+*payload.AdministradoresAcademicos[i].Apellidos_administrador_academico)
	// 	}
	// }

	if len(administradores_academicos_error) > 0 {
		api_helpers.RespondJson(c, 201, administradores_academicos_error)
	} else {
		api_helpers.RespondJson(c, 200, "OK")
	}
}

package administrador_academico

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type massiveAddRequest struct {
	AdministradoresAcademicos []struct {
		Id_cursos                                  []*int  `json:"id_cursos"`
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
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_administrador_academico     body    massive_add_input     true        "AdministradoresAcademico a agregar"
// @Success 200 {object} massive_add_input "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
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
		administrador_academico := models.AdministradorAcademico{
			Id_rol:                                     3,
			Cursos_administrador_academico:             []models.Curso{},
			Rut_administrador_academico:                *payload.AdministradoresAcademicos[i].Rut_administrador_academico,
			Nombres_administrador_academico:            *payload.AdministradoresAcademicos[i].Nombres_administrador_academico,
			Apellidos_administrador_academico:          *payload.AdministradoresAcademicos[i].Apellidos_administrador_academico,
			Hash_contrasena_administrador_academico:    *payload.AdministradoresAcademicos[i].Hash_contrasena_administrador_academico,
			Correo_electronico_administrador_academico: *payload.AdministradoresAcademicos[i].Correo_electronico_administrador_academico,
			Telefono_fijo_administrador_academico:      *payload.AdministradoresAcademicos[i].Telefono_fijo_administrador_academico,
			Telefono_celular_administrador_academico:   *payload.AdministradoresAcademicos[i].Telefono_celular_administrador_academico,
		}

		for j := 0; j < len(payload.AdministradoresAcademicos[i].Id_cursos); j++ {
			var crs models.Curso
			if err := repositories.GetOneCurso(&crs, utils.ConvertIntToString(*payload.AdministradoresAcademicos[i].Id_cursos[j])); err == nil {
				administrador_academico.Cursos_administrador_academico = append(administrador_academico.Cursos_administrador_academico, crs)
			}
		}

		if err := repositories.AddAdministradorAcademico(&administrador_academico); err != nil {
			administradores_academicos_error = append(administradores_academicos_error, "["+*payload.AdministradoresAcademicos[i].Rut_administrador_academico+"] "+*payload.AdministradoresAcademicos[i].Nombres_administrador_academico+" "+*payload.AdministradoresAcademicos[i].Apellidos_administrador_academico)
		}
	}

	if len(administradores_academicos_error) > 0 {
		api_helpers.RespondJSON(c, 201, administradores_academicos_error)
	} else {
		api_helpers.RespondJSON(c, 200, "OK")
	}
}

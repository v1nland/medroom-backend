package f_input

import (
	"medroom-backend/messages/Request"
	"medroom-backend/utils"
	"strings"
)

func ListEstudiantes(u *Request.ListEstudiantes) {
}

func GetOneEstudiante(u *Request.GetOneEstudiante) {
}

func AddNewEstudiante(u *Request.AddNewEstudiante) {
	if u.Rut_estudiante != nil {
		*u.Rut_estudiante = strings.TrimSpace(*u.Rut_estudiante)
		*u.Rut_estudiante = strings.ToUpper(*u.Rut_estudiante)
		*u.Rut_estudiante = utils.RemoveAccents(*u.Rut_estudiante)
	}
	if u.Nombres_estudiante != nil {
		*u.Nombres_estudiante = strings.TrimSpace(*u.Nombres_estudiante)
		*u.Nombres_estudiante = strings.ToUpper(*u.Nombres_estudiante)
		*u.Nombres_estudiante = utils.RemoveAccents(*u.Nombres_estudiante)
	}
	if u.Apellidos_estudiante != nil {
		*u.Apellidos_estudiante = strings.TrimSpace(*u.Apellidos_estudiante)
		*u.Apellidos_estudiante = strings.ToUpper(*u.Apellidos_estudiante)
		*u.Apellidos_estudiante = utils.RemoveAccents(*u.Apellidos_estudiante)
	}
	if u.Correo_electronico_estudiante != nil {
		*u.Correo_electronico_estudiante = strings.TrimSpace(*u.Correo_electronico_estudiante)
		*u.Correo_electronico_estudiante = strings.ToUpper(*u.Correo_electronico_estudiante)
		*u.Correo_electronico_estudiante = utils.RemoveAccents(*u.Correo_electronico_estudiante)
	}
	if u.Telefono_fijo_estudiante != nil {
		*u.Telefono_fijo_estudiante = strings.TrimSpace(*u.Telefono_fijo_estudiante)
		*u.Telefono_fijo_estudiante = strings.ToUpper(*u.Telefono_fijo_estudiante)
		*u.Telefono_fijo_estudiante = utils.RemoveAccents(*u.Telefono_fijo_estudiante)
	}
	if u.Telefono_celular_estudiante != nil {
		*u.Telefono_celular_estudiante = strings.TrimSpace(*u.Telefono_celular_estudiante)
		*u.Telefono_celular_estudiante = strings.ToUpper(*u.Telefono_celular_estudiante)
		*u.Telefono_celular_estudiante = utils.RemoveAccents(*u.Telefono_celular_estudiante)
	}
}

func PutOneEstudiante(u *Request.PutOneEstudiante) {
	if u.Rut_estudiante != nil {
		*u.Rut_estudiante = strings.TrimSpace(*u.Rut_estudiante)
		*u.Rut_estudiante = strings.ToUpper(*u.Rut_estudiante)
		*u.Rut_estudiante = utils.RemoveAccents(*u.Rut_estudiante)
	}
	if u.Nombres_estudiante != nil {
		*u.Nombres_estudiante = strings.TrimSpace(*u.Nombres_estudiante)
		*u.Nombres_estudiante = strings.ToUpper(*u.Nombres_estudiante)
		*u.Nombres_estudiante = utils.RemoveAccents(*u.Nombres_estudiante)
	}
	if u.Apellidos_estudiante != nil {
		*u.Apellidos_estudiante = strings.TrimSpace(*u.Apellidos_estudiante)
		*u.Apellidos_estudiante = strings.ToUpper(*u.Apellidos_estudiante)
		*u.Apellidos_estudiante = utils.RemoveAccents(*u.Apellidos_estudiante)
	}
	if u.Correo_electronico_estudiante != nil {
		*u.Correo_electronico_estudiante = strings.TrimSpace(*u.Correo_electronico_estudiante)
		*u.Correo_electronico_estudiante = strings.ToUpper(*u.Correo_electronico_estudiante)
		*u.Correo_electronico_estudiante = utils.RemoveAccents(*u.Correo_electronico_estudiante)
	}
	if u.Telefono_fijo_estudiante != nil {
		*u.Telefono_fijo_estudiante = strings.TrimSpace(*u.Telefono_fijo_estudiante)
		*u.Telefono_fijo_estudiante = strings.ToUpper(*u.Telefono_fijo_estudiante)
		*u.Telefono_fijo_estudiante = utils.RemoveAccents(*u.Telefono_fijo_estudiante)
	}
	if u.Telefono_celular_estudiante != nil {
		*u.Telefono_celular_estudiante = strings.TrimSpace(*u.Telefono_celular_estudiante)
		*u.Telefono_celular_estudiante = strings.ToUpper(*u.Telefono_celular_estudiante)
		*u.Telefono_celular_estudiante = utils.RemoveAccents(*u.Telefono_celular_estudiante)
	}
}

func DeleteEstudiante(u *Request.DeleteEstudiante) {
}

func GetMyEstudiante(u *Request.GetMyEstudiante) {
}

func PutMyEstudiante(u *Request.PutMyEstudiante) {
	if u.Telefono_fijo_estudiante != nil {
		*u.Telefono_fijo_estudiante = strings.TrimSpace(*u.Telefono_fijo_estudiante)
		*u.Telefono_fijo_estudiante = strings.ToUpper(*u.Telefono_fijo_estudiante)
		*u.Telefono_fijo_estudiante = utils.RemoveAccents(*u.Telefono_fijo_estudiante)
	}
	if u.Telefono_celular_estudiante != nil {
		*u.Telefono_celular_estudiante = strings.TrimSpace(*u.Telefono_celular_estudiante)
		*u.Telefono_celular_estudiante = strings.ToUpper(*u.Telefono_celular_estudiante)
		*u.Telefono_celular_estudiante = utils.RemoveAccents(*u.Telefono_celular_estudiante)
	}
}

func AddEstudianteToGrupo(u *Request.AddEstudianteToGrupo) {
}

func AddEstudianteToCurso(u *Request.AddEstudianteToCurso) {
}

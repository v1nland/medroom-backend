package f_input

import (
	"medroom-backend/app/messages/Request"
	"medroom-backend/app/utils"
	"strings"
)

func ListAdministradoresAcademicos(u *Request.ListAdministradoresAcademicos) {

}

func GetOneAdministradorAcademico(u *Request.GetOneAdministradorAcademico) {

}

func GetMyAdministradorAcademico(u *Request.GetMyAdministradorAcademico) {

}

func AddNewAdministradorAcademico(u *Request.AddNewAdministradorAcademico) {
	if u.Rut_administrador_academico != nil {
		*u.Rut_administrador_academico = strings.TrimSpace(*u.Rut_administrador_academico)
		*u.Rut_administrador_academico = strings.ToUpper(*u.Rut_administrador_academico)
		*u.Rut_administrador_academico = utils.RemoveAccents(*u.Rut_administrador_academico)
	}

	if u.Nombres_administrador_academico != nil {
		*u.Nombres_administrador_academico = strings.TrimSpace(*u.Nombres_administrador_academico)
		*u.Nombres_administrador_academico = strings.ToUpper(*u.Nombres_administrador_academico)
		*u.Nombres_administrador_academico = utils.RemoveAccents(*u.Nombres_administrador_academico)
	}

	if u.Apellidos_administrador_academico != nil {
		*u.Apellidos_administrador_academico = strings.TrimSpace(*u.Apellidos_administrador_academico)
		*u.Apellidos_administrador_academico = strings.ToUpper(*u.Apellidos_administrador_academico)
		*u.Apellidos_administrador_academico = utils.RemoveAccents(*u.Apellidos_administrador_academico)
	}

	if u.Correo_electronico_administrador_academico != nil {
		*u.Correo_electronico_administrador_academico = strings.TrimSpace(*u.Correo_electronico_administrador_academico)
		*u.Correo_electronico_administrador_academico = strings.ToUpper(*u.Correo_electronico_administrador_academico)
		*u.Correo_electronico_administrador_academico = utils.RemoveAccents(*u.Correo_electronico_administrador_academico)
	}

	if u.Telefono_fijo_administrador_academico != nil {
		*u.Telefono_fijo_administrador_academico = strings.TrimSpace(*u.Telefono_fijo_administrador_academico)
		*u.Telefono_fijo_administrador_academico = strings.ToUpper(*u.Telefono_fijo_administrador_academico)
		*u.Telefono_fijo_administrador_academico = utils.RemoveAccents(*u.Telefono_fijo_administrador_academico)
	}

	if u.Telefono_celular_administrador_academico != nil {
		*u.Telefono_celular_administrador_academico = strings.TrimSpace(*u.Telefono_celular_administrador_academico)
		*u.Telefono_celular_administrador_academico = strings.ToUpper(*u.Telefono_celular_administrador_academico)
		*u.Telefono_celular_administrador_academico = utils.RemoveAccents(*u.Telefono_celular_administrador_academico)
	}
}

func PutOneAdministradorAcademico(u *Request.PutOneAdministradorAcademico) {
	if u.Rut_administrador_academico != nil {
		*u.Rut_administrador_academico = strings.TrimSpace(*u.Rut_administrador_academico)
		*u.Rut_administrador_academico = strings.ToUpper(*u.Rut_administrador_academico)
		*u.Rut_administrador_academico = utils.RemoveAccents(*u.Rut_administrador_academico)
	}

	if u.Nombres_administrador_academico != nil {
		*u.Nombres_administrador_academico = strings.TrimSpace(*u.Nombres_administrador_academico)
		*u.Nombres_administrador_academico = strings.ToUpper(*u.Nombres_administrador_academico)
		*u.Nombres_administrador_academico = utils.RemoveAccents(*u.Nombres_administrador_academico)
	}

	if u.Apellidos_administrador_academico != nil {
		*u.Apellidos_administrador_academico = strings.TrimSpace(*u.Apellidos_administrador_academico)
		*u.Apellidos_administrador_academico = strings.ToUpper(*u.Apellidos_administrador_academico)
		*u.Apellidos_administrador_academico = utils.RemoveAccents(*u.Apellidos_administrador_academico)
	}

	if u.Correo_electronico_administrador_academico != nil {
		*u.Correo_electronico_administrador_academico = strings.TrimSpace(*u.Correo_electronico_administrador_academico)
		*u.Correo_electronico_administrador_academico = strings.ToUpper(*u.Correo_electronico_administrador_academico)
		*u.Correo_electronico_administrador_academico = utils.RemoveAccents(*u.Correo_electronico_administrador_academico)
	}

	if u.Telefono_fijo_administrador_academico != nil {
		*u.Telefono_fijo_administrador_academico = strings.TrimSpace(*u.Telefono_fijo_administrador_academico)
		*u.Telefono_fijo_administrador_academico = strings.ToUpper(*u.Telefono_fijo_administrador_academico)
		*u.Telefono_fijo_administrador_academico = utils.RemoveAccents(*u.Telefono_fijo_administrador_academico)
	}

	if u.Telefono_celular_administrador_academico != nil {
		*u.Telefono_celular_administrador_academico = strings.TrimSpace(*u.Telefono_celular_administrador_academico)
		*u.Telefono_celular_administrador_academico = strings.ToUpper(*u.Telefono_celular_administrador_academico)
		*u.Telefono_celular_administrador_academico = utils.RemoveAccents(*u.Telefono_celular_administrador_academico)
	}
}

func PutMyAdministradorAcademico(u *Request.PutMyAdministradorAcademico) {
	if u.Rut_administrador_academico != nil {
		*u.Rut_administrador_academico = strings.TrimSpace(*u.Rut_administrador_academico)
		*u.Rut_administrador_academico = strings.ToUpper(*u.Rut_administrador_academico)
		*u.Rut_administrador_academico = utils.RemoveAccents(*u.Rut_administrador_academico)
	}

	if u.Nombres_administrador_academico != nil {
		*u.Nombres_administrador_academico = strings.TrimSpace(*u.Nombres_administrador_academico)
		*u.Nombres_administrador_academico = strings.ToUpper(*u.Nombres_administrador_academico)
		*u.Nombres_administrador_academico = utils.RemoveAccents(*u.Nombres_administrador_academico)
	}

	if u.Apellidos_administrador_academico != nil {
		*u.Apellidos_administrador_academico = strings.TrimSpace(*u.Apellidos_administrador_academico)
		*u.Apellidos_administrador_academico = strings.ToUpper(*u.Apellidos_administrador_academico)
		*u.Apellidos_administrador_academico = utils.RemoveAccents(*u.Apellidos_administrador_academico)
	}

	if u.Correo_electronico_administrador_academico != nil {
		*u.Correo_electronico_administrador_academico = strings.TrimSpace(*u.Correo_electronico_administrador_academico)
		*u.Correo_electronico_administrador_academico = strings.ToUpper(*u.Correo_electronico_administrador_academico)
		*u.Correo_electronico_administrador_academico = utils.RemoveAccents(*u.Correo_electronico_administrador_academico)
	}

	if u.Telefono_fijo_administrador_academico != nil {
		*u.Telefono_fijo_administrador_academico = strings.TrimSpace(*u.Telefono_fijo_administrador_academico)
		*u.Telefono_fijo_administrador_academico = strings.ToUpper(*u.Telefono_fijo_administrador_academico)
		*u.Telefono_fijo_administrador_academico = utils.RemoveAccents(*u.Telefono_fijo_administrador_academico)
	}

	if u.Telefono_celular_administrador_academico != nil {
		*u.Telefono_celular_administrador_academico = strings.TrimSpace(*u.Telefono_celular_administrador_academico)
		*u.Telefono_celular_administrador_academico = strings.ToUpper(*u.Telefono_celular_administrador_academico)
		*u.Telefono_celular_administrador_academico = utils.RemoveAccents(*u.Telefono_celular_administrador_academico)
	}
}

func DeleteAdministradorAcademico(u *Request.DeleteAdministradorAcademico) {

}

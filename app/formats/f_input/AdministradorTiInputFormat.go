package f_input

import (
	"medroom-backend/app/messages/Request"
	"medroom-backend/app/utils"
	"strings"
)

func ListAdministradoresTi(u *Request.ListAdministradoresTi) {
}

func GetOneAdministradorTi(u *Request.GetOneAdministradorTi) {
}

func GetMyAdministradorTi(u *Request.GetMyAdministradorTi) {
}

func AddNewAdministradorTi(u *Request.AddNewAdministradorTi) {
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

func PutOneAdministradorTi(u *Request.PutOneAdministradorTi) {
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

func PutMyAdministradorTi(u *Request.PutMyAdministradorTi) {
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

func DeleteAdministradorTi(u *Request.DeleteAdministradorTi) {
}

package InputFormats

import (
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"
	"strings"
)

func GetAdministradoresTiInput(u *RequestMessages.ListAdministradoresTiPayload) {

}

func GetOneAdministradorTiInput(u *RequestMessages.GetOneAdministradorTiPayload) {

}

func AddNewAdministradorTiInput(u *RequestMessages.AddNewAdministradorTiPayload) {
	u.Rut_administrador_ti = strings.TrimSpace(u.Rut_administrador_ti)
	u.Rut_administrador_ti = strings.ToUpper(u.Rut_administrador_ti)
	u.Rut_administrador_ti = Utils.RemoveAccents(u.Rut_administrador_ti)

	u.Nombres_administrador_ti = strings.TrimSpace(u.Nombres_administrador_ti)
	u.Nombres_administrador_ti = strings.ToUpper(u.Nombres_administrador_ti)
	u.Nombres_administrador_ti = Utils.RemoveAccents(u.Nombres_administrador_ti)

	u.Apellidos_administrador_ti = strings.TrimSpace(u.Apellidos_administrador_ti)
	u.Apellidos_administrador_ti = strings.ToUpper(u.Apellidos_administrador_ti)
	u.Apellidos_administrador_ti = Utils.RemoveAccents(u.Apellidos_administrador_ti)

	u.Correo_electronico_administrador_ti = strings.TrimSpace(u.Correo_electronico_administrador_ti)
	u.Correo_electronico_administrador_ti = strings.ToUpper(u.Correo_electronico_administrador_ti)
	u.Correo_electronico_administrador_ti = Utils.RemoveAccents(u.Correo_electronico_administrador_ti)

	u.Telefono_fijo_administrador_ti = strings.TrimSpace(u.Telefono_fijo_administrador_ti)
	u.Telefono_fijo_administrador_ti = strings.ToUpper(u.Telefono_fijo_administrador_ti)
	u.Telefono_fijo_administrador_ti = Utils.RemoveAccents(u.Telefono_fijo_administrador_ti)

	u.Telefono_celular_administrador_ti = strings.TrimSpace(u.Telefono_celular_administrador_ti)
	u.Telefono_celular_administrador_ti = strings.ToUpper(u.Telefono_celular_administrador_ti)
	u.Telefono_celular_administrador_ti = Utils.RemoveAccents(u.Telefono_celular_administrador_ti)
}

func PutOneAdministradorTiInput(u *RequestMessages.PutOneAdministradorTiPayload) {
	u.Rut_administrador_ti = strings.TrimSpace(u.Rut_administrador_ti)
	u.Rut_administrador_ti = strings.ToUpper(u.Rut_administrador_ti)
	u.Rut_administrador_ti = Utils.RemoveAccents(u.Rut_administrador_ti)

	u.Nombres_administrador_ti = strings.TrimSpace(u.Nombres_administrador_ti)
	u.Nombres_administrador_ti = strings.ToUpper(u.Nombres_administrador_ti)
	u.Nombres_administrador_ti = Utils.RemoveAccents(u.Nombres_administrador_ti)

	u.Apellidos_administrador_ti = strings.TrimSpace(u.Apellidos_administrador_ti)
	u.Apellidos_administrador_ti = strings.ToUpper(u.Apellidos_administrador_ti)
	u.Apellidos_administrador_ti = Utils.RemoveAccents(u.Apellidos_administrador_ti)

	u.Correo_electronico_administrador_ti = strings.TrimSpace(u.Correo_electronico_administrador_ti)
	u.Correo_electronico_administrador_ti = strings.ToUpper(u.Correo_electronico_administrador_ti)
	u.Correo_electronico_administrador_ti = Utils.RemoveAccents(u.Correo_electronico_administrador_ti)

	u.Telefono_fijo_administrador_ti = strings.TrimSpace(u.Telefono_fijo_administrador_ti)
	u.Telefono_fijo_administrador_ti = strings.ToUpper(u.Telefono_fijo_administrador_ti)
	u.Telefono_fijo_administrador_ti = Utils.RemoveAccents(u.Telefono_fijo_administrador_ti)

	u.Telefono_celular_administrador_ti = strings.TrimSpace(u.Telefono_celular_administrador_ti)
	u.Telefono_celular_administrador_ti = strings.ToUpper(u.Telefono_celular_administrador_ti)
	u.Telefono_celular_administrador_ti = Utils.RemoveAccents(u.Telefono_celular_administrador_ti)
}

func DeleteAdministradorTiInput(u *RequestMessages.DeleteAdministradorTiPayload) {

}

package InputFormats

import (
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"
	"strings"
)

func GetAdministradoresAcademicosInput(u *RequestMessages.ListAdministradoresAcademicosPayload) {

}

func GetOneAdministradorAcademicoInput(u *RequestMessages.GetOneAdministradorAcademicoPayload) {

}

func GetMyAdministradorAcademicoInput(u *RequestMessages.GetMyAdministradorAcademicoPayload) {

}

func AddNewAdministradorAcademicoInput(u *RequestMessages.AddNewAdministradorAcademicoPayload) {
	u.Rut_administrador_academico = strings.TrimSpace(u.Rut_administrador_academico)
	u.Rut_administrador_academico = strings.ToUpper(u.Rut_administrador_academico)
	u.Rut_administrador_academico = Utils.RemoveAccents(u.Rut_administrador_academico)

	u.Nombres_administrador_academico = strings.TrimSpace(u.Nombres_administrador_academico)
	u.Nombres_administrador_academico = strings.ToUpper(u.Nombres_administrador_academico)
	u.Nombres_administrador_academico = Utils.RemoveAccents(u.Nombres_administrador_academico)

	u.Apellidos_administrador_academico = strings.TrimSpace(u.Apellidos_administrador_academico)
	u.Apellidos_administrador_academico = strings.ToUpper(u.Apellidos_administrador_academico)
	u.Apellidos_administrador_academico = Utils.RemoveAccents(u.Apellidos_administrador_academico)

	u.Correo_electronico_administrador_academico = strings.TrimSpace(u.Correo_electronico_administrador_academico)
	u.Correo_electronico_administrador_academico = strings.ToUpper(u.Correo_electronico_administrador_academico)
	u.Correo_electronico_administrador_academico = Utils.RemoveAccents(u.Correo_electronico_administrador_academico)

	u.Telefono_fijo_administrador_academico = strings.TrimSpace(u.Telefono_fijo_administrador_academico)
	u.Telefono_fijo_administrador_academico = strings.ToUpper(u.Telefono_fijo_administrador_academico)
	u.Telefono_fijo_administrador_academico = Utils.RemoveAccents(u.Telefono_fijo_administrador_academico)

	u.Telefono_celular_administrador_academico = strings.TrimSpace(u.Telefono_celular_administrador_academico)
	u.Telefono_celular_administrador_academico = strings.ToUpper(u.Telefono_celular_administrador_academico)
	u.Telefono_celular_administrador_academico = Utils.RemoveAccents(u.Telefono_celular_administrador_academico)
}

func PutOneAdministradorAcademicoInput(u *RequestMessages.PutOneAdministradorAcademicoPayload) {
	u.Rut_administrador_academico = strings.TrimSpace(u.Rut_administrador_academico)
	u.Rut_administrador_academico = strings.ToUpper(u.Rut_administrador_academico)
	u.Rut_administrador_academico = Utils.RemoveAccents(u.Rut_administrador_academico)

	u.Nombres_administrador_academico = strings.TrimSpace(u.Nombres_administrador_academico)
	u.Nombres_administrador_academico = strings.ToUpper(u.Nombres_administrador_academico)
	u.Nombres_administrador_academico = Utils.RemoveAccents(u.Nombres_administrador_academico)

	u.Apellidos_administrador_academico = strings.TrimSpace(u.Apellidos_administrador_academico)
	u.Apellidos_administrador_academico = strings.ToUpper(u.Apellidos_administrador_academico)
	u.Apellidos_administrador_academico = Utils.RemoveAccents(u.Apellidos_administrador_academico)

	u.Correo_electronico_administrador_academico = strings.TrimSpace(u.Correo_electronico_administrador_academico)
	u.Correo_electronico_administrador_academico = strings.ToUpper(u.Correo_electronico_administrador_academico)
	u.Correo_electronico_administrador_academico = Utils.RemoveAccents(u.Correo_electronico_administrador_academico)

	u.Telefono_fijo_administrador_academico = strings.TrimSpace(u.Telefono_fijo_administrador_academico)
	u.Telefono_fijo_administrador_academico = strings.ToUpper(u.Telefono_fijo_administrador_academico)
	u.Telefono_fijo_administrador_academico = Utils.RemoveAccents(u.Telefono_fijo_administrador_academico)

	u.Telefono_celular_administrador_academico = strings.TrimSpace(u.Telefono_celular_administrador_academico)
	u.Telefono_celular_administrador_academico = strings.ToUpper(u.Telefono_celular_administrador_academico)
	u.Telefono_celular_administrador_academico = Utils.RemoveAccents(u.Telefono_celular_administrador_academico)
}

func PutMyAdministradorAcademicoInput(u *RequestMessages.PutMyAdministradorAcademicoPayload) {
	u.Rut_administrador_academico = strings.TrimSpace(u.Rut_administrador_academico)
	u.Rut_administrador_academico = strings.ToUpper(u.Rut_administrador_academico)
	u.Rut_administrador_academico = Utils.RemoveAccents(u.Rut_administrador_academico)

	u.Nombres_administrador_academico = strings.TrimSpace(u.Nombres_administrador_academico)
	u.Nombres_administrador_academico = strings.ToUpper(u.Nombres_administrador_academico)
	u.Nombres_administrador_academico = Utils.RemoveAccents(u.Nombres_administrador_academico)

	u.Apellidos_administrador_academico = strings.TrimSpace(u.Apellidos_administrador_academico)
	u.Apellidos_administrador_academico = strings.ToUpper(u.Apellidos_administrador_academico)
	u.Apellidos_administrador_academico = Utils.RemoveAccents(u.Apellidos_administrador_academico)

	u.Correo_electronico_administrador_academico = strings.TrimSpace(u.Correo_electronico_administrador_academico)
	u.Correo_electronico_administrador_academico = strings.ToUpper(u.Correo_electronico_administrador_academico)
	u.Correo_electronico_administrador_academico = Utils.RemoveAccents(u.Correo_electronico_administrador_academico)

	u.Telefono_fijo_administrador_academico = strings.TrimSpace(u.Telefono_fijo_administrador_academico)
	u.Telefono_fijo_administrador_academico = strings.ToUpper(u.Telefono_fijo_administrador_academico)
	u.Telefono_fijo_administrador_academico = Utils.RemoveAccents(u.Telefono_fijo_administrador_academico)

	u.Telefono_celular_administrador_academico = strings.TrimSpace(u.Telefono_celular_administrador_academico)
	u.Telefono_celular_administrador_academico = strings.ToUpper(u.Telefono_celular_administrador_academico)
	u.Telefono_celular_administrador_academico = Utils.RemoveAccents(u.Telefono_celular_administrador_academico)
}

func DeleteAdministradorAcademicoInput(u *RequestMessages.DeleteAdministradorAcademicoPayload) {

}

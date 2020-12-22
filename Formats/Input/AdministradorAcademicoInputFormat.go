package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func GetAdministradoresAcademicosInput(u *Request.ListAdministradoresAcademicosPayload) {

}

func GetOneAdministradorAcademicoInput(u *Request.GetOneAdministradorAcademicoPayload) {

}

func GetMyAdministradorAcademicoInput(u *Request.GetMyAdministradorAcademicoPayload) {

}

func AddNewAdministradorAcademicoInput(u *Request.AddNewAdministradorAcademicoPayload) {
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

func PutOneAdministradorAcademicoInput(u *Request.PutOneAdministradorAcademicoPayload) {
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

func PutMyAdministradorAcademicoInput(u *Request.PutMyAdministradorAcademicoPayload) {
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

func DeleteAdministradorAcademicoInput(u *Request.DeleteAdministradorAcademicoPayload) {

}

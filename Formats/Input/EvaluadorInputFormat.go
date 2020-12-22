package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func GetEvaluadoresInput(u *Request.ListEvaluadoresPayload) {

}

func GetOneEvaluadorInput(u *Request.GetOneEvaluadorPayload) {

}

func AddNewEvaluadorInput(u *Request.AddNewEvaluadorPayload) {
	u.Rut_evaluador = strings.TrimSpace(u.Rut_evaluador)
	u.Rut_evaluador = strings.ToUpper(u.Rut_evaluador)
	u.Rut_evaluador = Utils.RemoveAccents(u.Rut_evaluador)

	u.Nombres_evaluador = strings.TrimSpace(u.Nombres_evaluador)
	u.Nombres_evaluador = strings.ToUpper(u.Nombres_evaluador)
	u.Nombres_evaluador = Utils.RemoveAccents(u.Nombres_evaluador)

	u.Apellidos_evaluador = strings.TrimSpace(u.Apellidos_evaluador)
	u.Apellidos_evaluador = strings.ToUpper(u.Apellidos_evaluador)
	u.Apellidos_evaluador = Utils.RemoveAccents(u.Apellidos_evaluador)

	u.Correo_electronico_evaluador = strings.TrimSpace(u.Correo_electronico_evaluador)
	u.Correo_electronico_evaluador = strings.ToUpper(u.Correo_electronico_evaluador)
	u.Correo_electronico_evaluador = Utils.RemoveAccents(u.Correo_electronico_evaluador)

	u.Telefono_fijo_evaluador = strings.TrimSpace(u.Telefono_fijo_evaluador)
	u.Telefono_fijo_evaluador = strings.ToUpper(u.Telefono_fijo_evaluador)
	u.Telefono_fijo_evaluador = Utils.RemoveAccents(u.Telefono_fijo_evaluador)

	u.Telefono_celular_evaluador = strings.TrimSpace(u.Telefono_celular_evaluador)
	u.Telefono_celular_evaluador = strings.ToUpper(u.Telefono_celular_evaluador)
	u.Telefono_celular_evaluador = Utils.RemoveAccents(u.Telefono_celular_evaluador)

	u.Recinto_evaluador = strings.TrimSpace(u.Recinto_evaluador)
	u.Recinto_evaluador = strings.ToUpper(u.Recinto_evaluador)
	u.Recinto_evaluador = Utils.RemoveAccents(u.Recinto_evaluador)

	u.Cargo_evaluador = strings.TrimSpace(u.Cargo_evaluador)
	u.Cargo_evaluador = strings.ToUpper(u.Cargo_evaluador)
	u.Cargo_evaluador = Utils.RemoveAccents(u.Cargo_evaluador)
}

func PutOneEvaluadorInput(u *Request.PutOneEvaluadorPayload) {
	u.Rut_evaluador = strings.TrimSpace(u.Rut_evaluador)
	u.Rut_evaluador = strings.ToUpper(u.Rut_evaluador)
	u.Rut_evaluador = Utils.RemoveAccents(u.Rut_evaluador)

	u.Nombres_evaluador = strings.TrimSpace(u.Nombres_evaluador)
	u.Nombres_evaluador = strings.ToUpper(u.Nombres_evaluador)
	u.Nombres_evaluador = Utils.RemoveAccents(u.Nombres_evaluador)

	u.Apellidos_evaluador = strings.TrimSpace(u.Apellidos_evaluador)
	u.Apellidos_evaluador = strings.ToUpper(u.Apellidos_evaluador)
	u.Apellidos_evaluador = Utils.RemoveAccents(u.Apellidos_evaluador)

	u.Correo_electronico_evaluador = strings.TrimSpace(u.Correo_electronico_evaluador)
	u.Correo_electronico_evaluador = strings.ToUpper(u.Correo_electronico_evaluador)
	u.Correo_electronico_evaluador = Utils.RemoveAccents(u.Correo_electronico_evaluador)

	u.Telefono_fijo_evaluador = strings.TrimSpace(u.Telefono_fijo_evaluador)
	u.Telefono_fijo_evaluador = strings.ToUpper(u.Telefono_fijo_evaluador)
	u.Telefono_fijo_evaluador = Utils.RemoveAccents(u.Telefono_fijo_evaluador)

	u.Telefono_celular_evaluador = strings.TrimSpace(u.Telefono_celular_evaluador)
	u.Telefono_celular_evaluador = strings.ToUpper(u.Telefono_celular_evaluador)
	u.Telefono_celular_evaluador = Utils.RemoveAccents(u.Telefono_celular_evaluador)

	u.Recinto_evaluador = strings.TrimSpace(u.Recinto_evaluador)
	u.Recinto_evaluador = strings.ToUpper(u.Recinto_evaluador)
	u.Recinto_evaluador = Utils.RemoveAccents(u.Recinto_evaluador)

	u.Cargo_evaluador = strings.TrimSpace(u.Cargo_evaluador)
	u.Cargo_evaluador = strings.ToUpper(u.Cargo_evaluador)
	u.Cargo_evaluador = Utils.RemoveAccents(u.Cargo_evaluador)
}

func DeleteEvaluadorInput(u *Request.DeleteEvaluadorPayload) {

}

func GetMyEvaluadorInput(u *Request.GetMyEvaluadorPayload) {

}

func PutMyEvaluadorInput(u *Request.PutMyEvaluadorPayload) {
	u.Rut_evaluador = strings.TrimSpace(u.Rut_evaluador)
	u.Rut_evaluador = strings.ToUpper(u.Rut_evaluador)
	u.Rut_evaluador = Utils.RemoveAccents(u.Rut_evaluador)

	u.Nombres_evaluador = strings.TrimSpace(u.Nombres_evaluador)
	u.Nombres_evaluador = strings.ToUpper(u.Nombres_evaluador)
	u.Nombres_evaluador = Utils.RemoveAccents(u.Nombres_evaluador)

	u.Apellidos_evaluador = strings.TrimSpace(u.Apellidos_evaluador)
	u.Apellidos_evaluador = strings.ToUpper(u.Apellidos_evaluador)
	u.Apellidos_evaluador = Utils.RemoveAccents(u.Apellidos_evaluador)

	u.Correo_electronico_evaluador = strings.TrimSpace(u.Correo_electronico_evaluador)
	u.Correo_electronico_evaluador = strings.ToUpper(u.Correo_electronico_evaluador)
	u.Correo_electronico_evaluador = Utils.RemoveAccents(u.Correo_electronico_evaluador)

	u.Telefono_fijo_evaluador = strings.TrimSpace(u.Telefono_fijo_evaluador)
	u.Telefono_fijo_evaluador = strings.ToUpper(u.Telefono_fijo_evaluador)
	u.Telefono_fijo_evaluador = Utils.RemoveAccents(u.Telefono_fijo_evaluador)

	u.Telefono_celular_evaluador = strings.TrimSpace(u.Telefono_celular_evaluador)
	u.Telefono_celular_evaluador = strings.ToUpper(u.Telefono_celular_evaluador)
	u.Telefono_celular_evaluador = Utils.RemoveAccents(u.Telefono_celular_evaluador)

	u.Recinto_evaluador = strings.TrimSpace(u.Recinto_evaluador)
	u.Recinto_evaluador = strings.ToUpper(u.Recinto_evaluador)
	u.Recinto_evaluador = Utils.RemoveAccents(u.Recinto_evaluador)

	u.Cargo_evaluador = strings.TrimSpace(u.Cargo_evaluador)
	u.Cargo_evaluador = strings.ToUpper(u.Cargo_evaluador)
	u.Cargo_evaluador = Utils.RemoveAccents(u.Cargo_evaluador)
}

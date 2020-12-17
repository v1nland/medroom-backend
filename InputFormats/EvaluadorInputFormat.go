package InputFormats

import (
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"
	"strings"
)

func GetEvaluadoresInput(u *RequestMessages.ListEvaluadoresPayload) {

}

func GetOneEvaluadorInput(u *RequestMessages.GetOneEvaluadorPayload) {

}

func AddNewEvaluadorInput(u *RequestMessages.AddNewEvaluadorPayload) {
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

func PutOneEvaluadorInput(u *RequestMessages.PutOneEvaluadorPayload) {
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

func DeleteEvaluadorInput(u *RequestMessages.DeleteEvaluadorPayload) {

}

func GetMyEvaluadorInput(u *RequestMessages.GetMyEvaluadorPayload) {

}

func PutMyEvaluadorInput(u *RequestMessages.PutMyEvaluadorPayload) {
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

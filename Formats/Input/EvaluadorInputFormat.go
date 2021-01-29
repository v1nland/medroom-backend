package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func ListEvaluadores(u *Request.ListEvaluadores) {
}

func GetOneEvaluador(u *Request.GetOneEvaluador) {
}

func AddNewEvaluador(u *Request.AddNewEvaluador) {
	if u.Rut_evaluador != nil {
		*u.Rut_evaluador = strings.TrimSpace(*u.Rut_evaluador)
		*u.Rut_evaluador = strings.ToUpper(*u.Rut_evaluador)
		*u.Rut_evaluador = Utils.RemoveAccents(*u.Rut_evaluador)
	}
	if u.Nombres_evaluador != nil {
		*u.Nombres_evaluador = strings.TrimSpace(*u.Nombres_evaluador)
		*u.Nombres_evaluador = strings.ToUpper(*u.Nombres_evaluador)
		*u.Nombres_evaluador = Utils.RemoveAccents(*u.Nombres_evaluador)
	}
	if u.Apellidos_evaluador != nil {
		*u.Apellidos_evaluador = strings.TrimSpace(*u.Apellidos_evaluador)
		*u.Apellidos_evaluador = strings.ToUpper(*u.Apellidos_evaluador)
		*u.Apellidos_evaluador = Utils.RemoveAccents(*u.Apellidos_evaluador)
	}
	if u.Correo_electronico_evaluador != nil {
		*u.Correo_electronico_evaluador = strings.TrimSpace(*u.Correo_electronico_evaluador)
		*u.Correo_electronico_evaluador = strings.ToUpper(*u.Correo_electronico_evaluador)
		*u.Correo_electronico_evaluador = Utils.RemoveAccents(*u.Correo_electronico_evaluador)
	}
	if u.Telefono_fijo_evaluador != nil {
		*u.Telefono_fijo_evaluador = strings.TrimSpace(*u.Telefono_fijo_evaluador)
		*u.Telefono_fijo_evaluador = strings.ToUpper(*u.Telefono_fijo_evaluador)
		*u.Telefono_fijo_evaluador = Utils.RemoveAccents(*u.Telefono_fijo_evaluador)
	}
	if u.Telefono_celular_evaluador != nil {
		*u.Telefono_celular_evaluador = strings.TrimSpace(*u.Telefono_celular_evaluador)
		*u.Telefono_celular_evaluador = strings.ToUpper(*u.Telefono_celular_evaluador)
		*u.Telefono_celular_evaluador = Utils.RemoveAccents(*u.Telefono_celular_evaluador)
	}
	if u.Recinto_evaluador != nil {
		*u.Recinto_evaluador = strings.TrimSpace(*u.Recinto_evaluador)
		*u.Recinto_evaluador = strings.ToUpper(*u.Recinto_evaluador)
		*u.Recinto_evaluador = Utils.RemoveAccents(*u.Recinto_evaluador)
	}
	if u.Cargo_evaluador != nil {
		*u.Cargo_evaluador = strings.TrimSpace(*u.Cargo_evaluador)
		*u.Cargo_evaluador = strings.ToUpper(*u.Cargo_evaluador)
		*u.Cargo_evaluador = Utils.RemoveAccents(*u.Cargo_evaluador)
	}
}

func PutOneEvaluador(u *Request.PutOneEvaluador) {
	if u.Rut_evaluador != nil {
		*u.Rut_evaluador = strings.TrimSpace(*u.Rut_evaluador)
		*u.Rut_evaluador = strings.ToUpper(*u.Rut_evaluador)
		*u.Rut_evaluador = Utils.RemoveAccents(*u.Rut_evaluador)
	}
	if u.Nombres_evaluador != nil {
		*u.Nombres_evaluador = strings.TrimSpace(*u.Nombres_evaluador)
		*u.Nombres_evaluador = strings.ToUpper(*u.Nombres_evaluador)
		*u.Nombres_evaluador = Utils.RemoveAccents(*u.Nombres_evaluador)
	}
	if u.Apellidos_evaluador != nil {
		*u.Apellidos_evaluador = strings.TrimSpace(*u.Apellidos_evaluador)
		*u.Apellidos_evaluador = strings.ToUpper(*u.Apellidos_evaluador)
		*u.Apellidos_evaluador = Utils.RemoveAccents(*u.Apellidos_evaluador)
	}
	if u.Correo_electronico_evaluador != nil {
		*u.Correo_electronico_evaluador = strings.TrimSpace(*u.Correo_electronico_evaluador)
		*u.Correo_electronico_evaluador = strings.ToUpper(*u.Correo_electronico_evaluador)
		*u.Correo_electronico_evaluador = Utils.RemoveAccents(*u.Correo_electronico_evaluador)
	}
	if u.Telefono_fijo_evaluador != nil {
		*u.Telefono_fijo_evaluador = strings.TrimSpace(*u.Telefono_fijo_evaluador)
		*u.Telefono_fijo_evaluador = strings.ToUpper(*u.Telefono_fijo_evaluador)
		*u.Telefono_fijo_evaluador = Utils.RemoveAccents(*u.Telefono_fijo_evaluador)
	}
	if u.Telefono_celular_evaluador != nil {
		*u.Telefono_celular_evaluador = strings.TrimSpace(*u.Telefono_celular_evaluador)
		*u.Telefono_celular_evaluador = strings.ToUpper(*u.Telefono_celular_evaluador)
		*u.Telefono_celular_evaluador = Utils.RemoveAccents(*u.Telefono_celular_evaluador)
	}
	if u.Recinto_evaluador != nil {
		*u.Recinto_evaluador = strings.TrimSpace(*u.Recinto_evaluador)
		*u.Recinto_evaluador = strings.ToUpper(*u.Recinto_evaluador)
		*u.Recinto_evaluador = Utils.RemoveAccents(*u.Recinto_evaluador)
	}
	if u.Cargo_evaluador != nil {
		*u.Cargo_evaluador = strings.TrimSpace(*u.Cargo_evaluador)
		*u.Cargo_evaluador = strings.ToUpper(*u.Cargo_evaluador)
		*u.Cargo_evaluador = Utils.RemoveAccents(*u.Cargo_evaluador)
	}
}

func DeleteEvaluador(u *Request.DeleteEvaluador) {
}

func GetMyEvaluador(u *Request.GetMyEvaluador) {
}

func PutMyEvaluador(u *Request.PutMyEvaluador) {
	if u.Telefono_fijo_evaluador != nil {
		*u.Telefono_fijo_evaluador = strings.TrimSpace(*u.Telefono_fijo_evaluador)
		*u.Telefono_fijo_evaluador = strings.ToUpper(*u.Telefono_fijo_evaluador)
		*u.Telefono_fijo_evaluador = Utils.RemoveAccents(*u.Telefono_fijo_evaluador)
	}
	if u.Telefono_celular_evaluador != nil {
		*u.Telefono_celular_evaluador = strings.TrimSpace(*u.Telefono_celular_evaluador)
		*u.Telefono_celular_evaluador = strings.ToUpper(*u.Telefono_celular_evaluador)
		*u.Telefono_celular_evaluador = Utils.RemoveAccents(*u.Telefono_celular_evaluador)
	}
	if u.Cargo_evaluador != nil {
		*u.Cargo_evaluador = strings.TrimSpace(*u.Cargo_evaluador)
		*u.Cargo_evaluador = strings.ToUpper(*u.Cargo_evaluador)
		*u.Cargo_evaluador = Utils.RemoveAccents(*u.Cargo_evaluador)
	}
}

func AddEvaluadorToGrupo(u *Request.AddEvaluadorToGrupo) {
}

package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func GetEstudiantesInput(u *Request.ListEstudiantesPayload) {

}

func GetOneEstudianteInput(u *Request.GetOneEstudiantePayload) {

}

func AddNewEstudianteInput(u *Request.AddNewEstudiantePayload) {
	u.Rut_estudiante = strings.TrimSpace(u.Rut_estudiante)
	u.Rut_estudiante = strings.ToUpper(u.Rut_estudiante)
	u.Rut_estudiante = Utils.RemoveAccents(u.Rut_estudiante)

	u.Nombres_estudiante = strings.TrimSpace(u.Nombres_estudiante)
	u.Nombres_estudiante = strings.ToUpper(u.Nombres_estudiante)
	u.Nombres_estudiante = Utils.RemoveAccents(u.Nombres_estudiante)

	u.Apellidos_estudiante = strings.TrimSpace(u.Apellidos_estudiante)
	u.Apellidos_estudiante = strings.ToUpper(u.Apellidos_estudiante)
	u.Apellidos_estudiante = Utils.RemoveAccents(u.Apellidos_estudiante)

	u.Correo_electronico_estudiante = strings.TrimSpace(u.Correo_electronico_estudiante)
	u.Correo_electronico_estudiante = strings.ToUpper(u.Correo_electronico_estudiante)
	u.Correo_electronico_estudiante = Utils.RemoveAccents(u.Correo_electronico_estudiante)

	u.Telefono_fijo_estudiante = strings.TrimSpace(u.Telefono_fijo_estudiante)
	u.Telefono_fijo_estudiante = strings.ToUpper(u.Telefono_fijo_estudiante)
	u.Telefono_fijo_estudiante = Utils.RemoveAccents(u.Telefono_fijo_estudiante)

	u.Telefono_celular_estudiante = strings.TrimSpace(u.Telefono_celular_estudiante)
	u.Telefono_celular_estudiante = strings.ToUpper(u.Telefono_celular_estudiante)
	u.Telefono_celular_estudiante = Utils.RemoveAccents(u.Telefono_celular_estudiante)
}

func PutOneEstudianteInput(u *Request.PutOneEstudiantePayload) {
	u.Rut_estudiante = strings.TrimSpace(u.Rut_estudiante)
	u.Rut_estudiante = strings.ToUpper(u.Rut_estudiante)
	u.Rut_estudiante = Utils.RemoveAccents(u.Rut_estudiante)

	u.Nombres_estudiante = strings.TrimSpace(u.Nombres_estudiante)
	u.Nombres_estudiante = strings.ToUpper(u.Nombres_estudiante)
	u.Nombres_estudiante = Utils.RemoveAccents(u.Nombres_estudiante)

	u.Apellidos_estudiante = strings.TrimSpace(u.Apellidos_estudiante)
	u.Apellidos_estudiante = strings.ToUpper(u.Apellidos_estudiante)
	u.Apellidos_estudiante = Utils.RemoveAccents(u.Apellidos_estudiante)

	u.Correo_electronico_estudiante = strings.TrimSpace(u.Correo_electronico_estudiante)
	u.Correo_electronico_estudiante = strings.ToUpper(u.Correo_electronico_estudiante)
	u.Correo_electronico_estudiante = Utils.RemoveAccents(u.Correo_electronico_estudiante)

	u.Telefono_fijo_estudiante = strings.TrimSpace(u.Telefono_fijo_estudiante)
	u.Telefono_fijo_estudiante = strings.ToUpper(u.Telefono_fijo_estudiante)
	u.Telefono_fijo_estudiante = Utils.RemoveAccents(u.Telefono_fijo_estudiante)

	u.Telefono_celular_estudiante = strings.TrimSpace(u.Telefono_celular_estudiante)
	u.Telefono_celular_estudiante = strings.ToUpper(u.Telefono_celular_estudiante)
	u.Telefono_celular_estudiante = Utils.RemoveAccents(u.Telefono_celular_estudiante)
}

func DeleteEstudianteInput(u *Request.DeleteEstudiantePayload) {

}

func GetMyEstudianteInput(u *Request.GetMyEstudiantePayload) {

}

func PutMyEstudianteInput(u *Request.PutMyEstudiantePayload) {
	u.Rut_estudiante = strings.TrimSpace(u.Rut_estudiante)
	u.Rut_estudiante = strings.ToUpper(u.Rut_estudiante)
	u.Rut_estudiante = Utils.RemoveAccents(u.Rut_estudiante)

	u.Nombres_estudiante = strings.TrimSpace(u.Nombres_estudiante)
	u.Nombres_estudiante = strings.ToUpper(u.Nombres_estudiante)
	u.Nombres_estudiante = Utils.RemoveAccents(u.Nombres_estudiante)

	u.Apellidos_estudiante = strings.TrimSpace(u.Apellidos_estudiante)
	u.Apellidos_estudiante = strings.ToUpper(u.Apellidos_estudiante)
	u.Apellidos_estudiante = Utils.RemoveAccents(u.Apellidos_estudiante)

	u.Correo_electronico_estudiante = strings.TrimSpace(u.Correo_electronico_estudiante)
	u.Correo_electronico_estudiante = strings.ToUpper(u.Correo_electronico_estudiante)
	u.Correo_electronico_estudiante = Utils.RemoveAccents(u.Correo_electronico_estudiante)

	u.Telefono_fijo_estudiante = strings.TrimSpace(u.Telefono_fijo_estudiante)
	u.Telefono_fijo_estudiante = strings.ToUpper(u.Telefono_fijo_estudiante)
	u.Telefono_fijo_estudiante = Utils.RemoveAccents(u.Telefono_fijo_estudiante)

	u.Telefono_celular_estudiante = strings.TrimSpace(u.Telefono_celular_estudiante)
	u.Telefono_celular_estudiante = strings.ToUpper(u.Telefono_celular_estudiante)
	u.Telefono_celular_estudiante = Utils.RemoveAccents(u.Telefono_celular_estudiante)
}

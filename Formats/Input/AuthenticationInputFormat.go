package Input

import (
	"medroom-backend/Messages/Request"
	"strings"
)

func FormatLoginEstudianteMessage(msg *Request.LoginEstudiantePayload) {
	// format trim
	msg.Hash_contrasena_estudiante = strings.TrimSpace(msg.Hash_contrasena_estudiante)
	msg.Correo_electronico_estudiante = strings.ToUpper(msg.Correo_electronico_estudiante)
}

func FormatLoginEvaluadorMessage(msg *Request.LoginEvaluadorPayload) {
	// format trim
	msg.Hash_contrasena_evaluador = strings.TrimSpace(msg.Hash_contrasena_evaluador)
	msg.Correo_electronico_evaluador = strings.ToUpper(msg.Correo_electronico_evaluador)
}

func FormatLoginAdministradorAcademicoMessage(msg *Request.LoginAdministradorAcademicoPayload) {
	// format trim
	msg.Hash_contrasena_administrador_academico = strings.TrimSpace(msg.Hash_contrasena_administrador_academico)
	msg.Correo_electronico_administrador_academico = strings.ToUpper(msg.Correo_electronico_administrador_academico)
}

func FormatLoginAdministradorTiMessage(msg *Request.LoginAdministradorTiPayload) {
	// format trim
	msg.Hash_contrasena_administrador_ti = strings.TrimSpace(msg.Hash_contrasena_administrador_ti)
	msg.Correo_electronico_administrador_ti = strings.ToUpper(msg.Correo_electronico_administrador_ti)
}

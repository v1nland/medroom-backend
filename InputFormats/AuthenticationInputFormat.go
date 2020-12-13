package InputFormats

import (
	"medroom-backend/RequestMessages"
	"strings"
)

func FormatLoginEstudianteMessage(msg *RequestMessages.LoginEstudiantePayload) {
	// format trim
	msg.Hash_contrasena_estudiante = strings.TrimSpace(msg.Hash_contrasena_estudiante)

	// format uppercase
	msg.Hash_contrasena_estudiante = strings.ToUpper(msg.Hash_contrasena_estudiante)
}

func FormatLoginEvaluadorMessage(msg *RequestMessages.LoginEvaluadorPayload) {
	// format trim
	msg.Hash_contrasena_evaluador = strings.TrimSpace(msg.Hash_contrasena_evaluador)

	// format uppercase
	msg.Hash_contrasena_evaluador = strings.ToUpper(msg.Hash_contrasena_evaluador)
}

func FormatLoginAdministradorAcademicoMessage(msg *RequestMessages.LoginAdministradorAcademicoPayload) {
	// format trim
	msg.Hash_contrasena_administrador_academico = strings.TrimSpace(msg.Hash_contrasena_administrador_academico)

	// format uppercase
	msg.Hash_contrasena_administrador_academico = strings.ToUpper(msg.Hash_contrasena_administrador_academico)
}

func FormatLoginAdministradorTiMessage(msg *RequestMessages.LoginAdministradorTiPayload) {
	// format trim
	msg.Hash_contrasena_administrador_ti = strings.TrimSpace(msg.Hash_contrasena_administrador_ti)

	// format uppercase
	msg.Hash_contrasena_administrador_ti = strings.ToUpper(msg.Hash_contrasena_administrador_ti)
}

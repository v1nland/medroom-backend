package f_input

import (
	"medroom-backend/app/messages/Request"
	"strings"
)

func LoginEstudiante(msg *Request.LoginEstudiante) {
	msg.Hash_contrasena_estudiante = strings.TrimSpace(msg.Hash_contrasena_estudiante)
	msg.Correo_electronico_estudiante = strings.ToUpper(msg.Correo_electronico_estudiante)
}

func LoginEvaluador(msg *Request.LoginEvaluador) {
	msg.Hash_contrasena_evaluador = strings.TrimSpace(msg.Hash_contrasena_evaluador)
	msg.Correo_electronico_evaluador = strings.ToUpper(msg.Correo_electronico_evaluador)
}

func LoginAdministradorAcademico(msg *Request.LoginAdministradorAcademico) {
	msg.Hash_contrasena_administrador_academico = strings.TrimSpace(msg.Hash_contrasena_administrador_academico)
	msg.Correo_electronico_administrador_academico = strings.ToUpper(msg.Correo_electronico_administrador_academico)
}

func LoginAdministradorTi(msg *Request.LoginAdministradorTi) {
	msg.Hash_contrasena_administrador_ti = strings.TrimSpace(msg.Hash_contrasena_administrador_ti)
	msg.Correo_electronico_administrador_ti = strings.ToUpper(msg.Correo_electronico_administrador_ti)
}

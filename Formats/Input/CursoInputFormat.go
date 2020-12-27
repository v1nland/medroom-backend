package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func ListCursosInput(u *Request.ListCursosPayload) {
}

func GetOneCursoInput(u *Request.GetOneCursoPayload) {
}

func GetOneCursoEstudianteInput(u *Request.GetOneCursoEstudiantePayload) {
}

func GetCursosEstudianteInput(u *Request.GetCursosEstudiantePayload) {
}

func GetOneCursoEvaluadorInput(u *Request.GetOneCursoEvaluadorPayload) {
}

func GetCursosEvaluadorInput(u *Request.GetCursosEvaluadorPayload) {
}

func AddNewCursoInput(u *Request.AddNewCursoPayload) {
	if u.Nombre_curso != nil {
		*u.Nombre_curso = strings.TrimSpace(*u.Nombre_curso)
		*u.Nombre_curso = strings.ToUpper(*u.Nombre_curso)
		*u.Nombre_curso = Utils.RemoveAccents(*u.Nombre_curso)
	}
	if u.Sigla_curso != nil {
		*u.Sigla_curso = strings.TrimSpace(*u.Sigla_curso)
		*u.Sigla_curso = strings.ToUpper(*u.Sigla_curso)
		*u.Sigla_curso = Utils.RemoveAccents(*u.Sigla_curso)
	}
}

func PutOneCursoInput(u *Request.PutOneCursoPayload) {
	if u.Nombre_curso != nil {
		*u.Nombre_curso = strings.TrimSpace(*u.Nombre_curso)
		*u.Nombre_curso = strings.ToUpper(*u.Nombre_curso)
		*u.Nombre_curso = Utils.RemoveAccents(*u.Nombre_curso)
	}
	if u.Sigla_curso != nil {
		*u.Sigla_curso = strings.TrimSpace(*u.Sigla_curso)
		*u.Sigla_curso = strings.ToUpper(*u.Sigla_curso)
		*u.Sigla_curso = Utils.RemoveAccents(*u.Sigla_curso)
	}
}

func DeleteCursoInput(u *Request.DeleteCursoPayload) {

}

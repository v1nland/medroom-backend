package InputFormats

import (
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"
	"strings"
)

func GetCursosInput(u *RequestMessages.ListCursosPayload) {

}

func GetOneCursoInput(u *RequestMessages.GetOneCursoPayload) {

}

func GetCursoEstudianteInput(u *RequestMessages.GetCursoEstudiantePayload) {

}

func GetCursoEvaluadorInput(u *RequestMessages.GetCursoEvaluadorPayload) {

}

func AddNewCursoInput(u *RequestMessages.AddNewCursoPayload) {
	u.Nombre_curso = strings.TrimSpace(u.Nombre_curso)
	u.Nombre_curso = strings.ToUpper(u.Nombre_curso)
	u.Nombre_curso = Utils.RemoveAccents(u.Nombre_curso)

	u.Sigla_curso = strings.TrimSpace(u.Sigla_curso)
	u.Sigla_curso = strings.ToUpper(u.Sigla_curso)
	u.Sigla_curso = Utils.RemoveAccents(u.Sigla_curso)
}

func PutOneCursoInput(u *RequestMessages.PutOneCursoPayload) {
	u.Nombre_curso = strings.TrimSpace(u.Nombre_curso)
	u.Nombre_curso = strings.ToUpper(u.Nombre_curso)
	u.Nombre_curso = Utils.RemoveAccents(u.Nombre_curso)

	u.Sigla_curso = strings.TrimSpace(u.Sigla_curso)
	u.Sigla_curso = strings.ToUpper(u.Sigla_curso)
	u.Sigla_curso = Utils.RemoveAccents(u.Sigla_curso)
}

func DeleteCursoInput(u *RequestMessages.DeleteCursoPayload) {

}

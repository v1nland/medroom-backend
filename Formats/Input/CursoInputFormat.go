package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func GetCursosInput(u *Request.ListCursosPayload) {

}

func GetOneCursoInput(u *Request.GetOneCursoPayload) {

}

func GetCursoEstudianteInput(u *Request.GetCursoEstudiantePayload) {

}

func GetCursoEvaluadorInput(u *Request.GetCursoEvaluadorPayload) {

}

func AddNewCursoInput(u *Request.AddNewCursoPayload) {
	u.Nombre_curso = strings.TrimSpace(u.Nombre_curso)
	u.Nombre_curso = strings.ToUpper(u.Nombre_curso)
	u.Nombre_curso = Utils.RemoveAccents(u.Nombre_curso)

	u.Sigla_curso = strings.TrimSpace(u.Sigla_curso)
	u.Sigla_curso = strings.ToUpper(u.Sigla_curso)
	u.Sigla_curso = Utils.RemoveAccents(u.Sigla_curso)
}

func PutOneCursoInput(u *Request.PutOneCursoPayload) {
	u.Nombre_curso = strings.TrimSpace(u.Nombre_curso)
	u.Nombre_curso = strings.ToUpper(u.Nombre_curso)
	u.Nombre_curso = Utils.RemoveAccents(u.Nombre_curso)

	u.Sigla_curso = strings.TrimSpace(u.Sigla_curso)
	u.Sigla_curso = strings.ToUpper(u.Sigla_curso)
	u.Sigla_curso = Utils.RemoveAccents(u.Sigla_curso)
}

func DeleteCursoInput(u *Request.DeleteCursoPayload) {

}

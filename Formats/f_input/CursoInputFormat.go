package f_input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func ListCursos(u *Request.ListCursos) {
}

func GetOneCurso(u *Request.GetOneCurso) {
}

func GetOneCursoEstudiante(u *Request.GetOneCursoEstudiante) {
}

func GetCursosEstudiante(u *Request.GetCursosEstudiante) {
}

func GetOneCursoEvaluador(u *Request.GetOneCursoEvaluador) {
}

func GetCursosEvaluador(u *Request.GetCursosEvaluador) {
}

func AddNewCurso(u *Request.AddNewCurso) {
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

func PutOneCurso(u *Request.PutOneCurso) {
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

func DeleteCurso(u *Request.DeleteCurso) {

}

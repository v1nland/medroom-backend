package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func ListEvaluacionesGrupoEstudiante(u *Request.ListEvaluacionesGrupoEstudiante) {
}

func ListEvaluacionesGrupoEvaluador(u *Request.ListEvaluacionesGrupoEvaluador) {
}

func ListEvaluaciones(u *Request.ListEvaluaciones) {
}

func GetEvaluacionesEstudiante(u *Request.ListEvaluacionesEstudiante) {
}

func GetOneEvaluacion(u *Request.GetOneEvaluacion) {
}

func AddNewEvaluacion(u *Request.AddNewEvaluacion) {
	if u.Nombre_evaluacion != nil {
		*u.Nombre_evaluacion = strings.TrimSpace(*u.Nombre_evaluacion)
		*u.Nombre_evaluacion = strings.ToUpper(*u.Nombre_evaluacion)
		*u.Nombre_evaluacion = Utils.RemoveAccents(*u.Nombre_evaluacion)
	}
}
func PutOneEvaluacion(u *Request.PutOneEvaluacion) {
	if u.Nombre_evaluacion != nil {
		*u.Nombre_evaluacion = strings.TrimSpace(*u.Nombre_evaluacion)
		*u.Nombre_evaluacion = strings.ToUpper(*u.Nombre_evaluacion)
		*u.Nombre_evaluacion = Utils.RemoveAccents(*u.Nombre_evaluacion)
	}
}

func DeleteEvaluacion(u *Request.DeleteEvaluacion) {
}

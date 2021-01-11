package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func ListEvaluacionesGrupoEstudianteInput(u *Request.ListEvaluacionesGrupoEstudiantePayload) {
}
func ListEvaluacionesGrupoEvaluadorInput(u *Request.ListEvaluacionesGrupoEvaluadorPayload) {
}
func ListEvaluacionesInput(u *Request.ListEvaluacionesPayload) {
}
func GetEvaluacionesEstudianteInput(u *Request.ListEvaluacionesEstudiantePayload) {
}
func GetOneEvaluacionInput(u *Request.GetOneEvaluacionPayload) {
}

func AddNewEvaluacionInput(u *Request.AddNewEvaluacionPayload) {
	if u.Nombre_evaluacion != nil {
		*u.Nombre_evaluacion = strings.TrimSpace(*u.Nombre_evaluacion)
		*u.Nombre_evaluacion = strings.ToUpper(*u.Nombre_evaluacion)
		*u.Nombre_evaluacion = Utils.RemoveAccents(*u.Nombre_evaluacion)
	}
}
func PutOneEvaluacionInput(u *Request.PutOneEvaluacionPayload) {
	if u.Nombre_evaluacion != nil {
		*u.Nombre_evaluacion = strings.TrimSpace(*u.Nombre_evaluacion)
		*u.Nombre_evaluacion = strings.ToUpper(*u.Nombre_evaluacion)
		*u.Nombre_evaluacion = Utils.RemoveAccents(*u.Nombre_evaluacion)
	}
}

func DeleteEvaluacionInput(u *Request.DeleteEvaluacionPayload) {
}

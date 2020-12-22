package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func ListPuntajesInput(u *Request.ListPuntajesPayload) {

}

func GetOnePuntajeInput(u *Request.GetOnePuntajePayload) {

}

func AddNewPuntajeInput(u *Request.AddNewPuntajePayload) {
	u.Nombre_competencia_puntaje = strings.TrimSpace(u.Nombre_competencia_puntaje)
	u.Nombre_competencia_puntaje = strings.ToUpper(u.Nombre_competencia_puntaje)
	u.Nombre_competencia_puntaje = Utils.RemoveAccents(u.Nombre_competencia_puntaje)

	u.Codigo_competencia_puntaje = strings.TrimSpace(u.Codigo_competencia_puntaje)
	u.Codigo_competencia_puntaje = strings.ToUpper(u.Codigo_competencia_puntaje)
	u.Codigo_competencia_puntaje = Utils.RemoveAccents(u.Codigo_competencia_puntaje)

	u.Feedback_puntaje = strings.TrimSpace(u.Feedback_puntaje)
	u.Feedback_puntaje = strings.ToUpper(u.Feedback_puntaje)
	u.Feedback_puntaje = Utils.RemoveAccents(u.Feedback_puntaje)
}

func PutOnePuntajeInput(u *Request.PutOnePuntajePayload) {
	u.Nombre_competencia_puntaje = strings.TrimSpace(u.Nombre_competencia_puntaje)
	u.Nombre_competencia_puntaje = strings.ToUpper(u.Nombre_competencia_puntaje)
	u.Nombre_competencia_puntaje = Utils.RemoveAccents(u.Nombre_competencia_puntaje)

	u.Codigo_competencia_puntaje = strings.TrimSpace(u.Codigo_competencia_puntaje)
	u.Codigo_competencia_puntaje = strings.ToUpper(u.Codigo_competencia_puntaje)
	u.Codigo_competencia_puntaje = Utils.RemoveAccents(u.Codigo_competencia_puntaje)

	u.Feedback_puntaje = strings.TrimSpace(u.Feedback_puntaje)
	u.Feedback_puntaje = strings.ToUpper(u.Feedback_puntaje)
	u.Feedback_puntaje = Utils.RemoveAccents(u.Feedback_puntaje)
}

func DeletePuntajeInput(u *Request.DeletePuntajePayload) {

}

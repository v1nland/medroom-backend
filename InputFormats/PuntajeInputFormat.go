package InputFormats

import (
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"
	"strings"
)

func GetPuntajesInput(u *RequestMessages.ListPuntajesPayload) {

}

func GetOnePuntajeInput(u *RequestMessages.GetOnePuntajePayload) {

}

func AddNewPuntajeInput(u *RequestMessages.AddNewPuntajePayload) {
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

func PutOnePuntajeInput(u *RequestMessages.PutOnePuntajePayload) {
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

func DeletePuntajeInput(u *RequestMessages.DeletePuntajePayload) {

}

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
	u.Feedback_puntaje = strings.TrimSpace(u.Feedback_puntaje)
	u.Feedback_puntaje = strings.ToUpper(u.Feedback_puntaje)
	u.Feedback_puntaje = Utils.RemoveAccents(u.Feedback_puntaje)
}

func PutOnePuntajeInput(u *RequestMessages.PutOnePuntajePayload) {
	u.Feedback_puntaje = strings.TrimSpace(u.Feedback_puntaje)
	u.Feedback_puntaje = strings.ToUpper(u.Feedback_puntaje)
	u.Feedback_puntaje = Utils.RemoveAccents(u.Feedback_puntaje)
}

func DeletePuntajeInput(u *RequestMessages.DeletePuntajePayload) {

}

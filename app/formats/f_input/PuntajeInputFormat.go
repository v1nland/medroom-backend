package f_input

import (
	"medroom-backend/app/messages/Request"
	"medroom-backend/app/utils"
	"strings"
)

func ListPuntajes(u *Request.ListPuntajes) {
}

func GetOnePuntaje(u *Request.GetOnePuntaje) {
}

func AddNewPuntaje(u *Request.AddNewPuntaje) {
	u.Nombre_competencia_puntaje = strings.TrimSpace(u.Nombre_competencia_puntaje)
	u.Nombre_competencia_puntaje = strings.ToUpper(u.Nombre_competencia_puntaje)
	u.Nombre_competencia_puntaje = utils.RemoveAccents(u.Nombre_competencia_puntaje)

	u.Codigo_competencia_puntaje = strings.TrimSpace(u.Codigo_competencia_puntaje)
	u.Codigo_competencia_puntaje = strings.ToUpper(u.Codigo_competencia_puntaje)
	u.Codigo_competencia_puntaje = utils.RemoveAccents(u.Codigo_competencia_puntaje)

	u.Feedback_puntaje = strings.TrimSpace(u.Feedback_puntaje)
	u.Feedback_puntaje = strings.ToUpper(u.Feedback_puntaje)
	u.Feedback_puntaje = utils.RemoveAccents(u.Feedback_puntaje)
}

func PutOnePuntaje(u *Request.PutOnePuntaje) {
	u.Nombre_competencia_puntaje = strings.TrimSpace(u.Nombre_competencia_puntaje)
	u.Nombre_competencia_puntaje = strings.ToUpper(u.Nombre_competencia_puntaje)
	u.Nombre_competencia_puntaje = utils.RemoveAccents(u.Nombre_competencia_puntaje)

	u.Codigo_competencia_puntaje = strings.TrimSpace(u.Codigo_competencia_puntaje)
	u.Codigo_competencia_puntaje = strings.ToUpper(u.Codigo_competencia_puntaje)
	u.Codigo_competencia_puntaje = utils.RemoveAccents(u.Codigo_competencia_puntaje)

	u.Feedback_puntaje = strings.TrimSpace(u.Feedback_puntaje)
	u.Feedback_puntaje = strings.ToUpper(u.Feedback_puntaje)
	u.Feedback_puntaje = utils.RemoveAccents(u.Feedback_puntaje)
}

func DeletePuntaje(u *Request.DeletePuntaje) {
}

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
	u.Nivel_logro_puntaje = strings.TrimSpace(u.Nivel_logro_puntaje)
	u.Nivel_logro_puntaje = strings.ToUpper(u.Nivel_logro_puntaje)
	u.Nivel_logro_puntaje = Utils.RemoveAccents(u.Nivel_logro_puntaje)
}

func PutOnePuntajeInput(u *RequestMessages.PutOnePuntajePayload) {
	u.Nivel_logro_puntaje = strings.TrimSpace(u.Nivel_logro_puntaje)
	u.Nivel_logro_puntaje = strings.ToUpper(u.Nivel_logro_puntaje)
	u.Nivel_logro_puntaje = Utils.RemoveAccents(u.Nivel_logro_puntaje)
}

func DeletePuntajeInput(u *RequestMessages.DeletePuntajePayload) {

}

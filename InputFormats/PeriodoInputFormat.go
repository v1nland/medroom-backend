package InputFormats

import (
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"
	"strings"
)

func GetPeriodosInput(u *RequestMessages.ListPeriodosPayload) {

}

func GetOnePeriodoInput(u *RequestMessages.GetOnePeriodoPayload) {

}

func AddNewPeriodoInput(u *RequestMessages.AddNewPeriodoPayload) {
	u.Nombre_periodo = strings.TrimSpace(u.Nombre_periodo)
	u.Nombre_periodo = strings.ToUpper(u.Nombre_periodo)
	u.Nombre_periodo = Utils.RemoveAccents(u.Nombre_periodo)
}

func PutOnePeriodoInput(u *RequestMessages.PutOnePeriodoPayload) {
	u.Nombre_periodo = strings.TrimSpace(u.Nombre_periodo)
	u.Nombre_periodo = strings.ToUpper(u.Nombre_periodo)
	u.Nombre_periodo = Utils.RemoveAccents(u.Nombre_periodo)
}

func DeletePeriodoInput(u *RequestMessages.DeletePeriodoPayload) {

}

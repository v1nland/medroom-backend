package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func GetPeriodosInput(u *Request.ListPeriodosPayload) {

}

func GetOnePeriodoInput(u *Request.GetOnePeriodoPayload) {

}

func AddNewPeriodoInput(u *Request.AddNewPeriodoPayload) {
	u.Nombre_periodo = strings.TrimSpace(u.Nombre_periodo)
	u.Nombre_periodo = strings.ToUpper(u.Nombre_periodo)
	u.Nombre_periodo = Utils.RemoveAccents(u.Nombre_periodo)
}

func PutOnePeriodoInput(u *Request.PutOnePeriodoPayload) {
	u.Nombre_periodo = strings.TrimSpace(u.Nombre_periodo)
	u.Nombre_periodo = strings.ToUpper(u.Nombre_periodo)
	u.Nombre_periodo = Utils.RemoveAccents(u.Nombre_periodo)
}

func DeletePeriodoInput(u *Request.DeletePeriodoPayload) {

}

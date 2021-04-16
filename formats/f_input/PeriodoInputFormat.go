package f_input

import (
	"medroom-backend/messages/Request"
	"medroom-backend/utils"
	"strings"
)

func ListPeriodos(u *Request.ListPeriodos) {
}

func GetOnePeriodo(u *Request.GetOnePeriodo) {
}

func AddNewPeriodo(u *Request.AddNewPeriodo) {
	if u.Nombre_periodo != nil {
		*u.Nombre_periodo = strings.TrimSpace(*u.Nombre_periodo)
		*u.Nombre_periodo = strings.ToUpper(*u.Nombre_periodo)
		*u.Nombre_periodo = utils.RemoveAccents(*u.Nombre_periodo)
	}
}

func PutOnePeriodo(u *Request.PutOnePeriodo) {
	if u.Nombre_periodo != nil {
		*u.Nombre_periodo = strings.TrimSpace(*u.Nombre_periodo)
		*u.Nombre_periodo = strings.ToUpper(*u.Nombre_periodo)
		*u.Nombre_periodo = utils.RemoveAccents(*u.Nombre_periodo)
	}
}

func DeletePeriodo(u *Request.DeletePeriodo) {
}

package InputFormats

import (
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"
	"strings"
)

func GetRolsInput(u *RequestMessages.ListRolsPayload) {

}

func GetOneRolInput(u *RequestMessages.GetOneRolPayload) {

}

func AddNewRolInput(u *RequestMessages.AddNewRolPayload) {
	u.Nombre_rol = strings.TrimSpace(u.Nombre_rol)
	u.Nombre_rol = strings.ToUpper(u.Nombre_rol)
	u.Nombre_rol = Utils.RemoveAccents(u.Nombre_rol)
}

func PutOneRolInput(u *RequestMessages.PutOneRolPayload) {
	u.Nombre_rol = strings.TrimSpace(u.Nombre_rol)
	u.Nombre_rol = strings.ToUpper(u.Nombre_rol)
	u.Nombre_rol = Utils.RemoveAccents(u.Nombre_rol)
}

func DeleteRolInput(u *RequestMessages.DeleteRolPayload) {

}

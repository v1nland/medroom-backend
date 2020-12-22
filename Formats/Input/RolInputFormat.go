package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func GetRolesInput(u *Request.ListRolesPayload) {

}

func GetOneRolInput(u *Request.GetOneRolPayload) {

}

func AddNewRolInput(u *Request.AddNewRolPayload) {
	u.Nombre_rol = strings.TrimSpace(u.Nombre_rol)
	u.Nombre_rol = strings.ToUpper(u.Nombre_rol)
	u.Nombre_rol = Utils.RemoveAccents(u.Nombre_rol)
}

func PutOneRolInput(u *Request.PutOneRolPayload) {
	u.Nombre_rol = strings.TrimSpace(u.Nombre_rol)
	u.Nombre_rol = strings.ToUpper(u.Nombre_rol)
	u.Nombre_rol = Utils.RemoveAccents(u.Nombre_rol)
}

func DeleteRolInput(u *Request.DeleteRolPayload) {

}

package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func ListRoles(u *Request.ListRoles) {
}

func GetOneRol(u *Request.GetOneRol) {
}

func AddNewRol(u *Request.AddNewRol) {
	if u.Nombre_rol != nil {
		*u.Nombre_rol = strings.TrimSpace(*u.Nombre_rol)
		*u.Nombre_rol = strings.ToUpper(*u.Nombre_rol)
		*u.Nombre_rol = Utils.RemoveAccents(*u.Nombre_rol)
	}
}

func PutOneRol(u *Request.PutOneRol) {
	if u.Nombre_rol != nil {
		*u.Nombre_rol = strings.TrimSpace(*u.Nombre_rol)
		*u.Nombre_rol = strings.ToUpper(*u.Nombre_rol)
		*u.Nombre_rol = Utils.RemoveAccents(*u.Nombre_rol)
	}
}

func DeleteRol(u *Request.DeleteRol) {
}

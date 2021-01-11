package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func ListGruposInput(u *Request.ListGruposPayload) {
}

func GetOneGrupoInput(u *Request.GetOneGrupoPayload) {
}

func GetGrupoEstudianteInput(u *Request.GetGrupoEstudiantePayload) {
}

func GetGrupoEvaluadorInput(u *Request.GetGrupoEvaluadorPayload) {
}

func AddNewGrupoInput(u *Request.AddNewGrupoPayload) {
	if u.Nombre_grupo != nil {
		*u.Nombre_grupo = strings.TrimSpace(*u.Nombre_grupo)
		*u.Nombre_grupo = strings.ToUpper(*u.Nombre_grupo)
		*u.Nombre_grupo = Utils.RemoveAccents(*u.Nombre_grupo)
	}
	if u.Sigla_grupo != nil {
		*u.Sigla_grupo = strings.TrimSpace(*u.Sigla_grupo)
		*u.Sigla_grupo = strings.ToUpper(*u.Sigla_grupo)
		*u.Sigla_grupo = Utils.RemoveAccents(*u.Sigla_grupo)
	}
}

func PutOneGrupoInput(u *Request.PutOneGrupoPayload) {
	if u.Nombre_grupo != nil {
		*u.Nombre_grupo = strings.TrimSpace(*u.Nombre_grupo)
		*u.Nombre_grupo = strings.ToUpper(*u.Nombre_grupo)
		*u.Nombre_grupo = Utils.RemoveAccents(*u.Nombre_grupo)
	}
	if u.Sigla_grupo != nil {
		*u.Sigla_grupo = strings.TrimSpace(*u.Sigla_grupo)
		*u.Sigla_grupo = strings.ToUpper(*u.Sigla_grupo)
		*u.Sigla_grupo = Utils.RemoveAccents(*u.Sigla_grupo)
	}
}

func DeleteGrupoInput(u *Request.DeleteGrupoPayload) {

}

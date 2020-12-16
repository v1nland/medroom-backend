package InputFormats

import (
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"
	"strings"
)

func GetGruposInput(u *RequestMessages.ListGruposPayload) {

}

func GetOneGrupoInput(u *RequestMessages.GetOneGrupoPayload) {

}

func GetGrupoEstudianteInput(u *RequestMessages.GetGrupoEstudiantePayload) {

}

func AddNewGrupoInput(u *RequestMessages.AddNewGrupoPayload) {
	u.Nombre_grupo = strings.TrimSpace(u.Nombre_grupo)
	u.Nombre_grupo = strings.ToUpper(u.Nombre_grupo)
	u.Nombre_grupo = Utils.RemoveAccents(u.Nombre_grupo)

	u.Sigla_grupo = strings.TrimSpace(u.Sigla_grupo)
	u.Sigla_grupo = strings.ToUpper(u.Sigla_grupo)
	u.Sigla_grupo = Utils.RemoveAccents(u.Sigla_grupo)
}

func PutOneGrupoInput(u *RequestMessages.PutOneGrupoPayload) {
	u.Nombre_grupo = strings.TrimSpace(u.Nombre_grupo)
	u.Nombre_grupo = strings.ToUpper(u.Nombre_grupo)
	u.Nombre_grupo = Utils.RemoveAccents(u.Nombre_grupo)

	u.Sigla_grupo = strings.TrimSpace(u.Sigla_grupo)
	u.Sigla_grupo = strings.ToUpper(u.Sigla_grupo)
	u.Sigla_grupo = Utils.RemoveAccents(u.Sigla_grupo)
}

func DeleteGrupoInput(u *RequestMessages.DeleteGrupoPayload) {

}

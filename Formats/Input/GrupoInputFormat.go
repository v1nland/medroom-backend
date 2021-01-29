package Input

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"strings"
)

func ListGrupos(u *Request.ListGrupos) {
}

func GetOneGrupo(u *Request.GetOneGrupo) {
}

func GetGrupoEstudiante(u *Request.GetGrupoEstudiante) {
}

func GetGrupoEvaluador(u *Request.GetGrupoEvaluador) {
}

func AddNewGrupo(u *Request.AddNewGrupo) {
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

func PutOneGrupo(u *Request.PutOneGrupo) {
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

func DeleteGrupo(u *Request.DeleteGrupo) {

}

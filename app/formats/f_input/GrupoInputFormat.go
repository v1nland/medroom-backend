package f_input

import (
	"medroom-backend/app/messages/Request"
	"medroom-backend/app/utils"
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
		*u.Nombre_grupo = utils.RemoveAccents(*u.Nombre_grupo)
	}
	if u.Sigla_grupo != nil {
		*u.Sigla_grupo = strings.TrimSpace(*u.Sigla_grupo)
		*u.Sigla_grupo = strings.ToUpper(*u.Sigla_grupo)
		*u.Sigla_grupo = utils.RemoveAccents(*u.Sigla_grupo)
	}
}

func PutOneGrupo(u *Request.PutOneGrupo) {
	if u.Nombre_grupo != nil {
		*u.Nombre_grupo = strings.TrimSpace(*u.Nombre_grupo)
		*u.Nombre_grupo = strings.ToUpper(*u.Nombre_grupo)
		*u.Nombre_grupo = utils.RemoveAccents(*u.Nombre_grupo)
	}
	if u.Sigla_grupo != nil {
		*u.Sigla_grupo = strings.TrimSpace(*u.Sigla_grupo)
		*u.Sigla_grupo = strings.ToUpper(*u.Sigla_grupo)
		*u.Sigla_grupo = utils.RemoveAccents(*u.Sigla_grupo)
	}
}

func DeleteGrupo(u *Request.DeleteGrupo) {

}

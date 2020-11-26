package InputFormats

import (
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"
	"strings"
)

func GetCompetenciasInput(u *RequestMessages.ListCompetenciasPayload) {

}

func GetOneCompetenciaInput(u *RequestMessages.GetOneCompetenciaPayload) {

}

func AddNewCompetenciaInput(u *RequestMessages.AddNewCompetenciaPayload) {
	u.Nombre_competencia = strings.TrimSpace(u.Nombre_competencia)
	u.Nombre_competencia = strings.ToUpper(u.Nombre_competencia)
	u.Nombre_competencia = Utils.RemoveAccents(u.Nombre_competencia)
}

func PutOneCompetenciaInput(u *RequestMessages.PutOneCompetenciaPayload) {
	u.Nombre_competencia = strings.TrimSpace(u.Nombre_competencia)
	u.Nombre_competencia = strings.ToUpper(u.Nombre_competencia)
	u.Nombre_competencia = Utils.RemoveAccents(u.Nombre_competencia)
}

func DeleteCompetenciaInput(u *RequestMessages.DeleteCompetenciaPayload) {

}

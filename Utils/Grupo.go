package Utils

import (
	"medroom-backend/Models"
)

func SearchIndexGrupoBySigla(lista_grupos []Models.Grupo, sigla_grupo string) (bool, int) {
	for i := 0; i < len(lista_grupos); i++ {
		if lista_grupos[i].Sigla_grupo == sigla_grupo {
			return true, i
		}
	}

	return false, -1
}

func SearchIdGrupoBySigla(lista_grupos []Models.Grupo, sigla_grupo string) (bool, int) {
	for i := 0; i < len(lista_grupos); i++ {
		if lista_grupos[i].Sigla_grupo == sigla_grupo {
			return true, lista_grupos[i].Id
		}
	}

	return false, -1
}

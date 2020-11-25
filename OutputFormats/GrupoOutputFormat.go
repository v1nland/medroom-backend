package OutputFormats

import (
	"medroom-backend/Models"
	"medroom-backend/ResponseMessages"
)

func GetGruposOutput(u []Models.Grupo) (output []ResponseMessages.ListGruposResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, ResponseMessages.ListGruposResponse{
			Nombre_grupo: u[i].Nombre_grupo,
			Sigla_grupo:  u[i].Sigla_grupo,
		})
	}

	return output
}

func GetOneGrupoOutput(u Models.Grupo) (output ResponseMessages.GetOneGrupoResponse) {
	return ResponseMessages.GetOneGrupoResponse{
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

func AddNewGrupoOutput(u Models.Grupo) (output ResponseMessages.AddNewGrupoResponse) {
	return ResponseMessages.AddNewGrupoResponse{
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

func PutOneGrupoOutput(u Models.Grupo) (output ResponseMessages.PutOneGrupoResponse) {
	return ResponseMessages.PutOneGrupoResponse{
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

func DeleteGrupoOutput(u Models.Grupo) (output ResponseMessages.DeleteGrupoResponse) {
	return ResponseMessages.DeleteGrupoResponse{
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

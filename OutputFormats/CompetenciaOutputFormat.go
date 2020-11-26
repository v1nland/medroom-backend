package OutputFormats

import (
	"medroom-backend/Models"
	"medroom-backend/ResponseMessages"
)

func GetCompetenciasOutput(u []Models.Competencia) (output []ResponseMessages.ListCompetenciasResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, ResponseMessages.ListCompetenciasResponse{
			Nombre_competencia: u[i].Nombre_competencia,
		})
	}

	return output
}

func GetOneCompetenciaOutput(u Models.Competencia) (output ResponseMessages.GetOneCompetenciaResponse) {
	return ResponseMessages.GetOneCompetenciaResponse{
		Nombre_competencia: u.Nombre_competencia,
	}
}

func AddNewCompetenciaOutput(u Models.Competencia) (output ResponseMessages.AddNewCompetenciaResponse) {
	return ResponseMessages.AddNewCompetenciaResponse{
		Nombre_competencia: u.Nombre_competencia,
	}
}

func PutOneCompetenciaOutput(u Models.Competencia) (output ResponseMessages.PutOneCompetenciaResponse) {
	return ResponseMessages.PutOneCompetenciaResponse{
		Nombre_competencia: u.Nombre_competencia,
	}
}

func DeleteCompetenciaOutput(u Models.Competencia) (output ResponseMessages.DeleteCompetenciaResponse) {
	return ResponseMessages.DeleteCompetenciaResponse{
		Nombre_competencia: u.Nombre_competencia,
	}
}

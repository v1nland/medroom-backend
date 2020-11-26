package OutputFormats

import (
	"medroom-backend/Models"
	"medroom-backend/ResponseMessages"
)

func GetPuntajesOutput(u []Models.Puntaje) (output []ResponseMessages.ListPuntajesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, ResponseMessages.ListPuntajesResponse{
			// Evaluacion_puntaje:   GetOneEvaluacionOutput(u[i].Evaluacion_puntaje),
			Competencia_puntaje:  GetOneCompetenciaOutput(u[i].Competencia_puntaje),
			Calificacion_puntaje: u[i].Calificacion_puntaje,
			Nivel_logro_puntaje:  u[i].Nivel_logro_puntaje,
		})
	}

	return output
}

func GetOnePuntajeOutput(u Models.Puntaje) (output ResponseMessages.GetOnePuntajeResponse) {
	return ResponseMessages.GetOnePuntajeResponse{
		// Evaluacion_puntaje:   GetOneEvaluacionOutput(u.Evaluacion_puntaje),
		Competencia_puntaje:  GetOneCompetenciaOutput(u.Competencia_puntaje),
		Calificacion_puntaje: u.Calificacion_puntaje,
		Nivel_logro_puntaje:  u.Nivel_logro_puntaje,
	}
}

func AddNewPuntajeOutput(u Models.Puntaje) (output ResponseMessages.AddNewPuntajeResponse) {
	return ResponseMessages.AddNewPuntajeResponse{
		Id_evaluacion:        u.Id_evaluacion,
		Id_competencia:       u.Id_competencia,
		Calificacion_puntaje: u.Calificacion_puntaje,
		Nivel_logro_puntaje:  u.Nivel_logro_puntaje,
	}
}

func PutOnePuntajeOutput(u Models.Puntaje) (output ResponseMessages.PutOnePuntajeResponse) {
	return ResponseMessages.PutOnePuntajeResponse{
		Id_evaluacion:        u.Id_evaluacion,
		Id_competencia:       u.Id_competencia,
		Calificacion_puntaje: u.Calificacion_puntaje,
		Nivel_logro_puntaje:  u.Nivel_logro_puntaje,
	}
}

func DeletePuntajeOutput(u Models.Puntaje) (output ResponseMessages.DeletePuntajeResponse) {
	return ResponseMessages.DeletePuntajeResponse{
		Id_evaluacion:        u.Id_evaluacion,
		Id_competencia:       u.Id_competencia,
		Calificacion_puntaje: u.Calificacion_puntaje,
		Nivel_logro_puntaje:  u.Nivel_logro_puntaje,
	}
}

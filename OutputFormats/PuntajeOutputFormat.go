package OutputFormats

import (
	"medroom-backend/Models"
	"medroom-backend/ResponseMessages"
)

func GetPuntajesOutput(u []Models.Puntaje) (output []ResponseMessages.ListPuntajesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, ResponseMessages.ListPuntajesResponse{
			Competencia_puntaje:  GetOneCompetenciaOutput(u[i].Competencia_puntaje),
			Calificacion_puntaje: u[i].Calificacion_puntaje,
			Feedback_puntaje:     u[i].Feedback_puntaje,
		})
	}

	return output
}

func GetOnePuntajeOutput(u Models.Puntaje) (output ResponseMessages.GetOnePuntajeResponse) {
	return ResponseMessages.GetOnePuntajeResponse{
		Competencia_puntaje:  GetOneCompetenciaOutput(u.Competencia_puntaje),
		Calificacion_puntaje: u.Calificacion_puntaje,
		Feedback_puntaje:     u.Feedback_puntaje,
	}
}

func AddNewPuntajeOutput(u Models.Puntaje) (output ResponseMessages.AddNewPuntajeResponse) {
	return ResponseMessages.AddNewPuntajeResponse{
		Id_competencia:       u.Id_competencia,
		Calificacion_puntaje: u.Calificacion_puntaje,
		Feedback_puntaje:     u.Feedback_puntaje,
	}
}

func PutOnePuntajeOutput(u Models.Puntaje) (output ResponseMessages.PutOnePuntajeResponse) {
	return ResponseMessages.PutOnePuntajeResponse{
		Id_competencia:       u.Id_competencia,
		Calificacion_puntaje: u.Calificacion_puntaje,
		Feedback_puntaje:     u.Feedback_puntaje,
	}
}

func DeletePuntajeOutput(u Models.Puntaje) (output ResponseMessages.DeletePuntajeResponse) {
	return ResponseMessages.DeletePuntajeResponse{
		Id_competencia:       u.Id_competencia,
		Calificacion_puntaje: u.Calificacion_puntaje,
		Feedback_puntaje:     u.Feedback_puntaje,
	}
}

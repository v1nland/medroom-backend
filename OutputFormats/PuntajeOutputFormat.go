package OutputFormats

import (
	"medroom-backend/Models"
	"medroom-backend/ResponseMessages"
)

func GetPuntajesOutput(u []Models.Puntaje) (output []ResponseMessages.ListPuntajesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, ResponseMessages.ListPuntajesResponse{
			Id_evaluacion:              u[i].Id_evaluacion,
			Nombre_competencia_puntaje: u[i].Nombre_competencia_puntaje,
			Calificacion_puntaje:       u[i].Calificacion_puntaje,
			Feedback_puntaje:           u[i].Feedback_puntaje,
		})
	}

	return output
}

func GetOnePuntajeOutput(u Models.Puntaje) (output ResponseMessages.GetOnePuntajeResponse) {
	return ResponseMessages.GetOnePuntajeResponse{
		Id_evaluacion:              u.Id_evaluacion,
		Nombre_competencia_puntaje: u.Nombre_competencia_puntaje,
		Calificacion_puntaje:       u.Calificacion_puntaje,
		Feedback_puntaje:           u.Feedback_puntaje,
	}
}

func AddNewPuntajeOutput(u Models.Puntaje) (output ResponseMessages.AddNewPuntajeResponse) {
	return ResponseMessages.AddNewPuntajeResponse{
		Id_evaluacion:              u.Id_evaluacion,
		Nombre_competencia_puntaje: u.Nombre_competencia_puntaje,
		Calificacion_puntaje:       u.Calificacion_puntaje,
		Feedback_puntaje:           u.Feedback_puntaje,
	}
}

func PutOnePuntajeOutput(u Models.Puntaje) (output ResponseMessages.PutOnePuntajeResponse) {
	return ResponseMessages.PutOnePuntajeResponse{
		Id_evaluacion:              u.Id_evaluacion,
		Nombre_competencia_puntaje: u.Nombre_competencia_puntaje,
		Calificacion_puntaje:       u.Calificacion_puntaje,
		Feedback_puntaje:           u.Feedback_puntaje,
	}
}

func DeletePuntajeOutput(u Models.Puntaje) (output ResponseMessages.DeletePuntajeResponse) {
	return ResponseMessages.DeletePuntajeResponse{
		Id_evaluacion:              u.Id_evaluacion,
		Nombre_competencia_puntaje: u.Nombre_competencia_puntaje,
		Calificacion_puntaje:       u.Calificacion_puntaje,
		Feedback_puntaje:           u.Feedback_puntaje,
	}
}

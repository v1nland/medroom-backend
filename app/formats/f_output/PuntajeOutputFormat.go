package f_output

import (
	"medroom-backend/app/Messages/Response"
	"medroom-backend/app/models"
)

func ListPuntajes(u []models.Puntaje) (output []Response.ListPuntajesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListPuntajesResponse{
			// Id_evaluacion:              u[i].Id_evaluacion,
			// Nombre_competencia_puntaje: u[i].Nombre_competencia_puntaje,
			// Codigo_competencia_puntaje: u[i].Codigo_competencia_puntaje,
			// Calificacion_puntaje:       u[i].Calificacion_puntaje,
			// Feedback_puntaje:           u[i].Feedback_puntaje,
		})
	}

	return output
}

func GetOnePuntaje(u models.Puntaje) (output Response.GetOnePuntajeResponse) {
	return Response.GetOnePuntajeResponse{
		// Id_evaluacion:              u.Id_evaluacion,
		// Nombre_competencia_puntaje: u.Nombre_competencia_puntaje,
		// Codigo_competencia_puntaje: u.Codigo_competencia_puntaje,
		// Calificacion_puntaje:       u.Calificacion_puntaje,
		// Feedback_puntaje:           u.Feedback_puntaje,
	}
}

func AddNewPuntaje(u models.Puntaje) (output Response.AddNewPuntajeResponse) {
	return Response.AddNewPuntajeResponse{
		Id: u.Id,
	}
}

func PutOnePuntaje(u models.Puntaje) (output Response.PutOnePuntajeResponse) {
	return Response.PutOnePuntajeResponse{
		Id: u.Id,
	}
}

func DeletePuntaje(u models.Puntaje) (output Response.DeletePuntajeResponse) {
	return Response.DeletePuntajeResponse{
		Id: u.Id,
	}
}

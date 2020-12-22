package Output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/Models"
)

func GetPuntajesOutput(u []Models.Puntaje) (output []Response.ListPuntajesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListPuntajesResponse{
			Id_evaluacion:              u[i].Id_evaluacion,
			Nombre_competencia_puntaje: u[i].Nombre_competencia_puntaje,
			Codigo_competencia_puntaje: u[i].Codigo_competencia_puntaje,
			Calificacion_puntaje:       u[i].Calificacion_puntaje,
			Feedback_puntaje:           u[i].Feedback_puntaje,
		})
	}

	return output
}

func GetOnePuntajeOutput(u Models.Puntaje) (output Response.GetOnePuntajeResponse) {
	return Response.GetOnePuntajeResponse{
		Id_evaluacion:              u.Id_evaluacion,
		Nombre_competencia_puntaje: u.Nombre_competencia_puntaje,
		Codigo_competencia_puntaje: u.Codigo_competencia_puntaje,
		Calificacion_puntaje:       u.Calificacion_puntaje,
		Feedback_puntaje:           u.Feedback_puntaje,
	}
}

func AddNewPuntajeOutput(u Models.Puntaje) (output Response.AddNewPuntajeResponse) {
	return Response.AddNewPuntajeResponse{
		Id_evaluacion:              u.Id_evaluacion,
		Nombre_competencia_puntaje: u.Nombre_competencia_puntaje,
		Codigo_competencia_puntaje: u.Codigo_competencia_puntaje,
		Calificacion_puntaje:       u.Calificacion_puntaje,
		Feedback_puntaje:           u.Feedback_puntaje,
	}
}

func PutOnePuntajeOutput(u Models.Puntaje) (output Response.PutOnePuntajeResponse) {
	return Response.PutOnePuntajeResponse{
		Id_evaluacion:              u.Id_evaluacion,
		Nombre_competencia_puntaje: u.Nombre_competencia_puntaje,
		Codigo_competencia_puntaje: u.Codigo_competencia_puntaje,
		Calificacion_puntaje:       u.Calificacion_puntaje,
		Feedback_puntaje:           u.Feedback_puntaje,
	}
}

func DeletePuntajeOutput(u Models.Puntaje) (output Response.DeletePuntajeResponse) {
	return Response.DeletePuntajeResponse{
		Id_evaluacion:              u.Id_evaluacion,
		Nombre_competencia_puntaje: u.Nombre_competencia_puntaje,
		Codigo_competencia_puntaje: u.Codigo_competencia_puntaje,
		Calificacion_puntaje:       u.Calificacion_puntaje,
		Feedback_puntaje:           u.Feedback_puntaje,
	}
}

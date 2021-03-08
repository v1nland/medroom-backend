package Output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/models"
)

func ListGruposOutput(u []models.Grupo) (output []Response.ListGruposResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListGruposResponse{
			Id_curso: u[i].Id_curso,
			// Evaluador_grupo: GetOneEvaluadorOutput(u[i].Evaluador_grupo),
			// Estudiantes_grupo: GetEstudiantesOutput(u[i].Estudiantes_grupo),
			Nombre_grupo: u[i].Nombre_grupo,
			Sigla_grupo:  u[i].Sigla_grupo,
		})
	}

	return output
}

func GetOneGrupoOutput(u models.Grupo) (output Response.GetOneGrupoResponse) {
	return Response.GetOneGrupoResponse{
		Id_curso: u.Id_curso,
		// Evaluador_grupo: GetOneEvaluadorOutput(u.Evaluador_grupo),
		// Estudiantes_grupo: GetEstudiantesOutput(u.Estudiantes_grupo),
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

func GetOneGrupoEstudianteOutput(u models.Grupo) (output Response.GetOneGrupoEstudianteResponse) {
	return Response.GetOneGrupoEstudianteResponse{
		Id_curso: u.Id_curso,
		// Evaluador_grupo: GetOneEvaluadorOutput(u.Evaluador_grupo),
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

func GetGruposEstudianteOutput(u []models.Grupo) (output []Response.GetGruposEstudianteResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.GetGruposEstudianteResponse{
			Id_curso: u[i].Id_curso,
			// Evaluador_grupo: GetOneEvaluadorOutput(u[i].Evaluador_grupo),
			Nombre_grupo: u[i].Nombre_grupo,
			Sigla_grupo:  u[i].Sigla_grupo,
		})
	}

	return output
}

func GetOneGrupoEvaluadorOutput(u models.Grupo) (output Response.GetOneGrupoEvaluadorResponse) {
	return Response.GetOneGrupoEvaluadorResponse{
		Id_curso: u.Id_curso,
		// Evaluador_grupo: GetOneEvaluadorOutput(u.Evaluador_grupo),
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

func GetGruposEvaluadorOutput(u []models.Grupo) (output []Response.GetGruposEvaluadorResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.GetGruposEvaluadorResponse{
			Id_curso: u[i].Id_curso,
			// Evaluador_grupo: GetOneEvaluadorOutput(u[i].Evaluador_grupo),
			Nombre_grupo: u[i].Nombre_grupo,
			Sigla_grupo:  u[i].Sigla_grupo,
		})
	}

	return output
}

func AddNewGrupoOutput(u models.Grupo) (output Response.AddNewGrupoResponse) {
	return Response.AddNewGrupoResponse{
		Id: u.Id,
	}
}

func PutOneGrupoOutput(u models.Grupo) (output Response.PutOneGrupoResponse) {
	return Response.PutOneGrupoResponse{
		Id: u.Id,
	}
}

func DeleteGrupoOutput(u models.Grupo) (output Response.DeleteGrupoResponse) {
	return Response.DeleteGrupoResponse{
		Id: u.Id,
	}
}

package Output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/Models"
)

func ListGruposOutput(u []Models.Grupo) (output []Response.ListGruposResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListGruposResponse{
			Id_curso:        u[i].Id_curso,
			Evaluador_grupo: GetOneEvaluadorOutput(u[i].Evaluador_grupo),
			// Estudiantes_grupo: GetEstudiantesOutput(u[i].Estudiantes_grupo),
			Nombre_grupo: u[i].Nombre_grupo,
			Sigla_grupo:  u[i].Sigla_grupo,
		})
	}

	return output
}

func GetOneGrupoOutput(u Models.Grupo) (output Response.GetOneGrupoResponse) {
	return Response.GetOneGrupoResponse{
		Id_curso:        u.Id_curso,
		Evaluador_grupo: GetOneEvaluadorOutput(u.Evaluador_grupo),
		// Estudiantes_grupo: GetEstudiantesOutput(u.Estudiantes_grupo),
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

func GetGrupoEstudianteOutput(u Models.Grupo) (output Response.GetGrupoEstudianteResponse) {
	return Response.GetGrupoEstudianteResponse{
		Id_curso:        u.Id_curso,
		Evaluador_grupo: GetOneEvaluadorOutput(u.Evaluador_grupo),
		// Estudiantes_grupo: GetEstudiantesOutput(u.Estudiantes_grupo),
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

func GetGrupoEvaluadorOutput(u Models.Grupo) (output Response.GetGrupoEvaluadorResponse) {
	return Response.GetGrupoEvaluadorResponse{
		Id_curso:        u.Id_curso,
		Evaluador_grupo: GetOneEvaluadorOutput(u.Evaluador_grupo),
		// Estudiantes_grupo: GetEstudiantesOutput(u.Estudiantes_grupo),
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

func AddNewGrupoOutput(u Models.Grupo) (output Response.AddNewGrupoResponse) {
	return Response.AddNewGrupoResponse{
		Id_curso:     u.Id_curso,
		Id_evaluador: u.Id_evaluador,
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

func PutOneGrupoOutput(u Models.Grupo) (output Response.PutOneGrupoResponse) {
	return Response.PutOneGrupoResponse{
		Id_curso:     u.Id_curso,
		Id_evaluador: u.Id_evaluador,
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

func DeleteGrupoOutput(u Models.Grupo) (output Response.DeleteGrupoResponse) {
	return Response.DeleteGrupoResponse{
		Id_curso:     u.Id_curso,
		Id_evaluador: u.Id_evaluador,
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

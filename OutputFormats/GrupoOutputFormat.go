package OutputFormats

import (
	"medroom-backend/Models"
	"medroom-backend/ResponseMessages"
)

func GetGruposOutput(u []Models.Grupo) (output []ResponseMessages.ListGruposResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, ResponseMessages.ListGruposResponse{
			Id_curso:          u[i].Id_curso,
			Evaluador_grupo:   GetOneEvaluadorOutput(u[i].Evaluador_grupo),
			Estudiantes_grupo: GetEstudiantesOutput(u[i].Estudiantes_grupo),
			Nombre_grupo:      u[i].Nombre_grupo,
			Sigla_grupo:       u[i].Sigla_grupo,
		})
	}

	return output
}

func GetOneGrupoOutput(u Models.Grupo) (output ResponseMessages.GetOneGrupoResponse) {
	return ResponseMessages.GetOneGrupoResponse{
		Id_curso:          u.Id_curso,
		Evaluador_grupo:   GetOneEvaluadorOutput(u.Evaluador_grupo),
		Estudiantes_grupo: GetEstudiantesOutput(u.Estudiantes_grupo),
		Nombre_grupo:      u.Nombre_grupo,
		Sigla_grupo:       u.Sigla_grupo,
	}
}

func GetGrupoEstudianteOutput(u Models.Grupo) (output ResponseMessages.GetGrupoEstudianteResponse) {
	return ResponseMessages.GetGrupoEstudianteResponse{
		Id_curso:          u.Id_curso,
		Evaluador_grupo:   GetOneEvaluadorOutput(u.Evaluador_grupo),
		Estudiantes_grupo: GetEstudiantesOutput(u.Estudiantes_grupo),
		Nombre_grupo:      u.Nombre_grupo,
		Sigla_grupo:       u.Sigla_grupo,
	}
}

func GetGrupoEvaluadorOutput(u Models.Grupo) (output ResponseMessages.GetGrupoEvaluadorResponse) {
	return ResponseMessages.GetGrupoEvaluadorResponse{
		Id_curso:          u.Id_curso,
		Evaluador_grupo:   GetOneEvaluadorOutput(u.Evaluador_grupo),
		Estudiantes_grupo: GetEstudiantesOutput(u.Estudiantes_grupo),
		Nombre_grupo:      u.Nombre_grupo,
		Sigla_grupo:       u.Sigla_grupo,
	}
}

func AddNewGrupoOutput(u Models.Grupo) (output ResponseMessages.AddNewGrupoResponse) {
	return ResponseMessages.AddNewGrupoResponse{
		Id_curso:     u.Id_curso,
		Id_evaluador: u.Id_evaluador,
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

func PutOneGrupoOutput(u Models.Grupo) (output ResponseMessages.PutOneGrupoResponse) {
	return ResponseMessages.PutOneGrupoResponse{
		Id_curso:     u.Id_curso,
		Id_evaluador: u.Id_evaluador,
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

func DeleteGrupoOutput(u Models.Grupo) (output ResponseMessages.DeleteGrupoResponse) {
	return ResponseMessages.DeleteGrupoResponse{
		Id_curso:     u.Id_curso,
		Id_evaluador: u.Id_evaluador,
		Nombre_grupo: u.Nombre_grupo,
		Sigla_grupo:  u.Sigla_grupo,
	}
}

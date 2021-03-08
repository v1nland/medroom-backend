package f_output

import (
	"medroom-backend/app/Messages/Response"
	"medroom-backend/app/models"
)

func ListEstudiantes(u []models.Estudiante) (output []Response.ListEstudiantesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListEstudiantesResponse{
			Id:                            u[i].Id,
			Rol_estudiante:                GetOneRol(u[i].Rol_estudiante),
			Evaluaciones_estudiante:       ListEvaluacionesEstudiante(u[i].Calificaciones_estudiante),
			Rut_estudiante:                u[i].Rut_estudiante,
			Nombres_estudiante:            u[i].Nombres_estudiante,
			Apellidos_estudiante:          u[i].Apellidos_estudiante,
			Correo_electronico_estudiante: u[i].Correo_electronico_estudiante,
			Telefono_fijo_estudiante:      u[i].Telefono_fijo_estudiante,
			Telefono_celular_estudiante:   u[i].Telefono_celular_estudiante,
		})
	}

	return output
}

func GetOneEstudiante(u models.Estudiante) (output Response.GetOneEstudianteResponse) {
	return Response.GetOneEstudianteResponse{
		Rol_estudiante:                GetOneRol(u.Rol_estudiante),
		Evaluaciones_estudiante:       ListEvaluacionesEstudiante(u.Calificaciones_estudiante),
		Rut_estudiante:                u.Rut_estudiante,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}

func GetMyEstudiante(u models.Estudiante) (output Response.GetMyEstudianteResponse) {
	return Response.GetMyEstudianteResponse{
		Rol_estudiante:                GetOneRol(u.Rol_estudiante),
		Evaluaciones_estudiante:       ListEvaluacionesEstudiante(u.Calificaciones_estudiante),
		Rut_estudiante:                u.Rut_estudiante,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}

func AddNewEstudiante(u models.Estudiante) (output Response.AddNewEstudianteResponse) {
	return Response.AddNewEstudianteResponse{
		Id: u.Id,
	}
}

func PutOneEstudiante(u models.Estudiante) (output Response.PutOneEstudianteResponse) {
	return Response.PutOneEstudianteResponse{
		Id: u.Id,
	}
}

func DeleteEstudiante(u models.Estudiante) (output Response.DeleteEstudianteResponse) {
	return Response.DeleteEstudianteResponse{
		Id: u.Id,
	}
}

func PutMyEstudiante(u models.Estudiante) (output Response.PutMyEstudianteResponse) {
	return Response.PutMyEstudianteResponse{
		Id: u.Id,
	}
}

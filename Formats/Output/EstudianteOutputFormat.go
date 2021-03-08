package Output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/models"
)

func ListEstudiantesOutput(u []models.Estudiante) (output []Response.ListEstudiantesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListEstudiantesResponse{
			Id:                            u[i].Id,
			Rol_estudiante:                GetOneRolOutput(u[i].Rol_estudiante),
			Evaluaciones_estudiante:       ListEvaluacionesEstudianteOutput(u[i].Calificaciones_estudiante),
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

func GetOneEstudianteOutput(u models.Estudiante) (output Response.GetOneEstudianteResponse) {
	return Response.GetOneEstudianteResponse{
		Rol_estudiante:                GetOneRolOutput(u.Rol_estudiante),
		Evaluaciones_estudiante:       ListEvaluacionesEstudianteOutput(u.Calificaciones_estudiante),
		Rut_estudiante:                u.Rut_estudiante,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}

func GetMyEstudianteOutput(u models.Estudiante) (output Response.GetMyEstudianteResponse) {
	return Response.GetMyEstudianteResponse{
		Rol_estudiante:                GetOneRolOutput(u.Rol_estudiante),
		Evaluaciones_estudiante:       ListEvaluacionesEstudianteOutput(u.Calificaciones_estudiante),
		Rut_estudiante:                u.Rut_estudiante,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}

func AddNewEstudianteOutput(u models.Estudiante) (output Response.AddNewEstudianteResponse) {
	return Response.AddNewEstudianteResponse{
		Id: u.Id,
	}
}

func PutOneEstudianteOutput(u models.Estudiante) (output Response.PutOneEstudianteResponse) {
	return Response.PutOneEstudianteResponse{
		Id: u.Id,
	}
}

func DeleteEstudianteOutput(u models.Estudiante) (output Response.DeleteEstudianteResponse) {
	return Response.DeleteEstudianteResponse{
		Id: u.Id,
	}
}

func PutMyEstudianteOutput(u models.Estudiante) (output Response.PutMyEstudianteResponse) {
	return Response.PutMyEstudianteResponse{
		Id: u.Id,
	}
}

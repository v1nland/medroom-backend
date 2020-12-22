package Output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/Models"
)

func GetEstudiantesOutput(u []Models.Estudiante) (output []Response.ListEstudiantesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListEstudiantesResponse{
			Id:                      u[i].Id,
			Rol_estudiante:          GetOneRolOutput(u[i].Rol_estudiante),
			Evaluaciones_estudiante: GetEvaluacionesEstudianteOutput(u[i].Evaluaciones_estudiante),
			// Id_grupo:                      u[i].Id_grupo,
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

func GetOneEstudianteOutput(u Models.Estudiante) (output Response.GetOneEstudianteResponse) {
	return Response.GetOneEstudianteResponse{
		Rol_estudiante:          GetOneRolOutput(u.Rol_estudiante),
		Evaluaciones_estudiante: GetEvaluacionesEstudianteOutput(u.Evaluaciones_estudiante),
		// Id_grupo:                      u.Id_grupo,
		Rut_estudiante:                u.Rut_estudiante,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}

func AddNewEstudianteOutput(u Models.Estudiante) (output Response.AddNewEstudianteResponse) {
	return Response.AddNewEstudianteResponse{
		Id_rol: u.Id_rol,
		// Id_grupo:                      u.Id_grupo,
		Rut_estudiante:                u.Rut_estudiante,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}

func PutOneEstudianteOutput(u Models.Estudiante) (output Response.PutOneEstudianteResponse) {
	return Response.PutOneEstudianteResponse{
		Id_rol: u.Id_rol,
		// Id_grupo:                      u.Id_grupo,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Hash_contrasena_estudiante:    u.Hash_contrasena_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}

func DeleteEstudianteOutput(u Models.Estudiante) (output Response.DeleteEstudianteResponse) {
	return Response.DeleteEstudianteResponse{
		Id_rol: u.Id_rol,
		// Id_grupo:                      u.Id_grupo,
		Rut_estudiante:                u.Rut_estudiante,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}

func GetMyEstudianteOutput(u Models.Estudiante) (output Response.GetMyEstudianteResponse) {
	return Response.GetMyEstudianteResponse{
		Rol_estudiante:          GetOneRolOutput(u.Rol_estudiante),
		Evaluaciones_estudiante: GetEvaluacionesEstudianteOutput(u.Evaluaciones_estudiante),
		// Id_grupo:                      u.Id_grupo,
		Rut_estudiante:                u.Rut_estudiante,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}

func PutMyEstudianteOutput(u Models.Estudiante) (output Response.PutMyEstudianteResponse) {
	return Response.PutMyEstudianteResponse{
		Id_rol: u.Id_rol,
		// Id_grupo:                      u.Id_grupo,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Hash_contrasena_estudiante:    u.Hash_contrasena_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}

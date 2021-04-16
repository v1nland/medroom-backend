package f_output

import (
	"medroom-backend/messages/Response"
	"medroom-backend/models"
)

func ListEvaluadores(u []models.Evaluador) (output []Response.ListEvaluadoresResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListEvaluadoresResponse{
			Rol_evaluador:                GetOneRol(u[i].Rol_evaluador),
			Rut_evaluador:                u[i].Rut_evaluador,
			Nombres_evaluador:            u[i].Nombres_evaluador,
			Apellidos_evaluador:          u[i].Apellidos_evaluador,
			Correo_electronico_evaluador: u[i].Correo_electronico_evaluador,
			Telefono_fijo_evaluador:      u[i].Telefono_fijo_evaluador,
			Telefono_celular_evaluador:   u[i].Telefono_celular_evaluador,
			Recinto_evaluador:            u[i].Recinto_evaluador,
			Cargo_evaluador:              u[i].Cargo_evaluador,
		})
	}

	return output
}

func GetOneEvaluador(u models.Evaluador) (output Response.GetOneEvaluadorResponse) {
	return Response.GetOneEvaluadorResponse{
		Rol_evaluador:                GetOneRol(u.Rol_evaluador),
		Rut_evaluador:                u.Rut_evaluador,
		Nombres_evaluador:            u.Nombres_evaluador,
		Apellidos_evaluador:          u.Apellidos_evaluador,
		Correo_electronico_evaluador: u.Correo_electronico_evaluador,
		Telefono_fijo_evaluador:      u.Telefono_fijo_evaluador,
		Telefono_celular_evaluador:   u.Telefono_celular_evaluador,
		Recinto_evaluador:            u.Recinto_evaluador,
		Cargo_evaluador:              u.Cargo_evaluador,
	}
}

func GetMyEvaluador(u models.Evaluador) (output Response.GetMyEvaluadorResponse) {
	return Response.GetMyEvaluadorResponse{
		Rol_evaluador:                GetOneRol(u.Rol_evaluador),
		Rut_evaluador:                u.Rut_evaluador,
		Nombres_evaluador:            u.Nombres_evaluador,
		Apellidos_evaluador:          u.Apellidos_evaluador,
		Correo_electronico_evaluador: u.Correo_electronico_evaluador,
		Telefono_fijo_evaluador:      u.Telefono_fijo_evaluador,
		Telefono_celular_evaluador:   u.Telefono_celular_evaluador,
		Recinto_evaluador:            u.Recinto_evaluador,
		Cargo_evaluador:              u.Cargo_evaluador,
	}
}

func AddNewEvaluador(u models.Evaluador) (output Response.AddNewEvaluadorResponse) {
	return Response.AddNewEvaluadorResponse{
		Id: u.Id,
	}
}

func PutOneEvaluador(u models.Evaluador) (output Response.PutOneEvaluadorResponse) {
	return Response.PutOneEvaluadorResponse{
		Id: u.Id,
	}
}

func DeleteEvaluador(u models.Evaluador) (output Response.DeleteEvaluadorResponse) {
	return Response.DeleteEvaluadorResponse{
		Id: u.Id,
	}
}

func PutMyEvaluador(u models.Evaluador) (output Response.PutMyEvaluadorResponse) {
	return Response.PutMyEvaluadorResponse{
		Id: u.Id,
	}
}

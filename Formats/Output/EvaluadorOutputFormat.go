package Output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/Models"
)

func ListEvaluadoresOutput(u []Models.Evaluador) (output []Response.ListEvaluadoresResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListEvaluadoresResponse{
			Rol_evaluador:                GetOneRolOutput(u[i].Rol_evaluador),
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

func GetOneEvaluadorOutput(u Models.Evaluador) (output Response.GetOneEvaluadorResponse) {
	return Response.GetOneEvaluadorResponse{
		Rol_evaluador:                GetOneRolOutput(u.Rol_evaluador),
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

func GetMyEvaluadorOutput(u Models.Evaluador) (output Response.GetMyEvaluadorResponse) {
	return Response.GetMyEvaluadorResponse{
		Rol_evaluador:                GetOneRolOutput(u.Rol_evaluador),
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

func AddNewEvaluadorOutput(u Models.Evaluador) (output Response.AddNewEvaluadorResponse) {
	return Response.AddNewEvaluadorResponse{
		Id: u.Id,
	}
}

func PutOneEvaluadorOutput(u Models.Evaluador) (output Response.PutOneEvaluadorResponse) {
	return Response.PutOneEvaluadorResponse{
		Id: u.Id,
	}
}

func DeleteEvaluadorOutput(u Models.Evaluador) (output Response.DeleteEvaluadorResponse) {
	return Response.DeleteEvaluadorResponse{
		Id: u.Id,
	}
}

func PutMyEvaluadorOutput(u Models.Evaluador) (output Response.PutMyEvaluadorResponse) {
	return Response.PutMyEvaluadorResponse{
		Id: u.Id,
	}
}

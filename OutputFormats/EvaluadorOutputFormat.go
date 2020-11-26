package OutputFormats

import (
	"medroom-backend/Models"
	"medroom-backend/ResponseMessages"
)

func GetEvaluadorsOutput(u []Models.Evaluador) (output []ResponseMessages.ListEvaluadorsResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, ResponseMessages.ListEvaluadorsResponse{
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

func GetOneEvaluadorOutput(u Models.Evaluador) (output ResponseMessages.GetOneEvaluadorResponse) {
	return ResponseMessages.GetOneEvaluadorResponse{
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

func AddNewEvaluadorOutput(u Models.Evaluador) (output ResponseMessages.AddNewEvaluadorResponse) {
	return ResponseMessages.AddNewEvaluadorResponse{
		Id_rol:                       u.Id_rol,
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

func PutOneEvaluadorOutput(u Models.Evaluador) (output ResponseMessages.PutOneEvaluadorResponse) {
	return ResponseMessages.PutOneEvaluadorResponse{
		Id_rol:                       u.Id_rol,
		Rut_evaluador:                u.Rut_evaluador,
		Nombres_evaluador:            u.Nombres_evaluador,
		Apellidos_evaluador:          u.Apellidos_evaluador,
		Hash_contrasena_evaluador:    u.Hash_contrasena_evaluador,
		Correo_electronico_evaluador: u.Correo_electronico_evaluador,
		Telefono_fijo_evaluador:      u.Telefono_fijo_evaluador,
		Telefono_celular_evaluador:   u.Telefono_celular_evaluador,
		Recinto_evaluador:            u.Recinto_evaluador,
		Cargo_evaluador:              u.Cargo_evaluador,
	}
}

func DeleteEvaluadorOutput(u Models.Evaluador) (output ResponseMessages.DeleteEvaluadorResponse) {
	return ResponseMessages.DeleteEvaluadorResponse{
		Id_rol:                       u.Id_rol,
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

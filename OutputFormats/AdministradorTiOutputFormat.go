package OutputFormats

import (
	"medroom-backend/Models"
	"medroom-backend/ResponseMessages"
)

func GetAdministradoresTiOutput(u []Models.AdministradorTi) (output []ResponseMessages.ListAdministradoresTiResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, ResponseMessages.ListAdministradoresTiResponse{
			Rol_administrador_ti:                GetOneRolOutput(u[i].Rol_administrador_ti),
			Rut_administrador_ti:                u[i].Rut_administrador_ti,
			Nombres_administrador_ti:            u[i].Nombres_administrador_ti,
			Apellidos_administrador_ti:          u[i].Apellidos_administrador_ti,
			Correo_electronico_administrador_ti: u[i].Correo_electronico_administrador_ti,
			Telefono_fijo_administrador_ti:      u[i].Telefono_fijo_administrador_ti,
			Telefono_celular_administrador_ti:   u[i].Telefono_celular_administrador_ti,
		})
	}

	return output
}

func GetOneAdministradorTiOutput(u Models.AdministradorTi) (output ResponseMessages.GetOneAdministradorTiResponse) {
	return ResponseMessages.GetOneAdministradorTiResponse{
		Rol_administrador_ti:                GetOneRolOutput(u.Rol_administrador_ti),
		Rut_administrador_ti:                u.Rut_administrador_ti,
		Nombres_administrador_ti:            u.Nombres_administrador_ti,
		Apellidos_administrador_ti:          u.Apellidos_administrador_ti,
		Correo_electronico_administrador_ti: u.Correo_electronico_administrador_ti,
		Telefono_fijo_administrador_ti:      u.Telefono_fijo_administrador_ti,
		Telefono_celular_administrador_ti:   u.Telefono_celular_administrador_ti,
	}
}

func GetMyAdministradorTiOutput(u Models.AdministradorTi) (output ResponseMessages.GetMyAdministradorTiResponse) {
	return ResponseMessages.GetMyAdministradorTiResponse{
		Rol_administrador_ti:                GetOneRolOutput(u.Rol_administrador_ti),
		Rut_administrador_ti:                u.Rut_administrador_ti,
		Nombres_administrador_ti:            u.Nombres_administrador_ti,
		Apellidos_administrador_ti:          u.Apellidos_administrador_ti,
		Correo_electronico_administrador_ti: u.Correo_electronico_administrador_ti,
		Telefono_fijo_administrador_ti:      u.Telefono_fijo_administrador_ti,
		Telefono_celular_administrador_ti:   u.Telefono_celular_administrador_ti,
	}
}

func AddNewAdministradorTiOutput(u Models.AdministradorTi) (output ResponseMessages.AddNewAdministradorTiResponse) {
	return ResponseMessages.AddNewAdministradorTiResponse{
		Id_rol:                              u.Id_rol,
		Rut_administrador_ti:                u.Rut_administrador_ti,
		Nombres_administrador_ti:            u.Nombres_administrador_ti,
		Apellidos_administrador_ti:          u.Apellidos_administrador_ti,
		Correo_electronico_administrador_ti: u.Correo_electronico_administrador_ti,
		Telefono_fijo_administrador_ti:      u.Telefono_fijo_administrador_ti,
		Telefono_celular_administrador_ti:   u.Telefono_celular_administrador_ti,
	}
}

func PutOneAdministradorTiOutput(u Models.AdministradorTi) (output ResponseMessages.PutOneAdministradorTiResponse) {
	return ResponseMessages.PutOneAdministradorTiResponse{
		Id_rol:                              u.Id_rol,
		Nombres_administrador_ti:            u.Nombres_administrador_ti,
		Apellidos_administrador_ti:          u.Apellidos_administrador_ti,
		Hash_contrasena_administrador_ti:    u.Hash_contrasena_administrador_ti,
		Correo_electronico_administrador_ti: u.Correo_electronico_administrador_ti,
		Telefono_fijo_administrador_ti:      u.Telefono_fijo_administrador_ti,
		Telefono_celular_administrador_ti:   u.Telefono_celular_administrador_ti,
	}
}

func PutMyAdministradorTiOutput(u Models.AdministradorTi) (output ResponseMessages.PutMyAdministradorTiResponse) {
	return ResponseMessages.PutMyAdministradorTiResponse{
		Id_rol:                              u.Id_rol,
		Nombres_administrador_ti:            u.Nombres_administrador_ti,
		Apellidos_administrador_ti:          u.Apellidos_administrador_ti,
		Hash_contrasena_administrador_ti:    u.Hash_contrasena_administrador_ti,
		Correo_electronico_administrador_ti: u.Correo_electronico_administrador_ti,
		Telefono_fijo_administrador_ti:      u.Telefono_fijo_administrador_ti,
		Telefono_celular_administrador_ti:   u.Telefono_celular_administrador_ti,
	}
}

func DeleteAdministradorTiOutput(u Models.AdministradorTi) (output ResponseMessages.DeleteAdministradorTiResponse) {
	return ResponseMessages.DeleteAdministradorTiResponse{
		Id_rol:                              u.Id_rol,
		Rut_administrador_ti:                u.Rut_administrador_ti,
		Nombres_administrador_ti:            u.Nombres_administrador_ti,
		Apellidos_administrador_ti:          u.Apellidos_administrador_ti,
		Correo_electronico_administrador_ti: u.Correo_electronico_administrador_ti,
		Telefono_fijo_administrador_ti:      u.Telefono_fijo_administrador_ti,
		Telefono_celular_administrador_ti:   u.Telefono_celular_administrador_ti,
	}
}

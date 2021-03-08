package f_output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/models"
)

func ListAdministradoresTi(u []models.AdministradorTi) (output []Response.ListAdministradoresTiResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListAdministradoresTiResponse{
			Rol_administrador_ti:                GetOneRol(u[i].Rol_administrador_ti),
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

func GetOneAdministradorTi(u models.AdministradorTi) (output Response.GetOneAdministradorTiResponse) {
	return Response.GetOneAdministradorTiResponse{
		Rol_administrador_ti:                GetOneRol(u.Rol_administrador_ti),
		Rut_administrador_ti:                u.Rut_administrador_ti,
		Nombres_administrador_ti:            u.Nombres_administrador_ti,
		Apellidos_administrador_ti:          u.Apellidos_administrador_ti,
		Correo_electronico_administrador_ti: u.Correo_electronico_administrador_ti,
		Telefono_fijo_administrador_ti:      u.Telefono_fijo_administrador_ti,
		Telefono_celular_administrador_ti:   u.Telefono_celular_administrador_ti,
	}
}

func GetMyAdministradorTi(u models.AdministradorTi) (output Response.GetMyAdministradorTiResponse) {
	return Response.GetMyAdministradorTiResponse{
		Rol_administrador_ti:                GetOneRol(u.Rol_administrador_ti),
		Rut_administrador_ti:                u.Rut_administrador_ti,
		Nombres_administrador_ti:            u.Nombres_administrador_ti,
		Apellidos_administrador_ti:          u.Apellidos_administrador_ti,
		Correo_electronico_administrador_ti: u.Correo_electronico_administrador_ti,
		Telefono_fijo_administrador_ti:      u.Telefono_fijo_administrador_ti,
		Telefono_celular_administrador_ti:   u.Telefono_celular_administrador_ti,
	}
}

func AddNewAdministradorTi(u models.AdministradorTi) (output Response.AddNewAdministradorTiResponse) {
	return Response.AddNewAdministradorTiResponse{
		Id: u.Id,
	}
}

func PutOneAdministradorTi(u models.AdministradorTi) (output Response.PutOneAdministradorTiResponse) {
	return Response.PutOneAdministradorTiResponse{
		Id: u.Id,
	}
}

func PutMyAdministradorTi(u models.AdministradorTi) (output Response.PutMyAdministradorTiResponse) {
	return Response.PutMyAdministradorTiResponse{
		Id: u.Id,
	}
}

func DeleteAdministradorTi(u models.AdministradorTi) (output Response.DeleteAdministradorTiResponse) {
	return Response.DeleteAdministradorTiResponse{
		Id: u.Id,
	}
}

package f_output

import (
	"medroom-backend/messages/Response"
	"medroom-backend/models"
)

func ListAdministradoresAcademicos(u []models.AdministradorAcademico) (output []Response.ListAdministradoresAcademicosResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListAdministradoresAcademicosResponse{
			Rol_administrador_academico:                GetOneRol(u[i].Rol_administrador_academico),
			Rut_administrador_academico:                u[i].Rut_administrador_academico,
			Nombres_administrador_academico:            u[i].Nombres_administrador_academico,
			Apellidos_administrador_academico:          u[i].Apellidos_administrador_academico,
			Correo_electronico_administrador_academico: u[i].Correo_electronico_administrador_academico,
			Telefono_fijo_administrador_academico:      u[i].Telefono_fijo_administrador_academico,
			Telefono_celular_administrador_academico:   u[i].Telefono_celular_administrador_academico,
		})
	}

	return output
}

func GetOneAdministradorAcademico(u models.AdministradorAcademico) (output Response.GetOneAdministradorAcademicoResponse) {
	return Response.GetOneAdministradorAcademicoResponse{
		Rol_administrador_academico:                GetOneRol(u.Rol_administrador_academico),
		Rut_administrador_academico:                u.Rut_administrador_academico,
		Nombres_administrador_academico:            u.Nombres_administrador_academico,
		Apellidos_administrador_academico:          u.Apellidos_administrador_academico,
		Correo_electronico_administrador_academico: u.Correo_electronico_administrador_academico,
		Telefono_fijo_administrador_academico:      u.Telefono_fijo_administrador_academico,
		Telefono_celular_administrador_academico:   u.Telefono_celular_administrador_academico,
	}
}

func GetMyAdministradorAcademico(u models.AdministradorAcademico) (output Response.GetMyAdministradorAcademicoResponse) {
	return Response.GetMyAdministradorAcademicoResponse{
		Rol_administrador_academico:                GetOneRol(u.Rol_administrador_academico),
		Rut_administrador_academico:                u.Rut_administrador_academico,
		Nombres_administrador_academico:            u.Nombres_administrador_academico,
		Apellidos_administrador_academico:          u.Apellidos_administrador_academico,
		Correo_electronico_administrador_academico: u.Correo_electronico_administrador_academico,
		Telefono_fijo_administrador_academico:      u.Telefono_fijo_administrador_academico,
		Telefono_celular_administrador_academico:   u.Telefono_celular_administrador_academico,
	}
}

func AddNewAdministradorAcademico(u models.AdministradorAcademico) (output Response.AddNewAdministradorAcademicoResponse) {
	return Response.AddNewAdministradorAcademicoResponse{
		Id: u.Id,
	}
}

func PutOneAdministradorAcademico(u models.AdministradorAcademico) (output Response.PutOneAdministradorAcademicoResponse) {
	return Response.PutOneAdministradorAcademicoResponse{
		Id: u.Id,
	}
}

func PutMyAdministradorAcademico(u models.AdministradorAcademico) (output Response.PutMyAdministradorAcademicoResponse) {
	return Response.PutMyAdministradorAcademicoResponse{
		Id: u.Id,
	}
}

func DeleteAdministradorAcademico(u models.AdministradorAcademico) (output Response.DeleteAdministradorAcademicoResponse) {
	return Response.DeleteAdministradorAcademicoResponse{
		Id: u.Id,
	}
}

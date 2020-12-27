package Output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/Models"
)

func ListAdministradoresAcademicosOutput(u []Models.AdministradorAcademico) (output []Response.ListAdministradoresAcademicosResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListAdministradoresAcademicosResponse{
			Rol_administrador_academico:                GetOneRolOutput(u[i].Rol_administrador_academico),
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

func GetOneAdministradorAcademicoOutput(u Models.AdministradorAcademico) (output Response.GetOneAdministradorAcademicoResponse) {
	return Response.GetOneAdministradorAcademicoResponse{
		Rol_administrador_academico:                GetOneRolOutput(u.Rol_administrador_academico),
		Rut_administrador_academico:                u.Rut_administrador_academico,
		Nombres_administrador_academico:            u.Nombres_administrador_academico,
		Apellidos_administrador_academico:          u.Apellidos_administrador_academico,
		Correo_electronico_administrador_academico: u.Correo_electronico_administrador_academico,
		Telefono_fijo_administrador_academico:      u.Telefono_fijo_administrador_academico,
		Telefono_celular_administrador_academico:   u.Telefono_celular_administrador_academico,
	}
}

func GetMyAdministradorAcademicoOutput(u Models.AdministradorAcademico) (output Response.GetMyAdministradorAcademicoResponse) {
	return Response.GetMyAdministradorAcademicoResponse{
		Rol_administrador_academico:                GetOneRolOutput(u.Rol_administrador_academico),
		Rut_administrador_academico:                u.Rut_administrador_academico,
		Nombres_administrador_academico:            u.Nombres_administrador_academico,
		Apellidos_administrador_academico:          u.Apellidos_administrador_academico,
		Correo_electronico_administrador_academico: u.Correo_electronico_administrador_academico,
		Telefono_fijo_administrador_academico:      u.Telefono_fijo_administrador_academico,
		Telefono_celular_administrador_academico:   u.Telefono_celular_administrador_academico,
	}
}

func AddNewAdministradorAcademicoOutput(u Models.AdministradorAcademico) (output Response.AddNewAdministradorAcademicoResponse) {
	return Response.AddNewAdministradorAcademicoResponse{
		Id: u.Id,
	}
}

func PutOneAdministradorAcademicoOutput(u Models.AdministradorAcademico) (output Response.PutOneAdministradorAcademicoResponse) {
	return Response.PutOneAdministradorAcademicoResponse{
		Id: u.Id,
	}
}

func PutMyAdministradorAcademicoOutput(u Models.AdministradorAcademico) (output Response.PutMyAdministradorAcademicoResponse) {
	return Response.PutMyAdministradorAcademicoResponse{
		Id: u.Id,
	}
}

func DeleteAdministradorAcademicoOutput(u Models.AdministradorAcademico) (output Response.DeleteAdministradorAcademicoResponse) {
	return Response.DeleteAdministradorAcademicoResponse{
		Id: u.Id,
	}
}

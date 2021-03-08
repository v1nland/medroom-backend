package f_output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/models"
)

func ListRoles(u []models.Rol) (output []Response.ListRolesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListRolesResponse{
			Nombre_rol: u[i].Nombre_rol,
		})
	}

	return output
}

func GetOneRol(u models.Rol) (output Response.GetOneRolResponse) {
	return Response.GetOneRolResponse{
		Nombre_rol: u.Nombre_rol,
	}
}

func AddNewRol(u models.Rol) (output Response.AddNewRolResponse) {
	return Response.AddNewRolResponse{
		Id: u.Id,
	}
}

func PutOneRol(u models.Rol) (output Response.PutOneRolResponse) {
	return Response.PutOneRolResponse{
		Id: u.Id,
	}
}

func DeleteRol(u models.Rol) (output Response.DeleteRolResponse) {
	return Response.DeleteRolResponse{
		Id: u.Id,
	}
}

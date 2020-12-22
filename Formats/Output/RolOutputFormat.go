package Output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/Models"
)

func ListRolesOutput(u []Models.Rol) (output []Response.ListRolesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListRolesResponse{
			Nombre_rol: u[i].Nombre_rol,
		})
	}

	return output
}

func GetOneRolOutput(u Models.Rol) (output Response.GetOneRolResponse) {
	return Response.GetOneRolResponse{
		Nombre_rol: u.Nombre_rol,
	}
}

func AddNewRolOutput(u Models.Rol) (output Response.AddNewRolResponse) {
	return Response.AddNewRolResponse{
		Nombre_rol: u.Nombre_rol,
	}
}

func PutOneRolOutput(u Models.Rol) (output Response.PutOneRolResponse) {
	return Response.PutOneRolResponse{
		Nombre_rol: u.Nombre_rol,
	}
}

func DeleteRolOutput(u Models.Rol) (output Response.DeleteRolResponse) {
	return Response.DeleteRolResponse{
		Nombre_rol: u.Nombre_rol,
	}
}

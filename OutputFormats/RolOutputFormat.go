package OutputFormats

import (
	"medroom-backend/Models"
	"medroom-backend/ResponseMessages"
)

func GetRolsOutput(u []Models.Rol) (output []ResponseMessages.ListRolsResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, ResponseMessages.ListRolsResponse{
			Nombre_rol: u[i].Nombre_rol,
		})
	}

	return output
}

func GetOneRolOutput(u Models.Rol) (output ResponseMessages.GetOneRolResponse) {
	return ResponseMessages.GetOneRolResponse{
		Nombre_rol: u.Nombre_rol,
	}
}

func AddNewRolOutput(u Models.Rol) (output ResponseMessages.AddNewRolResponse) {
	return ResponseMessages.AddNewRolResponse{
		Nombre_rol: u.Nombre_rol,
	}
}

func PutOneRolOutput(u Models.Rol) (output ResponseMessages.PutOneRolResponse) {
	return ResponseMessages.PutOneRolResponse{
		Nombre_rol: u.Nombre_rol,
	}
}

func DeleteRolOutput(u Models.Rol) (output ResponseMessages.DeleteRolResponse) {
	return ResponseMessages.DeleteRolResponse{
		Nombre_rol: u.Nombre_rol,
	}
}

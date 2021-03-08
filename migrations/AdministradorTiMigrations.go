package migrations

import (
	"fmt"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
)

func AdministradorTimigrations() {
	fmt.Println("===== ADMINISTRADOR TI =====")

	var rol models.Rol
	if err := repositories.GetOneRol(&rol, "4"); err != nil {
		panic("ROL ADMINISTRADOR TI NO EXISTE")
	}

	container := &models.AdministradorTi{
		Rol_administrador_ti:                rol,
		Rut_administrador_ti:                "11.111.111-1",
		Nombres_administrador_ti:            "NOMBRE ADMINISTRADOR",
		Apellidos_administrador_ti:          "APELLIDO TI",
		Hash_contrasena_administrador_ti:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
		Correo_electronico_administrador_ti: "ADMINISTRADOR.TI@MAIL.UDP.CL",
		Telefono_fijo_administrador_ti:      "12345678",
		Telefono_celular_administrador_ti:   "12345678",
	}

	if err := repositories.AddNewAdministradorTi(container); err != nil {
		panic("NO SE PUDO MIGRAR ADMINISTRADOR TI")
	}

	utils.StructToString(container)
}

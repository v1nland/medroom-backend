package Migrations

import (
	"fmt"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"
)

func AdministradorTiMigrations() {
	fmt.Println("===== ADMINISTRADOR TI =====")

	var rol Models.Rol
	if err := Repositories.GetOneRol(&rol, "4"); err != nil {
		panic("ROL ADMINISTRADOR TI NO EXISTE")
	}

	container := &Models.AdministradorTi{
		Rol_administrador_ti:                rol,
		Rut_administrador_ti:                "11.111.111-1",
		Nombres_administrador_ti:            "NOMBRE ADMINISTRADOR",
		Apellidos_administrador_ti:          "APELLIDO TI",
		Hash_contrasena_administrador_ti:    "hash",
		Correo_electronico_administrador_ti: "administrador.ti@udp.cl",
		Telefono_fijo_administrador_ti:      "12345678",
		Telefono_celular_administrador_ti:   "12345678",
	}

	if err := Repositories.AddNewAdministradorTi(container); err != nil {
		panic("NO SE PUDO MIGRAR ADMINISTRADOR TI")
	}

	Utils.StructToString(container)
}

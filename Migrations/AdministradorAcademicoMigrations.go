package Migrations

import (
	"fmt"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"
)

func AdministradorAcademicoMigrations() {
	fmt.Println("===== ADMINISTRADOR ACADEMICO =====")

	var rol Models.Rol
	if err := Repositories.GetOneRol(&rol, "3"); err != nil {
		panic("ROL ADMINISTRADOR ACADÉMICO NO EXISTE")
	}

	container := &Models.AdministradorAcademico{
		Rol_administrador_academico:                rol,
		Rut_administrador_academico:                "44.444.444-4",
		Nombres_administrador_academico:            "NOMBRE ADMINISTRADOR",
		Apellidos_administrador_academico:          "APELLIDO ACADEMICO",
		Hash_contrasena_administrador_academico:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
		Correo_electronico_administrador_academico: "administrador.academico@udp.cl",
		Telefono_fijo_administrador_academico:      "12345678",
		Telefono_celular_administrador_academico:   "12345678",
	}

	if err := Repositories.AddNewAdministradorAcademico(container); err != nil {
		panic("NO SE PUDO MIGRAR ADMINISTRADOR ACADEMICO")
	}

	Utils.StructToString(container)
}

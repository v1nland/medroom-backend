package migrations

import (
	"fmt"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/google/uuid"
)

func AdministradorTi() {
	fmt.Println("===== ADMINISTRADOR TI =====")

	container := &models.AdministradorTi{
		Id:                                  uuid.MustParse("75895023-f9db-4ccb-9045-c3d858dde28c"),
		Id_rol:                              4,
		Rut_administrador_ti:                "11.111.111-1",
		Nombres_administrador_ti:            "ADMINISTRADOR",
		Apellidos_administrador_ti:          "TI",
		Hash_contrasena_administrador_ti:    "2df47540a33e717aa32af8ec5d89dd8034123a9e0411c6f03f54d646840fc0d4",
		Correo_electronico_administrador_ti: "ADMINISTRADOR.TI@MAIL.UDP.CL",
		Telefono_fijo_administrador_ti:      "12345678",
		Telefono_celular_administrador_ti:   "12345678",
	}

	if err := repositories.AddNewAdministradorTi(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR ADMINISTRADOR TI")
	}

	utils.StructToString(container)
}

package migrations

import (
	"fmt"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
)

func Competenciamigrations() {
	fmt.Println("===== COMPETENCIA =====")

	// Anamnesis
	container := &models.Competencia{
		Id:                 "ANAM",
		Nombre_competencia: "ANAMNESIS",
	}
	if err := repositories.AddNewCompetencia(container); err != nil {
		panic("NO SE PUDO MIGRAR COMPETENCIA 'ANAMNESIS'")
	}
	utils.StructToString(container)

	// Exploración física
	container = &models.Competencia{
		Id:                 "EXFI",
		Nombre_competencia: "EXPLORACION FISICA",
	}
	if err := repositories.AddNewCompetencia(container); err != nil {
		panic("NO SE PUDO MIGRAR COMPETENCIA 'EXPLORACION FISICA'")
	}
	utils.StructToString(container)

	// Profesionalismo
	container = &models.Competencia{
		Id:                 "PROF",
		Nombre_competencia: "PROFESIONALISMO",
	}
	if err := repositories.AddNewCompetencia(container); err != nil {
		panic("NO SE PUDO MIGRAR COMPETENCIA 'PROFESIONALISMO'")
	}
	utils.StructToString(container)

	// Juicio Clínico
	container = &models.Competencia{
		Id:                 "JUCL",
		Nombre_competencia: "JUICIO CLINICO",
	}
	if err := repositories.AddNewCompetencia(container); err != nil {
		panic("NO SE PUDO MIGRAR COMPETENCIA 'JUICIO CLINICO'")
	}
	utils.StructToString(container)

	// Habilidades Comunicativas
	container = &models.Competencia{
		Id:                 "HACO",
		Nombre_competencia: "HABILIDADES COMUNICATIVAS",
	}
	if err := repositories.AddNewCompetencia(container); err != nil {
		panic("NO SE PUDO MIGRAR COMPETENCIA 'HABILIDADES COMUNICATIVAS'")
	}
	utils.StructToString(container)

	// Organización y Eficiencia
	container = &models.Competencia{
		Id:                 "OREF",
		Nombre_competencia: "ORGANIZACION Y EFICIENCIA",
	}
	if err := repositories.AddNewCompetencia(container); err != nil {
		panic("NO SE PUDO MIGRAR COMPETENCIA 'ORGANIZACION Y EFICIENCIA'")
	}
	utils.StructToString(container)
}

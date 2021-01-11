package Migrations

import (
	"fmt"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"
)

func CompetenciaMigrations() {
	fmt.Println("===== COMPETENCIA =====")

	// Anamnesis
	container := &Models.Competencia{
		Id:                 "ANAM",
		Nombre_competencia: "ANAMNESIS",
	}
	if err := Repositories.AddNewCompetencia(container); err != nil {
		panic("NO SE PUDO MIGRAR COMPETENCIA 'ANAMNESIS'")
	}
	Utils.StructToString(container)

	// Exploración física
	container = &Models.Competencia{
		Id:                 "EXFI",
		Nombre_competencia: "EXPLORACION FISICA",
	}
	if err := Repositories.AddNewCompetencia(container); err != nil {
		panic("NO SE PUDO MIGRAR COMPETENCIA 'EXPLORACION FISICA'")
	}
	Utils.StructToString(container)

	// Profesionalismo
	container = &Models.Competencia{
		Id:                 "PROF",
		Nombre_competencia: "PROFESIONALISMO",
	}
	if err := Repositories.AddNewCompetencia(container); err != nil {
		panic("NO SE PUDO MIGRAR COMPETENCIA 'PROFESIONALISMO'")
	}
	Utils.StructToString(container)

	// Juicio Clínico
	container = &Models.Competencia{
		Id:                 "JUCL",
		Nombre_competencia: "JUICIO CLINICO",
	}
	if err := Repositories.AddNewCompetencia(container); err != nil {
		panic("NO SE PUDO MIGRAR COMPETENCIA 'JUICIO CLINICO'")
	}
	Utils.StructToString(container)

	// Habilidades Comunicativas
	container = &Models.Competencia{
		Id:                 "HACO",
		Nombre_competencia: "HABILIDADES COMUNICATIVAS",
	}
	if err := Repositories.AddNewCompetencia(container); err != nil {
		panic("NO SE PUDO MIGRAR COMPETENCIA 'HABILIDADES COMUNICATIVAS'")
	}
	Utils.StructToString(container)

	// Organización y Eficiencia
	container = &Models.Competencia{
		Id:                 "OREF",
		Nombre_competencia: "ORGANIZACION Y EFICIENCIA",
	}
	if err := Repositories.AddNewCompetencia(container); err != nil {
		panic("NO SE PUDO MIGRAR COMPETENCIA 'ORGANIZACION Y EFICIENCIA'")
	}
	Utils.StructToString(container)

	// Valoración Global
	container = &Models.Competencia{
		Id:                 "VAGL",
		Nombre_competencia: "VALORACION GLOBAL",
	}
	if err := Repositories.AddNewCompetencia(container); err != nil {
		panic("NO SE PUDO MIGRAR COMPETENCIA 'VALORACION GLOBAL'")
	}
	Utils.StructToString(container)
}

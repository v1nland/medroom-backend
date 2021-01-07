package Repositories

import (
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// .Preload("Calificaciones_estudiante.Puntajes_calificacion_estudiante.Competencia_puntaje").Preload("Calificaciones_estudiante.Evaluacion_calificacion_estudiante").Preload("Calificaciones_estudiante.Periodo_calificacion_estudiante").Preload("Calificaciones_estudiante.Evaluador_calificacion_estudiante.Rol_evaluador")
func GetAllEstudiantes(u *[]Models.Estudiante) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Order("created_at asc").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneEstudiante(u *Models.Estudiante, id string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewEstudiante(u *Models.Estudiante) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneEstudiante(u *Models.Estudiante, id string) (err error) {
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Save(u)
	return nil
}

func DeleteEstudiante(u *Models.Estudiante, id string) (err error) {
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Delete(u)
	return nil
}

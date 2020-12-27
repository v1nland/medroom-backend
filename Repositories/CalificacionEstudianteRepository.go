package Repositories

import (
	"fmt"
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllCalificacionesEstudiante(u *[]Models.CalificacionEstudiante) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCalificacionEstudiante(u *Models.CalificacionEstudiante, id string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCalificacionEstudianteByIdEvaluacion(u *Models.CalificacionEstudiante, id_evaluacion string, id_estudiante string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id_evaluacion = ? AND id_estudiante = ?", id_evaluacion, id_estudiante).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewCalificacionEstudiante(u *Models.CalificacionEstudiante) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneCalificacionEstudiante(u *Models.CalificacionEstudiante, id string) (err error) {
	fmt.Println(u)
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Save(u)
	return nil
}

func DeleteCalificacionEstudiante(u *Models.CalificacionEstudiante, id string) (err error) {
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Delete(u)
	return nil
}

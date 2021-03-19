package repositories

import (
	"fmt"
	"medroom-backend/config"
	"medroom-backend/models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllEvaluaciones(u *[]models.Evaluacion) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneEvaluacion(u *models.Evaluacion, id string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewEvaluacion(u *models.Evaluacion) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneEvaluacion(u *models.Evaluacion, id string) (err error) {
	fmt.Println(u)
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Save(u)
	return nil
}

func DeleteEvaluacion(u *models.Evaluacion, id string) (err error) {
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Delete(u)
	return nil
}

func GetEvaluacionesRendidasEstudiante(u *[]models.Evaluacion, id_estudiante string) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Select("ev.*").Table("estudiantes e").Joins("JOIN calificaciones_estudiantes ce ON ce.id_estudiante = e.id").Joins("JOIN evaluaciones ev ON ce.id_evaluacion = ev.id").Where("e.id = ?", id_estudiante).Find(u).Error; err != nil {
		return err
	}
	return nil
}

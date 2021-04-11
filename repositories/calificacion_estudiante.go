package repositories

import (
	"fmt"
	"medroom-backend/config"
	"medroom-backend/models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllCalificacionesEstudiante(u *[]models.CalificacionEstudiante) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetOneCalificacionEstudiante(u *models.CalificacionEstudiante, id string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Where("id = ?", id).
		First(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetOneCalificacionEstudianteByIdEvaluacion(u *models.CalificacionEstudiante, id_evaluacion string, id_estudiante string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Preload("Puntajes_calificacion_estudiante.Competencia_puntaje").
		Where("id_evaluacion = ?", id_evaluacion).
		Where("id_estudiante = ?", id_estudiante).
		First(u).
		Error; err != nil {
		return err
	}
	return nil
}

func AddNewCalificacionEstudiante(u *models.CalificacionEstudiante) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Create(u).
		Error; err != nil {
		return err
	}
	return nil
}

func PutOneCalificacionEstudiante(u *models.CalificacionEstudiante, id string) (err error) {
	fmt.Println(u)
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Save(u)
	return nil
}

func DeleteCalificacionEstudiante(u *models.CalificacionEstudiante, id string) (err error) {
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Where("id = ?", id).
		Delete(u)
	return nil
}

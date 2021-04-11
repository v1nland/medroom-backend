package repositories

import (
	"fmt"
	"medroom-backend/config"
	"medroom-backend/models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ADMINISTRADOR ACADEMICO
func GetAllEvaluadoresCurso(u *[]models.Evaluador, id_curso string) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Order("created_at asc").
		Joins("JOIN evaluadores_grupos eg on eg.id_evaluador = evaluadores.id").
		Joins("JOIN grupos g on eg.id_grupo = g.id").
		Joins("join cursos c on g.id_curso = c.id").
		Where("c.id = ?", id_curso).
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetAllEvaluadoresCursoSinGrupo(u *[]models.Evaluador, id_curso string) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Order("created_at asc").
		Joins("JOIN evaluadores_grupos eg on eg.id_evaluador = evaluadores.id").
		Joins("JOIN grupos g on eg.id_grupo = g.id").
		Joins("join cursos c on g.id_curso = c.id").
		Where("g.sigla_grupo = 'SG' and c.id = ?", id_curso).
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetAllEvaluadores(u *[]models.Evaluador) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Order("created_at asc").
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetOneEvaluador(u *models.Evaluador, id string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Where("id = ?", id).
		First(u).
		Error; err != nil {
		return err
	}
	return nil
}

func AddNewEvaluador(u *models.Evaluador) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Create(u).
		Error; err != nil {
		return err
	}
	return nil
}

func PutOneEvaluador(u *models.Evaluador, id string) (err error) {
	fmt.Println(u)
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Save(u)
	return nil
}

func DeleteEvaluador(u *models.Evaluador, id string) (err error) {
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Where("id = ?", id).
		Delete(u)
	return nil
}

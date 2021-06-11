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
func GetAllEvaluadoresCurso(u *[]models.Evaluador, sigla_curso string, id_periodo string) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Order("created_at asc").
		Joins("JOIN evaluadores_grupos eg on eg.id_evaluador = evaluadores.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_grupo AND eg.sigla_curso = g.sigla_curso AND eg.id_periodo_curso = g.id_periodo_curso").
		Joins("JOIN cursos c ON g.sigla_curso = c.sigla_curso AND g.id_periodo_curso = c.id_periodo").
		Where("c.sigla_curso = ?", sigla_curso).
		Where("c.id_periodo = ?", id_periodo).
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetAllEvaluadoresCursoSinGrupo(u *[]models.Evaluador, sigla_curso string, id_periodo string) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Order("created_at asc").
		Joins("JOIN evaluadores_grupos eg on eg.id_evaluador = evaluadores.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_grupo AND eg.sigla_curso = g.sigla_curso AND eg.id_periodo_curso = g.id_periodo_curso").
		Joins("JOIN cursos c ON g.sigla_curso = c.sigla_curso AND g.id_periodo_curso = c.id_periodo").
		Where("g.sigla_grupo = 'SG'").
		Where("c.sigla_curso = ?", sigla_curso).
		Where("c.id_periodo = ?", id_periodo).
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

func GetOneEvaluadorByRut(u *models.Evaluador, rut string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Where("rut_evaluador = ?", rut).
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

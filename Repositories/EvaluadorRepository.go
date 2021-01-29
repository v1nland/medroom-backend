package Repositories

import (
	"fmt"
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllEvaluadores(u *[]Models.Evaluador) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Order("created_at asc").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetAllEvaluadoresCurso(u *[]Models.Evaluador, id_curso string) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Order("created_at asc").Joins("JOIN evaluadores_grupos eg on eg.id_evaluador = evaluadores.id").Joins("JOIN grupos g on eg.id_grupo = g.id").Joins("join cursos c on g.id_curso = c.id").Where("c.id = ?", id_curso).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetAllEvaluadoresCursoSinGrupo(u *[]Models.Evaluador, id_curso string) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Order("created_at asc").Joins("JOIN evaluadores_grupos eg on eg.id_evaluador = evaluadores.id").Joins("JOIN grupos g on eg.id_grupo = g.id").Joins("join cursos c on g.id_curso = c.id").Where("g.sigla_grupo = 'SG' and c.id = ?", id_curso).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneEvaluador(u *Models.Evaluador, id string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewEvaluador(u *Models.Evaluador) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneEvaluador(u *Models.Evaluador, id string) (err error) {
	fmt.Println(u)
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Save(u)
	return nil
}

func DeleteEvaluador(u *Models.Evaluador, id string) (err error) {
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Delete(u)
	return nil
}

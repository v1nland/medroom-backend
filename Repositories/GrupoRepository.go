package Repositories

import (
	"fmt"
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllGrupos(u *[]Models.Grupo) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupo(u *Models.Grupo, id string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetGruposEstudiante(u *[]Models.Grupo, id_curso string, id_estudiante string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("estudiantes").Select("g.*").Joins("JOIN estudiantes_grupos eg ON eg.id_estudiante = estudiantes.id").Joins("JOIN grupos g ON eg.id_grupo = g.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("c.id = ? AND estudiantes.id = ?", id_curso, id_estudiante).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupoEstudiante(u *Models.Grupo, id string, id_curso string, id_estudiante string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("estudiantes").Select("g.*").Joins("JOIN estudiantes_grupos eg ON eg.id_estudiante = estudiantes.id").Joins("JOIN grupos g ON eg.id_grupo = g.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("g.id = ? AND c.id = ? AND estudiantes.id = ?", id, id_curso, id_estudiante).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetGruposEvaluador(u *[]Models.Grupo, id_curso string, id_evaluador string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("evaluadores").Select("g.*").Joins("JOIN grupos g ON g.id_evaluador = evaluadores.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("c.id = ? AND evaluadores.id = ?", id_curso, id_evaluador).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupoEvaluador(u *Models.Grupo, id string, id_curso string, id_evaluador string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("evaluadores").Select("g.*").Joins("JOIN grupos g ON g.id_evaluador = evaluadores.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("g.id = ? AND c.id = ? AND evaluadores.id = ?", id, id_curso, id_evaluador).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewGrupo(u *Models.Grupo) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneGrupo(u *Models.Grupo, id string) (err error) {
	fmt.Println(u)
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Save(u)
	return nil
}

func DeleteGrupo(u *Models.Grupo, id string) (err error) {
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Delete(u)
	return nil
}

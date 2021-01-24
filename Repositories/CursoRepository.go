package Repositories

import (
	"fmt"
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllCursos(u *[]Models.Curso) (err error) {
	if err = Config.DB.Preload(clause.Associations).Order("created_at asc").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCurso(u *Models.Curso, id string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetCursosEstudiante(u *[]Models.Curso, id_estudiante string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("estudiantes").Select("c.*").Joins("JOIN estudiantes_grupos eg ON eg.id_estudiante = estudiantes.id").Joins("JOIN grupos g ON eg.id_grupo = g.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("estudiantes.id = ?", id_estudiante).Group("c.id").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCursoEstudiante(u *Models.Curso, id string, id_estudiante string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("estudiantes").Select("c.*").Joins("JOIN estudiantes_grupos eg ON eg.id_estudiante = estudiantes.id").Joins("JOIN grupos g ON eg.id_grupo = g.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("c.id = ? AND estudiantes.id = ?", id, id_estudiante).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetCursosEvaluador(u *[]Models.Curso, id_evaluador string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("evaluadores").Select("c.*").Joins("JOIN evaluadores_grupos eg ON eg.id_evaluador = evaluadores.id").Joins("JOIN grupos g ON eg.id_grupo = g.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("evaluadores.id = ?", id_evaluador).Group("c.id").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCursoEvaluador(u *Models.Curso, id string, id_evaluador string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("evaluadores").Select("c.*").Joins("JOIN evaluadores_grupos eg ON eg.id_evaluador = evaluadores.id").Joins("JOIN grupos g ON eg.id_grupo = g.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("c.id = ? AND evaluadores.id = ?", id, id_evaluador).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewCurso(u *Models.Curso) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneCurso(u *Models.Curso, id string) (err error) {
	fmt.Println(u)
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Save(u)
	return nil
}

func DeleteCurso(u *Models.Curso, id string) (err error) {
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Delete(u)
	return nil
}

func GetGrupoSinGrupo(u *Models.Curso, id string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Joins("JOIN grupos ON cursos.id = grupos.id_curso").First(u).Error; err != nil {
		return err
	}
	return nil
}

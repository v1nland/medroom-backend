package repositories

import (
	"fmt"
	"medroom-backend/config"
	"medroom-backend/models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ESTUDIANTE
func GetCursosEstudiante(u *[]models.Curso, id_estudiante string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("estudiantes").Select("c.*").Joins("JOIN estudiantes_grupos eg ON eg.id_estudiante = estudiantes.id").Joins("JOIN grupos g ON eg.id_grupo = g.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("estudiantes.id = ?", id_estudiante).Group("c.id").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCursoEstudiante(u *models.Curso, id string, id_estudiante string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("estudiantes").Select("c.*").Joins("JOIN estudiantes_grupos eg ON eg.id_estudiante = estudiantes.id").Joins("JOIN grupos g ON eg.id_grupo = g.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("c.id = ? AND estudiantes.id = ?", id, id_estudiante).First(u).Error; err != nil {
		return err
	}
	return nil
}

// EVALUADOR
func GetCursosEvaluador(u *[]models.Curso, id_evaluador string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("evaluadores").Select("c.*").Joins("JOIN evaluadores_grupos eg ON eg.id_evaluador = evaluadores.id").Joins("JOIN grupos g ON eg.id_grupo = g.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("evaluadores.id = ?", id_evaluador).Group("c.id").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCursoEvaluador(u *models.Curso, id string, id_evaluador string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("evaluadores").Select("c.*").Joins("JOIN evaluadores_grupos eg ON eg.id_evaluador = evaluadores.id").Joins("JOIN grupos g ON eg.id_grupo = g.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("c.id = ? AND evaluadores.id = ?", id, id_evaluador).First(u).Error; err != nil {
		return err
	}
	return nil
}

// ADMINISTRADOR ACADEMICO
func GetCursosAdministradorAcademico(u *[]models.Curso, id_administrador_academico string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("administradores_academicos").Select("c.*").Joins("JOIN administradores_academicos_cursos aac ON aac.id_administrador_academico = administradores_academicos.id").Joins("JOIN cursos c ON aac.id_curso = c.id").Where("administradores_academicos.id = ?", id_administrador_academico).Group("c.id").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCursoAdministradorAcademico(u *models.Curso, id string, id_administrador_academico string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("administradores_academicos").Select("c.*").Joins("JOIN administradores_academicos_cursos aac ON aac.id_administrador_academico = administradores_academicos.id").Joins("JOIN cursos c ON aac.id_curso = c.id").Where("c.id = ? AND administradores_academicos.id = ?", id, id_administrador_academico).First(u).Error; err != nil {
		return err
	}
	return nil
}

// ADMINISTRADOR TI
func GetAllCursos(u *[]models.Curso) (err error) {
	if err = config.DB.Preload(clause.Associations).Order("created_at asc").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCurso(u *models.Curso, id string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewCurso(u *models.Curso) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneCurso(u *models.Curso, id string) (err error) {
	fmt.Println(u)
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Save(u)
	return nil
}

func DeleteCurso(u *models.Curso, id string) (err error) {
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Delete(u)
	return nil
}

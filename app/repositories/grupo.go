package repositories

import (
	"fmt"
	"medroom-backend/app/config"
	"medroom-backend/app/models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllGrupos(u *[]models.Grupo) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Order("created_at asc").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupo(u *models.Grupo, id string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

// ESTUDIANTE
func GetGruposEstudiante(u *[]models.Grupo, id_curso string, id_estudiante string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("estudiantes").Select("g.*").Joins("JOIN estudiantes_grupos eg ON eg.id_estudiante = estudiantes.id").Joins("JOIN grupos g ON eg.id_grupo = g.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("c.id = ? AND estudiantes.id = ?", id_curso, id_estudiante).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupoEstudiante(u *models.Grupo, id string, id_curso string, id_estudiante string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("estudiantes").Select("g.*").Joins("JOIN estudiantes_grupos eg ON eg.id_estudiante = estudiantes.id").Joins("JOIN grupos g ON eg.id_grupo = g.id").Joins("JOIN cursos c ON g.id_curso = c.id").Where("g.id = ? AND c.id = ? AND estudiantes.id = ?", id, id_curso, id_estudiante).First(u).Error; err != nil {
		return err
	}
	return nil
}

// EVALUADOR
func GetGruposEvaluador(u *[]models.Grupo, id_curso string, id_evaluador string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("evaluadores").Select("g.*").Joins("JOIN evaluadores_grupos eg ON eg.id_evaluador = evaluadores.id").Joins("JOIN grupos g ON g.id = eg.id_grupo").Joins("JOIN cursos c ON g.id_curso = c.id").Where("c.id = ? AND evaluadores.id = ?", id_curso, id_evaluador).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupoEvaluador(u *models.Grupo, id string, id_curso string, id_evaluador string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("evaluadores").Select("g.*").Joins("JOIN evaluadores_grupos eg ON eg.id_evaluador = evaluadores.id").Joins("JOIN grupos g ON g.id = eg.id_grupo").Joins("JOIN cursos c ON g.id_curso = c.id").Where("g.id = ? AND c.id = ? AND evaluadores.id = ?", id, id_curso, id_evaluador).First(u).Error; err != nil {
		return err
	}
	return nil
}

// ADMINISTRADOR ACADEMICO
func GetGruposAdministradorAcademico(u *[]models.Grupo, id_curso string, id_administrador_academico string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("administradores_academicos").Select("g.*").Joins("JOIN administradores_academicos_cursos aac ON aac.id_administrador_academico = administradores_academicos.id").Joins("JOIN cursos c ON aac.id_curso = c.id").Joins("JOIN grupos g ON c.id = g.id_curso").Where("administradores_academicos.id = ?", id_administrador_academico).Group("g.id").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupoAdministradorAcademico(u *models.Grupo, id string, id_curso string, id_administrador_academico string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Table("administradores_academicos").Select("g.*").Joins("JOIN administradores_academicos_cursos aac ON aac.id_administrador_academico = administradores_academicos.id").Joins("JOIN cursos c ON aac.id_curso = c.id").Joins("JOIN grupos g ON c.id = g.id_curso").Where("g.id = ? AND administradores_academicos.id = ?", id_curso, id_administrador_academico).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewGrupo(u *models.Grupo) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneGrupo(u *models.Grupo, id string) (err error) {
	fmt.Println(u)
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Save(u)
	return nil
}

func DeleteGrupo(u *models.Grupo, id string) (err error) {
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Delete(u)
	return nil
}

func DeleteEstudianteGrupo(id_grupo string, id_estudiante string) (err error) {
	if err := config.DB.Debug().Exec(`DELETE FROM public.estudiantes_grupos WHERE estudiantes_grupos.id_grupo = ? AND estudiantes_grupos.id_estudiante = ?`, id_grupo, id_estudiante).Error; err != nil {
		return err
	}
	return nil
}

func DeleteEvaluadorGrupo(id_grupo string, id_evaluador string) (err error) {
	if err := config.DB.Debug().Exec(`DELETE FROM public.evaluadores_grupos WHERE evaluadores_grupos.id_grupo = ? AND evaluadores_grupos.id_evaluador = ?`, id_grupo, id_evaluador).Error; err != nil {
		return err
	}
	return nil
}

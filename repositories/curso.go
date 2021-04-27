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
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Table("estudiantes").
		Select("c.*").
		Joins("JOIN estudiantes_grupos eg ON eg.id_estudiante = estudiantes.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_grupo AND eg.sigla_curso = g.sigla_curso AND eg.id_periodo_curso = g.id_periodo_curso").
		Joins("JOIN cursos c ON g.sigla_curso = c.sigla_curso AND g.id_periodo_curso = c.id_periodo").
		Where("estudiantes.id = ?", id_estudiante).
		Group("c.sigla_curso, c.id_periodo").
		Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCursoEstudiante(u *models.Curso, sigla_curso string, id_periodo string, id_estudiante string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Select("c.*").
		Table("estudiantes").
		Joins("JOIN estudiantes_grupos eg ON eg.id_estudiante = estudiantes.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_grupo AND eg.sigla_curso = g.sigla_curso AND eg.id_periodo_curso = g.id_periodo_curso").
		Joins("JOIN cursos c ON g.sigla_curso = c.sigla_curso AND g.id_periodo_curso = c.id_periodo").
		Where("c.sigla_curso = ? AND c.id_periodo = ? AND estudiantes.id = ?", sigla_curso, id_periodo, id_estudiante).
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

// EVALUADOR
func GetCursosEvaluador(u *[]models.Curso, id_evaluador string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Table("evaluadores").
		Select("c.*").
		Joins("JOIN evaluadores_grupos eg ON eg.id_evaluador = evaluadores.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_grupo AND eg.sigla_curso = g.sigla_curso AND eg.id_periodo_curso = g.id_periodo_curso").
		Joins("JOIN cursos c ON g.sigla_curso = c.sigla_curso AND g.id_periodo_curso = c.id_periodo").
		Where("evaluadores.id = ?", id_evaluador).
		Group("c.sigla_curso, c.id_periodo").
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetOneCursoEvaluador(u *models.Curso, sigla_curso string, id_periodo string, id_evaluador string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Select("c.*").
		Table("evaluadores").
		Joins("JOIN evaluadores_grupos eg ON eg.id_evaluador = evaluadores.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_grupo AND eg.sigla_curso = g.sigla_curso AND eg.id_periodo_curso = g.id_periodo_curso").
		Joins("JOIN cursos c ON g.sigla_curso = c.sigla_curso AND g.id_periodo_curso = c.id_periodo").
		Where("c.sigla_curso = ? AND c.id_periodo = ? AND evaluadores.id = ?", sigla_curso, id_periodo, id_evaluador).
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

// ADMINISTRADOR ACADEMICO
func GetCursosAdministradorAcademico(u *[]models.Curso, id_administrador_academico string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Table("administradores_academicos").
		Select("c.*").
		Joins("JOIN administradores_academicos_cursos aac ON aac.id_administrador_academico = administradores_academicos.id").
		Joins("JOIN cursos c ON aac.sigla_curso = c.sigla_curso AND aac.id_periodo = c.id_periodo").
		Where("administradores_academicos.id = ?", id_administrador_academico).
		Group("c.sigla_curso,c.id_periodo").
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetOneCursoAdministradorAcademico(u *models.Curso, sigla_curso string, id_periodo string, id_administrador_academico string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Table("administradores_academicos").
		Select("c.*").
		Joins("JOIN administradores_academicos_cursos aac ON aac.id_administrador_academico = administradores_academicos.id").
		Joins("JOIN cursos c ON aac.sigla_curso = c.sigla_curso AND aac.id_periodo = c.id_periodo").
		Where("c.sigla = ? AND c.id_periodo = ? AND administradores_academicos.id = ?", sigla_curso, id_periodo, id_administrador_academico).
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

// ADMINISTRADOR TI
func GetAllCursos(u *[]models.Curso) (err error) {
	if err = config.DB.Preload(clause.Associations).
		Order("created_at asc").
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetOneCurso(u *models.Curso, sigla string, id_periodo string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Where("sigla_curso = ?", sigla).
		Where("id_periodo = ?", id_periodo).
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func AddNewCurso(u *models.Curso) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Create(u).
		Error; err != nil {
		return err
	}
	return nil
}

func PutOneCurso(u *models.Curso, sigla string, id_periodo string) (err error) {
	fmt.Println(u)
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Save(u)
	return nil
}

func DeleteCurso(u *models.Curso, sigla string, id_periodo string) (err error) {
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Where("sigla_curso = ?", sigla).
		Where("id_periodo = ?", id_periodo).
		Delete(u)
	return nil
}

func RemoveAdministradorAcademicoFromCurso(sigla string, id_periodo string, id_administrador_academico string) (err error) {
	if err := config.DB.Debug().
		Exec(`DELETE FROM public.administradores_academicos_cursos 
					WHERE administradores_academicos_cursos.sigla_curso = ? 
					WHERE administradores_academicos_cursos.id_periodo = ? 
					AND administradores_academicos_cursos.id_administrador_academico = ?`, sigla, id_periodo, id_administrador_academico).
		Error; err != nil {
		return err
	}
	return nil
}

func ClearCursoAssociations(u *models.Curso, sigla string, id_periodo string) (err error) {
	if err := config.DB.Debug().
		Exec(`DELETE FROM public.administradores_academicos_cursos 
					WHERE administradores_academicos_cursos.sigla_curso = ?
					AND administradores_academicos_cursos.id_periodo = ?`, sigla, id_periodo).
		Error; err != nil {
		return err
	}

	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Where("sigla_curso = ?", sigla).
		Where("id_periodo = ?", id_periodo).
		Delete(u)

	return nil
}

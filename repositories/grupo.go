package repositories

import (
	"fmt"
	"medroom-backend/config"
	"medroom-backend/models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllGrupos(u *[]models.Grupo) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Order("created_at asc").
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupo(u *models.Grupo, sigla_curso string, id_periodo_curso string, sigla_grupo string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Where("sigla_curso = ?", sigla_curso).
		Where("id_periodo_curso = ?", id_periodo_curso).
		Where("sigla_grupo = ?", sigla_grupo).
		First(u).
		Error; err != nil {
		return err
	}
	return nil
}

// ESTUDIANTE
func GetGruposEstudiante(u *[]models.Grupo, sigla_curso string, id_periodo string, id_estudiante string) (err error) {
	if err := config.DB.Debug().Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Table("estudiantes").
		Select("g.*").
		Joins("JOIN estudiantes_grupos eg ON eg.id_estudiante = estudiantes.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_grupo AND eg.sigla_curso = g.sigla_curso AND eg.id_periodo_curso = g.id_periodo_curso").
		Joins("JOIN cursos c ON g.sigla_curso = c.sigla_curso AND g.id_periodo_curso = c.id_periodo").
		Where("c.sigla_curso = ? AND c.id_periodo = ?", sigla_curso, id_periodo).
		Where("estudiantes.id = ?", id_estudiante).
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupoEstudiante(u *models.Grupo, sigla_grupo string, sigla_curso string, id_periodo string, id_estudiante string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Table("estudiantes").
		Select("g.*").
		Joins("JOIN estudiantes_grupos eg ON eg.id_estudiante = estudiantes.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_grupo AND eg.sigla_curso = g.sigla_curso AND eg.id_periodo_curso = g.id_periodo_curso").
		Joins("JOIN cursos c ON g.sigla_curso = c.sigla_curso AND g.id_periodo_curso = c.id_periodo").
		Where("estudiantes.id = ?", id_estudiante).
		Where("c.sigla_curso = ? AND c.id_periodo = ?", sigla_curso, id_periodo).
		Where("g.sigla_grupo = ? AND g.id_periodo_curso = ? AND g.sigla_curso", sigla_grupo, id_periodo, sigla_curso).
		First(u).
		Error; err != nil {
		return err
	}
	return nil
}

// EVALUADOR
func GetGruposEvaluador(u *[]models.Grupo, sigla_curso string, id_periodo string, id_evaluador string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Table("evaluadores").
		Select("g.*").
		Joins("JOIN evaluadores_grupos eg ON eg.id_evaluador = evaluadores.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_grupo AND eg.sigla_curso = g.sigla_curso AND eg.id_periodo_curso = g.id_periodo_curso").
		Joins("JOIN cursos c ON g.sigla_curso = c.sigla_curso AND g.id_periodo_curso = c.id_periodo").
		Where("evaluadores.id = ?", id_evaluador).
		Where("c.sigla_grupo = ? AND c.id_periodo = ?", sigla_curso, id_periodo).
		Where("g.sigla_grupo = ? AND g.id_periodo_curso = ?", sigla_curso, id_periodo).
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupoEvaluador(u *models.Grupo, sigla_grupo string, sigla_curso string, id_periodo string, id_evaluador string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Table("evaluadores").
		Select("g.*").
		Joins("JOIN evaluadores_grupos eg ON eg.id_evaluador = evaluadores.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_grupo AND eg.sigla_curso = g.sigla_curso AND eg.id_periodo_curso = g.id_periodo_curso").
		Joins("JOIN cursos c ON g.sigla_curso = c.sigla_curso AND g.id_periodo_curso = c.id_periodo").
		Where("evaluadores.id = ?", id_evaluador).
		Where("c.sigla_grupo = ? AND c.id_periodo = ?", sigla_curso, id_periodo).
		Where("g.sigla_grupo = ? AND g.id_periodo_curso = ?", sigla_grupo, id_periodo).
		First(u).
		Error; err != nil {
		return err
	}
	return nil
}

// ADMINISTRADOR ACADEMICO
func GetGruposAdministradorAcademico(u *[]models.Grupo, sigla_curso string, id_periodo string, id_administrador_academico string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Table("administradores_academicos").
		Select("g.*").
		Joins("JOIN administradores_academicos_cursos aac ON aac.id_administrador_academico = administradores_academicos.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_grupo AND eg.sigla_curso = g.sigla_curso AND eg.id_periodo_curso = g.id_periodo_curso").
		Joins("JOIN cursos c ON g.sigla_curso = c.sigla_curso AND g.id_periodo_curso = c.id_periodo").
		Where("administradores_academicos.id = ?", id_administrador_academico).
		Where("c.sigla_grupo = ? AND c.id_periodo = ?", sigla_curso, id_periodo).
		Where("g.sigla_grupo = ? AND g.id_periodo_curso = ?", sigla_curso, id_periodo).
		Group("g.sigla_curso, g.id_periodo_curso, g.sigla_grupo").
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupoAdministradorAcademico(u *models.Grupo, sigla_grupo string, sigla_curso string, id_periodo string, id_administrador_academico string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Table("administradores_academicos").
		Select("g.*").
		Joins("JOIN administradores_academicos_cursos aac ON aac.id_administrador_academico = administradores_academicos.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_grupo AND eg.sigla_curso = g.sigla_curso AND eg.id_periodo_curso = g.id_periodo_curso").
		Joins("JOIN cursos c ON g.sigla_curso = c.sigla_curso AND g.id_periodo_curso = c.id_periodo").
		Where("administradores_academicos.id = ?", id_administrador_academico).
		Where("c.sigla_grupo = ? AND c.id_periodo = ?", sigla_curso, id_periodo).
		Where("g.sigla_grupo = ? AND g.id_periodo_curso = ?", sigla_grupo, id_periodo).
		First(u).
		Error; err != nil {
		return err
	}
	return nil
}

func AddNewGrupo(u *models.Grupo) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Create(u).
		Error; err != nil {
		return err
	}
	return nil
}

func PutOneGrupo(u *models.Grupo, sigla_grupo string, sigla_curso string, id_periodo string) (err error) {
	fmt.Println(u)
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Save(u)
	return nil
}

func DeleteGrupo(u *models.Grupo) (err error) {
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Delete(u)
	return nil
}

func DeleteEstudianteGrupo(sigla_grupo string, sigla_curso string, id_periodo string, id_estudiante string) (err error) {
	if err := config.DB.Debug().
		Exec(`DELETE FROM public.estudiantes_grupos 
					WHERE estudiantes_grupos.sigla_curso = ? 
					AND estudiantes_grupos.id_periodo_curso = ?
					AND estudiantes_grupos.sigla_grupo = ?
					AND estudiantes_grupos.id_estudiante = ?`, sigla_curso, id_periodo, sigla_grupo, id_estudiante).
		Error; err != nil {
		return err
	}
	return nil
}

func DeleteEvaluadorGrupo(sigla_grupo string, sigla_curso string, id_periodo string, id_evaluador string) (err error) {
	if err := config.DB.Debug().
		Exec(`DELETE FROM public.evaluadores_grupos 
					WHERE estudiantes_grupos.sigla_curso = ? 
					AND estudiantes_grupos.id_periodo_curso = ?
					AND estudiantes_grupos.sigla_grupo = ?
					AND estudiantes_grupos.id_evaluador = ?`, sigla_curso, id_periodo, sigla_grupo, id_evaluador).
		Error; err != nil {
		return err
	}
	return nil
}

func ClearGrupo(sigla_grupo string, sigla_curso string, id_periodo string) (err error) {
	if err := config.DB.Debug().
		Exec(`DELETE FROM public.estudiantes_grupos 
					WHERE estudiantes_grupos.sigla_curso = ? 
					AND estudiantes_grupos.id_periodo_curso = ?
					AND estudiantes_grupos.sigla_grupo = ?`, sigla_curso, id_periodo, sigla_grupo).
		Error; err != nil {
		return err
	}
	if err := config.DB.Debug().
		Exec(`DELETE FROM public.evaluadores_grupos 
					WHERE estudiantes_grupos.sigla_curso = ? 
					AND estudiantes_grupos.id_periodo_curso = ?
					AND estudiantes_grupos.sigla_grupo = ?`, sigla_curso, id_periodo, sigla_grupo).
		Error; err != nil {
		return err
	}
	return nil
}

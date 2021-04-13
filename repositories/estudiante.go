package repositories

import (
	"medroom-backend/config"
	"medroom-backend/models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ADMINISTRADOR ACADEMICO
func GetAllEstudiantesCurso(u *[]models.Estudiante, sigla_curso string, id_periodo string) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Order("created_at asc").
		Joins("JOIN estudiantes_grupos eg on eg.id_estudiante = estudiantes.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_curso AND eg.sigla_curso = g.id_periodo_curso AND eg.id_periodo_curso = g.sigla_grupo").
		Joins("JOIN cursos c ON g.sigla_curso = c.sigla_curso AND g.id_periodo_curso = c.id_periodo").
		Where("g.sigla_grupo != 'SG'").
		Where("c.sigla_curso = ?", sigla_curso).
		Where("c.id_periodo = ?", id_periodo).
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetAllEstudiantesCursoSinGrupo(u *[]models.Estudiante, sigla_curso string, id_periodo string) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Order("created_at asc").
		Joins("JOIN estudiantes_grupos eg on eg.id_estudiante = estudiantes.id").
		Joins("JOIN grupos g ON eg.sigla_grupo = g.sigla_curso AND eg.sigla_curso = g.id_periodo_curso AND eg.id_periodo_curso = g.sigla_grupo").
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

func GetAllEstudiantes(u *[]models.Estudiante) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Order("created_at asc").
		Find(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetOneEstudiante(u *models.Estudiante, id string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Where("id = ?", id).
		First(u).
		Error; err != nil {
		return err
	}
	return nil
}

func GetReporteEstudiante(u *models.Estudiante, id string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Preload("Calificaciones_estudiante.Puntajes_calificacion_estudiante.Competencia_puntaje").
		Preload("Calificaciones_estudiante.Evaluador_calificacion_estudiante").
		Preload("Calificaciones_estudiante.Evaluacion_calificacion_estudiante").
		Where("id = ?", id).
		First(u).
		Error; err != nil {
		return err
	}
	return nil
}

func AddNewEstudiante(u *models.Estudiante) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Create(u).
		Error; err != nil {
		return err
	}
	return nil
}

func PutOneEstudiante(u *models.Estudiante, id string) (err error) {
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Save(u)
	return nil
}

func DeleteEstudiante(u *models.Estudiante, id string) (err error) {
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Preload(clause.Associations).
		Where("id = ?", id).
		Delete(u)
	return nil
}

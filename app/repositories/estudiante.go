package repositories

import (
	"medroom-backend/app/config"
	"medroom-backend/app/models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ADMINISTRADOR ACADEMICO
func GetAllEstudiantesCurso(u *[]models.Estudiante, id_curso string) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Order("created_at asc").Joins("JOIN estudiantes_grupos eg on eg.id_estudiante = estudiantes.id").Joins("JOIN grupos g on eg.id_grupo = g.id").Joins("join cursos c on g.id_curso = c.id").Where("g.sigla_grupo != 'SG' and c.id = ?", id_curso).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetAllEstudiantesCursoSinGrupo(u *[]models.Estudiante, id_curso string) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Order("created_at asc").Joins("JOIN estudiantes_grupos eg on eg.id_estudiante = estudiantes.id").Joins("JOIN grupos g on eg.id_grupo = g.id").Joins("join cursos c on g.id_curso = c.id").Where("g.sigla_grupo = 'SG' and c.id = ?", id_curso).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetAllEstudiantes(u *[]models.Estudiante) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Order("created_at asc").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneEstudiante(u *models.Estudiante, id string) (err error) {
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewEstudiante(u *models.Estudiante) (err error) {
	if err = config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneEstudiante(u *models.Estudiante, id string) (err error) {
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Save(u)
	return nil
}

func DeleteEstudiante(u *models.Estudiante, id string) (err error) {
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Delete(u)
	return nil
}

package Repositories

import (
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllEstudiantes(u *[]Models.Estudiante) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Order("created_at asc").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetAllEstudiantesCurso(u *[]Models.Estudiante, id_curso string) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Order("created_at asc").Joins("JOIN estudiantes_grupos eg on eg.id_estudiante = estudiantes.id").Joins("JOIN grupos g on eg.id_grupo = g.id").Joins("join cursos c on g.id_curso = c.id").Where("g.sigla_grupo != 'SG' and c.id = ?", id_curso).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetAllEstudiantesCursoSinGrupo(u *[]Models.Estudiante, id_curso string) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Order("created_at asc").Joins("JOIN estudiantes_grupos eg on eg.id_estudiante = estudiantes.id").Joins("JOIN grupos g on eg.id_grupo = g.id").Joins("join cursos c on g.id_curso = c.id").Where("g.sigla_grupo = 'SG' and c.id = ?", id_curso).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneEstudiante(u *Models.Estudiante, id string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewEstudiante(u *Models.Estudiante) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneEstudiante(u *Models.Estudiante, id string) (err error) {
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Save(u)
	return nil
}

func DeleteEstudiante(u *Models.Estudiante, id string) (err error) {
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Delete(u)
	return nil
}

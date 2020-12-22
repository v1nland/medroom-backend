package Repositories

import (
	"fmt"
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllEstudiantes(u *[]Models.Estudiante) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Find(u).Error; err != nil {
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
	fmt.Println(u)
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Save(u)
	return nil
}

func DeleteEstudiante(u *Models.Estudiante, id string) (err error) {
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Delete(u)
	return nil
}

func AuthEstudiante(u *Models.Estudiante, correo_electronico_estudiante string, hash_contrasena_estudiante string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("correo_electronico_estudiante = ? AND hash_contrasena_estudiante = ?", correo_electronico_estudiante, hash_contrasena_estudiante).First(u).Error; err != nil {
		return err
	}
	return nil
}

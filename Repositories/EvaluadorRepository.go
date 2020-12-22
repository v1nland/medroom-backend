package Repositories

import (
	"fmt"
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllEvaluadores(u *[]Models.Evaluador) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneEvaluador(u *Models.Evaluador, id string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewEvaluador(u *Models.Evaluador) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneEvaluador(u *Models.Evaluador, id string) (err error) {
	fmt.Println(u)
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Save(u)
	return nil
}

func DeleteEvaluador(u *Models.Evaluador, id string) (err error) {
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Delete(u)
	return nil
}

func AuthEvaluador(u *Models.Evaluador, correo_electronico_evaluador string, hash_contrasena_evaluador string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("correo_electronico_evaluador = ? AND hash_contrasena_evaluador = ?", correo_electronico_evaluador, hash_contrasena_evaluador).First(u).Error; err != nil {
		return err
	}
	return nil
}

package Repositories

import (
	"fmt"
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllAdministradoresTi(u *[]Models.AdministradorTi) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneAdministradorTi(u *Models.AdministradorTi, id string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewAdministradorTi(u *Models.AdministradorTi) (err error) {
	if err = Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneAdministradorTi(u *Models.AdministradorTi, id string) (err error) {
	fmt.Println(u)
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Save(u)
	return nil
}

func DeleteAdministradorTi(u *Models.AdministradorTi, id string) (err error) {
	Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("id = ?", id).Delete(u)
	return nil
}

func AuthAdministradorTi(u *Models.AdministradorTi, correo_electronico_administrador_ti string, hash_contrasena_administrador_ti string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("correo_electronico_administrador_ti = ? AND hash_contrasena_administrador_ti = ?", correo_electronico_administrador_ti, hash_contrasena_administrador_ti).First(u).Error; err != nil {
		return err
	}
	return nil
}

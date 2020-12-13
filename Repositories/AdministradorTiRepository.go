package Repositories

import (
	"fmt"
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetAllAdministradoresTi(u *[]Models.AdministradorTi) (err error) {
	if err = Config.DB.Set("gorm:auto_preload", true).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneAdministradorTi(u *Models.AdministradorTi, id string) (err error) {
	if err := Config.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewAdministradorTi(u *Models.AdministradorTi) (err error) {
	if err = Config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneAdministradorTi(u *Models.AdministradorTi, id string) (err error) {
	fmt.Println(u)
	Config.DB.Save(u)
	return nil
}

func DeleteAdministradorTi(u *Models.AdministradorTi, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(u)
	return nil
}

func AuthAdministradorTi(u *Models.AdministradorTi, correo_electronico_administrador_ti string, hash_contrasena_administrador_ti string) (err error) {
	if err := Config.DB.Set("gorm:auto_preload", true).Where("correo_electronico_administrador_ti = ? AND hash_contrasena_administrador_ti = ?", correo_electronico_administrador_ti, hash_contrasena_administrador_ti).First(u).Error; err != nil {
		return err
	}
	return nil
}

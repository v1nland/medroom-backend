package Repositories

import (
	"fmt"
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetAllAdministradoresAcademicos(u *[]Models.AdministradorAcademico) (err error) {
	if err = Config.DB.Set("gorm:auto_preload", true).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneAdministradorAcademico(u *Models.AdministradorAcademico, id string) (err error) {
	if err := Config.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewAdministradorAcademico(u *Models.AdministradorAcademico) (err error) {
	if err = Config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneAdministradorAcademico(u *Models.AdministradorAcademico, id string) (err error) {
	fmt.Println(u)
	Config.DB.Save(u)
	return nil
}

func DeleteAdministradorAcademico(u *Models.AdministradorAcademico, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(u)
	return nil
}

func AuthAdministradorAcademico(u *Models.AdministradorAcademico, correo_electronico_administrador_academico string, hash_contrasena_administrador_academico string) (err error) {
	if err := Config.DB.Set("gorm:auto_preload", true).Where("correo_electronico_administrador_academico = ? AND hash_contrasena_administrador_academico = ?", correo_electronico_administrador_academico, hash_contrasena_administrador_academico).First(u).Error; err != nil {
		return err
	}
	return nil
}

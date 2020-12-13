package Repositories

import (
	"fmt"
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetAllEstudiantes(u *[]Models.Estudiante) (err error) {
	if err = Config.DB.Set("gorm:auto_preload", true).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneEstudiante(u *Models.Estudiante, id string) (err error) {
	if err := Config.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewEstudiante(u *Models.Estudiante) (err error) {
	if err = Config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneEstudiante(u *Models.Estudiante, id string) (err error) {
	fmt.Println(u)
	Config.DB.Save(u)
	return nil
}

func DeleteEstudiante(u *Models.Estudiante, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(u)
	return nil
}

func AuthEstudiante(u *Models.Estudiante, correo_electronico_estudiante string, hash_contrasena_estudiante string) (err error) {
	if err := Config.DB.Set("gorm:auto_preload", true).Where("correo_electronico_estudiante = ? AND hash_contrasena_estudiante = ?", correo_electronico_estudiante, hash_contrasena_estudiante).First(u).Error; err != nil {
		return err
	}
	return nil
}

package Repositories

import (
	"fmt"
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetAllEvaluadores(u *[]Models.Evaluador) (err error) {
	if err = Config.DB.Set("gorm:auto_preload", true).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneEvaluador(u *Models.Evaluador, id string) (err error) {
	if err := Config.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewEvaluador(u *Models.Evaluador) (err error) {
	if err = Config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneEvaluador(u *Models.Evaluador, id string) (err error) {
	fmt.Println(u)
	Config.DB.Save(u)
	return nil
}

func DeleteEvaluador(u *Models.Evaluador, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(u)
	return nil
}

func AuthEvaluador(u *Models.Evaluador, correo_electronico_evaluador string, hash_contrasena_evaluador string) (err error) {
	if err := Config.DB.Set("gorm:auto_preload", true).Where("correo_electronico_evaluador = ? AND hash_contrasena_evaluador = ?", correo_electronico_evaluador, hash_contrasena_evaluador).First(u).Error; err != nil {
		return err
	}
	return nil
}

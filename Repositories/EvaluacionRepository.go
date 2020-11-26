package Repositories

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"medroom-backend/Config"
	"medroom-backend/Models"
)

func GetAllEvaluacions(u *[]Models.Evaluacion) (err error) {
	if err = Config.DB.Debug().Set("gorm:auto_preload", true).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneEvaluacion(u *Models.Evaluacion, id string) (err error) {
	if err := Config.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewEvaluacion(u *Models.Evaluacion) (err error) {
	if err = Config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneEvaluacion(u *Models.Evaluacion, id string) (err error) {
	fmt.Println(u)
	Config.DB.Save(u)
	return nil
}

func DeleteEvaluacion(u *Models.Evaluacion, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(u)
	return nil
}

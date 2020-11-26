package Repositories

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"medroom-backend/Config"
	"medroom-backend/Models"
)

func GetAllPuntajes(u *[]Models.Puntaje) (err error) {
	if err = Config.DB.Set("gorm:auto_preload", true).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOnePuntaje(u *Models.Puntaje, id string) (err error) {
	if err := Config.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewPuntaje(u *Models.Puntaje) (err error) {
	if err = Config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOnePuntaje(u *Models.Puntaje, id string) (err error) {
	fmt.Println(u)
	Config.DB.Save(u)
	return nil
}

func DeletePuntaje(u *Models.Puntaje, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(u)
	return nil
}

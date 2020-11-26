package Repositories

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"medroom-backend/Config"
	"medroom-backend/Models"
)

func GetAllPeriodos(u *[]Models.Periodo) (err error) {
	if err = Config.DB.Set("gorm:auto_preload", true).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOnePeriodo(u *Models.Periodo, id string) (err error) {
	if err := Config.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewPeriodo(u *Models.Periodo) (err error) {
	if err = Config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOnePeriodo(u *Models.Periodo, id string) (err error) {
	fmt.Println(u)
	Config.DB.Save(u)
	return nil
}

func DeletePeriodo(u *Models.Periodo, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(u)
	return nil
}

package Repositories

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"medroom-backend/Config"
	"medroom-backend/Models"
)

func GetAllGrupos(u *[]Models.Grupo) (err error) {
	if err = Config.DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupo(u *Models.Grupo, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewGrupo(u *Models.Grupo) (err error) {
	if err = Config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneGrupo(u *Models.Grupo, id string) (err error) {
	fmt.Println(u)
	Config.DB.Save(u)
	return nil
}

func DeleteGrupo(u *Models.Grupo, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(u)
	return nil
}

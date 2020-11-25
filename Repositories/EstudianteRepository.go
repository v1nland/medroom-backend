package Repositories

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"medroom-backend/Config"
	"medroom-backend/Models"
)

func GetAllEstudiantes(u *[]Models.Estudiante) (err error) {
	if err = Config.DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneEstudiante(u *Models.Estudiante, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(u).Error; err != nil {
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

package Repositories

import (
	"fmt"
	"medroom-backend/Config"
	"medroom-backend/Models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetAllGrupos(u *[]Models.Grupo) (err error) {
	if err = Config.DB.Set("gorm:auto_preload", true).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupo(u *Models.Grupo, id string) (err error) {
	if err := Config.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupoByEvaluadorId(u *Models.Grupo, id_evaluador string) (err error) {
	if err := Config.DB.Set("gorm:auto_preload", true).Where("id_evaluador = ?", id_evaluador).First(u).Error; err != nil {
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

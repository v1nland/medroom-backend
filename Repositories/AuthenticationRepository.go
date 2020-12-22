package Repositories

import (
	"medroom-backend/Config"
	"medroom-backend/Models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AuthenticateEstudiante(u *Models.Estudiante, correo_electronico_estudiante string, hash_contrasena_estudiante string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("correo_electronico_estudiante = ? AND hash_contrasena_estudiante = ?", correo_electronico_estudiante, hash_contrasena_estudiante).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AuthenticateEvaluador(u *Models.Evaluador, correo_electronico_evaluador string, hash_contrasena_evaluador string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("correo_electronico_evaluador = ? AND hash_contrasena_evaluador = ?", correo_electronico_evaluador, hash_contrasena_evaluador).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AuthenticateAdministradorAcademico(u *Models.AdministradorAcademico, correo_electronico_administrador_academico string, hash_contrasena_administrador_academico string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("correo_electronico_administrador_academico = ? AND hash_contrasena_administrador_academico = ?", correo_electronico_administrador_academico, hash_contrasena_administrador_academico).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AuthenticateAdministradorTi(u *Models.AdministradorTi, correo_electronico_administrador_ti string, hash_contrasena_administrador_ti string) (err error) {
	if err := Config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload(clause.Associations).Where("correo_electronico_administrador_ti = ? AND hash_contrasena_administrador_ti = ?", correo_electronico_administrador_ti, hash_contrasena_administrador_ti).First(u).Error; err != nil {
		return err
	}
	return nil
}

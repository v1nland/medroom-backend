package Controllers

import (
	"errors"
	"medroom-backend/ApiHelpers"
	"medroom-backend/Formats/Input"
	"medroom-backend/Formats/Output"
	"medroom-backend/Messages/Request"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Lista de administradores-academicos
// @Description Lista todos los administradores-academicos
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListAdministradoresAcademicosSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos [get]
func ListAdministradoresAcademicos(c *gin.Context) {
	var container []Models.AdministradorAcademico

	if err := Repositories.GetAllAdministradoresAcademicos(&container); err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			ApiHelpers.RespondJSON(c, 200, container)
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.ListAdministradoresAcademicosOutput(container))
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Obtiene un administrador_academico
// @Description Obtiene un administrador_academico según su UUID
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a buscar"
// @Success 200 {object} Swagger.GetOneAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos/{uuid_administrador_academico} [get]
func GetOneAdministradorAcademico(c *gin.Context) {
	id := c.Params.ByName("id")

	var container Models.AdministradorAcademico
	if err := Repositories.GetOneAdministradorAcademico(&container, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}
	}

	// ApiHelpers.RespondJSON(c, 200, Output.GetOneAdministradorAcademicoOutput(container))
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Agrega un nuevo administrador_academico
// @Description Genera un nuevo administrador_academico con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_administrador_academico     body    Request.AddNewAdministradorAcademicoPayload     true        "AdministradorAcademico a agregar"
// @Success 200 {object} Swagger.AddNewAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos [post]
func AddNewAdministradorAcademico(c *gin.Context) {
	var input Request.AddNewAdministradorAcademicoPayload

	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format
	Input.AddNewAdministradorAcademicoInput(&input)

	// generate model entity
	model_container := Models.AdministradorAcademico{
		Id_rol:                                     *input.Id_rol,
		Rut_administrador_academico:                *input.Rut_administrador_academico,
		Nombres_administrador_academico:            *input.Nombres_administrador_academico,
		Apellidos_administrador_academico:          *input.Apellidos_administrador_academico,
		Hash_contrasena_administrador_academico:    *input.Hash_contrasena_administrador_academico,
		Correo_electronico_administrador_academico: *input.Correo_electronico_administrador_academico,
		Telefono_fijo_administrador_academico:      *input.Telefono_fijo_administrador_academico,
		Telefono_celular_administrador_academico:   *input.Telefono_celular_administrador_academico,
	}

	// query
	err := Repositories.AddNewAdministradorAcademico(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.AddNewAdministradorAcademicoOutput(model_container))
}

// @Summary Modifica un administrador_academico
// @Description Modifica un administrador_academico con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a modificar"
// @Param   input_actualiza_administrador_academico     body    Request.PutOneAdministradorAcademicoPayload     true        "AdministradorAcademico a modificar"
// @Success 200 {object} Swagger.PutOneAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos/{uuid_administrador_academico} [put]
func PutOneAdministradorAcademico(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container Request.PutOneAdministradorAcademicoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutOneAdministradorAcademicoInput(&container)

	// generate model entity
	var model Models.AdministradorAcademico
	if err := Repositories.GetOneAdministradorAcademico(&model, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model = Models.AdministradorAcademico{
		Id:                                      model.Id,
		Id_rol:                                  model.Id_rol,
		Rut_administrador_academico:             Utils.CheckNullString(container.Rut_administrador_academico, model.Rut_administrador_academico),
		Nombres_administrador_academico:         Utils.CheckNullString(container.Nombres_administrador_academico, model.Nombres_administrador_academico),
		Apellidos_administrador_academico:       Utils.CheckNullString(container.Apellidos_administrador_academico, model.Apellidos_administrador_academico),
		Hash_contrasena_administrador_academico: Utils.CheckNullString(container.Hash_contrasena_administrador_academico, model.Hash_contrasena_administrador_academico),
		Correo_electronico_administrador_academico: Utils.CheckNullString(container.Correo_electronico_administrador_academico, model.Correo_electronico_administrador_academico),
		Telefono_fijo_administrador_academico:      Utils.CheckNullString(container.Telefono_fijo_administrador_academico, model.Telefono_fijo_administrador_academico),
		Telefono_celular_administrador_academico:   Utils.CheckNullString(container.Telefono_celular_administrador_academico, model.Telefono_celular_administrador_academico),
	}

	if err := Repositories.PutOneAdministradorAcademico(&model, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.PutOneAdministradorAcademicoOutput(model))
}

// @Summary Elimina un administrador_academico
// @Description Elimina un administrador_academico con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a eliminar"
// @Success 200 {object} Swagger.DeleteAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos/{uuid_administrador_academico} [delete]
func DeleteAdministradorAcademico(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.AdministradorAcademico

	// get query
	err := Repositories.GetOneAdministradorAcademico(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = Repositories.DeleteAdministradorAcademico(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.DeleteAdministradorAcademicoOutput(container))
}

// @Summary Obtiene el perfil del administrador academico
// @Description Obtiene el perfil del administrador academico según su token
// @Tags 04 - Administración Academica
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetMyAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/me [get]
func GetMyAdministradorAcademico(c *gin.Context) {
	// params
	id_administrador_academico := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	// model container
	var container Models.AdministradorAcademico

	// query
	err := Repositories.GetOneAdministradorAcademico(&container, id_administrador_academico)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetMyAdministradorAcademicoOutput(container))
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil de un administrador_academico con los datos entregados
// @Tags 04 - Administración Academica
// @Accept  json
// @Produce  json
// @Param   input_actualiza_administrador_academico     body    Request.PutMyAdministradorAcademicoPayload     true        "AdministradorAcademico a modificar"
// @Success 200 {object} Swagger.PutMyAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/me [put]
func PutMyAdministradorAcademico(c *gin.Context) {
	// params
	id_administrador_academico := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	// input input
	var input Request.PutMyAdministradorAcademicoPayload
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutMyAdministradorAcademicoInput(&input)

	// generate model entity
	var model Models.AdministradorAcademico
	if err := Repositories.GetOneAdministradorAcademico(&model, id_administrador_academico); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Administrador académico not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// replace data in model entity
	model = Models.AdministradorAcademico{
		Id:                                      model.Id,
		Id_rol:                                  model.Id_rol,
		Rut_administrador_academico:             model.Rut_administrador_academico,
		Nombres_administrador_academico:         model.Nombres_administrador_academico,
		Apellidos_administrador_academico:       model.Apellidos_administrador_academico,
		Hash_contrasena_administrador_academico: Utils.CheckNullString(input.Hash_contrasena_administrador_academico, model.Hash_contrasena_administrador_academico),
		Correo_electronico_administrador_academico: Utils.CheckNullString(input.Correo_electronico_administrador_academico, model.Correo_electronico_administrador_academico),
		Telefono_fijo_administrador_academico:      Utils.CheckNullString(input.Telefono_fijo_administrador_academico, model.Telefono_fijo_administrador_academico),
		Telefono_celular_administrador_academico:   Utils.CheckNullString(input.Telefono_celular_administrador_academico, model.Telefono_celular_administrador_academico),
	}

	// put query
	if err := Repositories.PutOneAdministradorAcademico(&model, id_administrador_academico); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Administrador académico not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.PutMyAdministradorAcademicoOutput(model))
}

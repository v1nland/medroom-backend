package Controllers

import (
	"errors"
	"medroom-backend/ApiHelpers"
	"medroom-backend/Formats/Input"
	"medroom-backend/Messages/Request"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Lista de evaluadores
// @Description Lista todos los evaluadores
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListEvaluadoresSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores [get]
func ListEvaluadores(c *gin.Context) {
	// model container
	var container []Models.Evaluador
	if err := Repositories.GetAllEvaluadores(&container); err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			ApiHelpers.RespondJSON(c, 200, container)
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.ListEvaluadoresOutput(container))
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Obtiene un evaluador
// @Description Obtiene un evaluador según su UUID
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a buscar"
// @Success 200 {object} Swagger.GetOneEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores/{uuid_evaluador} [get]
func GetOneEvaluador(c *gin.Context) {
	id := c.Params.ByName("id")

	// model container
	var container Models.Evaluador
	if err := Repositories.GetOneEvaluador(&container, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.GetOneEvaluadorOutput(container))
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Agrega un nuevo evaluador
// @Description Genera un nuevo evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_evaluador     body    Request.AddNewEvaluadorPayload     true        "Evaluador a agregar"
// @Success 200 {object} Swagger.AddNewEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores [post]
func AddNewEvaluador(c *gin.Context) {
	var input Request.AddNewEvaluadorPayload
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.AddNewEvaluadorInput(&input)

	model := Models.Evaluador{
		Id_rol:                       *input.Id_rol,
		Rut_evaluador:                *input.Rut_evaluador,
		Nombres_evaluador:            *input.Nombres_evaluador,
		Apellidos_evaluador:          *input.Apellidos_evaluador,
		Hash_contrasena_evaluador:    *input.Hash_contrasena_evaluador,
		Correo_electronico_evaluador: *input.Correo_electronico_evaluador,
		Telefono_fijo_evaluador:      *input.Telefono_fijo_evaluador,
		Telefono_celular_evaluador:   *input.Telefono_celular_evaluador,
		Recinto_evaluador:            *input.Recinto_evaluador,
		Cargo_evaluador:              *input.Cargo_evaluador,
	}

	err := Repositories.AddNewEvaluador(&model)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.AddNewEvaluadorOutput(model))
	ApiHelpers.RespondJSON(c, 200, model)
}

// @Summary Modifica un evaluador
// @Description Modifica un evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a modificar"
// @Param   input_actualiza_evaluador     body    Request.PutOneEvaluadorPayload     true        "Evaluador a modificar"
// @Success 200 {object} Swagger.PutOneEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores/{uuid_evaluador} [put]
func PutOneEvaluador(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input bind
	var input Request.PutOneEvaluadorPayload
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutOneEvaluadorInput(&input)

	// generate model entity
	var model Models.Evaluador
	if err := Repositories.GetOneEvaluador(&model, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Evaluador not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// replace data in model entity
	model = Models.Evaluador{
		Id:                           model.Id,
		Id_rol:                       model.Id_rol,
		Rol_evaluador:                model.Rol_evaluador,
		Rut_evaluador:                Utils.CheckNullString(input.Rut_evaluador, model.Rut_evaluador),
		Nombres_evaluador:            Utils.CheckNullString(input.Nombres_evaluador, model.Nombres_evaluador),
		Apellidos_evaluador:          Utils.CheckNullString(input.Apellidos_evaluador, model.Apellidos_evaluador),
		Hash_contrasena_evaluador:    Utils.CheckNullString(input.Hash_contrasena_evaluador, model.Hash_contrasena_evaluador),
		Correo_electronico_evaluador: Utils.CheckNullString(input.Correo_electronico_evaluador, model.Correo_electronico_evaluador),
		Telefono_fijo_evaluador:      Utils.CheckNullString(input.Telefono_fijo_evaluador, model.Telefono_fijo_evaluador),
		Telefono_celular_evaluador:   Utils.CheckNullString(input.Telefono_celular_evaluador, model.Telefono_celular_evaluador),
		Recinto_evaluador:            Utils.CheckNullString(input.Recinto_evaluador, model.Recinto_evaluador),
		Cargo_evaluador:              Utils.CheckNullString(input.Cargo_evaluador, model.Cargo_evaluador),
	}

	// put query
	if err := Repositories.PutOneEvaluador(&model, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.PutOneEvaluadorOutput(model))
	ApiHelpers.RespondJSON(c, 200, model)
}

// @Summary Elimina un evaluador
// @Description Elimina un evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a eliminar"
// @Success 200 {object} Swagger.DeleteEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores/{uuid_evaluador} [delete]
func DeleteEvaluador(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Evaluador
	if err := Repositories.GetOneEvaluador(&container, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	if err := Repositories.DeleteEvaluador(&container, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.DeleteEvaluadorOutput(container))
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Obtiene el perfil del evaluador
// @Description Obtiene el perfil del evaluador según su token
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetMyEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me [get]
func GetMyEvaluador(c *gin.Context) {
	// params
	id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	// model container
	var container Models.Evaluador
	if err := Repositories.GetOneEvaluador(&container, id_evaluador); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.GetMyEvaluadorOutput(container))
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil del propio evaluador con los datos entregados
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   input_actualiza_evaluador     body    Request.PutMyEvaluadorPayload     true        "Evaluador a modificar"
// @Success 200 {object} Swagger.PutMyEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me [put]
func PutMyEvaluador(c *gin.Context) {
	// params
	id := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	// input bind
	var container Request.PutMyEvaluadorPayload
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutMyEvaluadorInput(&container)

	// get model entity
	var model Models.Evaluador
	if err := Repositories.GetOneEvaluador(&model, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// replace data in model entity
	model = Models.Evaluador{
		Id:                           model.Id,
		Id_rol:                       model.Id_rol,
		Rol_evaluador:                model.Rol_evaluador,
		Rut_evaluador:                model.Rut_evaluador,
		Nombres_evaluador:            model.Nombres_evaluador,
		Apellidos_evaluador:          model.Apellidos_evaluador,
		Hash_contrasena_evaluador:    Utils.CheckNullString(container.Hash_contrasena_evaluador, model.Hash_contrasena_evaluador),
		Correo_electronico_evaluador: model.Correo_electronico_evaluador,
		Telefono_fijo_evaluador:      Utils.CheckNullString(container.Telefono_fijo_evaluador, model.Telefono_fijo_evaluador),
		Telefono_celular_evaluador:   Utils.CheckNullString(container.Telefono_celular_evaluador, model.Telefono_celular_evaluador),
		Recinto_evaluador:            model.Recinto_evaluador,
		Cargo_evaluador:              Utils.CheckNullString(container.Cargo_evaluador, model.Cargo_evaluador),
	}

	// put query
	if err := Repositories.PutOneEvaluador(&model, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.PutMyEvaluadorOutput(model))
	ApiHelpers.RespondJSON(c, 200, model)
}

package Controllers

import (
	"medroom-backend/ApiHelpers"
	"medroom-backend/Formats/Input"
	"medroom-backend/Formats/Output"
	"medroom-backend/Messages/Request"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de evaluadores
// @Description Lista todos los evaluadores
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} SwaggerMessages.ListEvaluadoresSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores [get]
func ListEvaluadores(c *gin.Context) {
	// model container
	var container []Models.Evaluador

	// query
	err := Repositories.GetAllEvaluadores(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetEvaluadoresOutput(container))
}

// @Summary Obtiene un evaluador
// @Description Obtiene un evaluador según su UUID
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a buscar"
// @Success 200 {object} SwaggerMessages.GetOneEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores/{uuid_evaluador} [get]
func GetOneEvaluador(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Evaluador

	// query
	err := Repositories.GetOneEvaluador(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetOneEvaluadorOutput(container))
}

// @Summary Agrega un nuevo evaluador
// @Description Genera un nuevo evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_evaluador     body    Request.AddNewEvaluadorPayload     true        "Evaluador a agregar"
// @Success 200 {object} SwaggerMessages.AddNewEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores [post]
func AddNewEvaluador(c *gin.Context) {
	// input container
	var container Request.AddNewEvaluadorPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.AddNewEvaluadorInput(&container)

	// generate model entity
	model_container := Models.Evaluador{
		Id_rol:                       container.Id_rol,
		Rut_evaluador:                container.Rut_evaluador,
		Nombres_evaluador:            container.Nombres_evaluador,
		Apellidos_evaluador:          container.Apellidos_evaluador,
		Hash_contrasena_evaluador:    container.Hash_contrasena_evaluador,
		Correo_electronico_evaluador: container.Correo_electronico_evaluador,
		Telefono_fijo_evaluador:      container.Telefono_fijo_evaluador,
		Telefono_celular_evaluador:   container.Telefono_celular_evaluador,
		Recinto_evaluador:            container.Recinto_evaluador,
		Cargo_evaluador:              container.Cargo_evaluador,
	}

	// query
	err := Repositories.AddNewEvaluador(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.AddNewEvaluadorOutput(model_container))
}

// @Summary Modifica un evaluador
// @Description Modifica un evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a modificar"
// @Param   input_actualiza_evaluador     body    Request.PutOneEvaluadorPayload     true        "Evaluador a modificar"
// @Success 200 {object} SwaggerMessages.PutOneEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores/{uuid_evaluador} [put]
func PutOneEvaluador(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container Request.PutOneEvaluadorPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutOneEvaluadorInput(&container)

	// generate model entity
	var model_container Models.Evaluador

	// get query
	err := Repositories.GetOneEvaluador(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.Evaluador{
		Id:                           model_container.Id,
		Id_rol:                       Utils.CheckUpdatedInt(container.Id_rol, model_container.Id_rol),
		Rut_evaluador:                Utils.CheckUpdatedString(container.Rut_evaluador, model_container.Rut_evaluador),
		Nombres_evaluador:            Utils.CheckUpdatedString(container.Nombres_evaluador, model_container.Nombres_evaluador),
		Apellidos_evaluador:          Utils.CheckUpdatedString(container.Apellidos_evaluador, model_container.Apellidos_evaluador),
		Hash_contrasena_evaluador:    Utils.CheckUpdatedString(container.Hash_contrasena_evaluador, model_container.Hash_contrasena_evaluador),
		Correo_electronico_evaluador: Utils.CheckUpdatedString(container.Correo_electronico_evaluador, model_container.Correo_electronico_evaluador),
		Telefono_fijo_evaluador:      Utils.CheckUpdatedString(container.Telefono_fijo_evaluador, model_container.Telefono_fijo_evaluador),
		Telefono_celular_evaluador:   Utils.CheckUpdatedString(container.Telefono_celular_evaluador, model_container.Telefono_celular_evaluador),
		Recinto_evaluador:            Utils.CheckUpdatedString(container.Recinto_evaluador, model_container.Recinto_evaluador),
		Cargo_evaluador:              Utils.CheckUpdatedString(container.Cargo_evaluador, model_container.Cargo_evaluador),
	}

	// update foreign entity
	err = Repositories.GetOneRol(&model_container.Rol_evaluador, Utils.ConvertIntToString(model_container.Id_rol))
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// put query
	err = Repositories.PutOneEvaluador(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.PutOneEvaluadorOutput(model_container))
}

// @Summary Elimina un evaluador
// @Description Elimina un evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a eliminar"
// @Success 200 {object} SwaggerMessages.DeleteEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores/{uuid_evaluador} [delete]
func DeleteEvaluador(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Evaluador

	// get query
	err := Repositories.GetOneEvaluador(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = Repositories.DeleteEvaluador(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.DeleteEvaluadorOutput(container))
}

// @Summary Obtiene el perfil del evaluador
// @Description Obtiene el perfil del evaluador según su token
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Success 200 {object} SwaggerMessages.GetMyEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me [get]
func GetMyEvaluador(c *gin.Context) {
	// params
	id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	// model container
	var container Models.Evaluador
	if err := Repositories.GetOneEvaluador(&container, id_evaluador); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetMyEvaluadorOutput(container))
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil del propio evaluador con los datos entregados
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   input_actualiza_evaluador     body    Request.PutMyEvaluadorPayload     true        "Evaluador a modificar"
// @Success 200 {object} SwaggerMessages.PutMyEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me [put]
func PutMyEvaluador(c *gin.Context) {
	// params
	id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	// input container
	var container Request.PutMyEvaluadorPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutMyEvaluadorInput(&container)

	// generate model entity
	var model_container Models.Evaluador

	// get query
	err := Repositories.GetOneEvaluador(&model_container, id_evaluador)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.Evaluador{
		Id:                           model_container.Id,
		Id_rol:                       Utils.CheckUpdatedInt(container.Id_rol, model_container.Id_rol),
		Rut_evaluador:                Utils.CheckUpdatedString(container.Rut_evaluador, model_container.Rut_evaluador),
		Nombres_evaluador:            Utils.CheckUpdatedString(container.Nombres_evaluador, model_container.Nombres_evaluador),
		Apellidos_evaluador:          Utils.CheckUpdatedString(container.Apellidos_evaluador, model_container.Apellidos_evaluador),
		Hash_contrasena_evaluador:    Utils.CheckUpdatedString(container.Hash_contrasena_evaluador, model_container.Hash_contrasena_evaluador),
		Correo_electronico_evaluador: Utils.CheckUpdatedString(container.Correo_electronico_evaluador, model_container.Correo_electronico_evaluador),
		Telefono_fijo_evaluador:      Utils.CheckUpdatedString(container.Telefono_fijo_evaluador, model_container.Telefono_fijo_evaluador),
		Telefono_celular_evaluador:   Utils.CheckUpdatedString(container.Telefono_celular_evaluador, model_container.Telefono_celular_evaluador),
		Recinto_evaluador:            Utils.CheckUpdatedString(container.Recinto_evaluador, model_container.Recinto_evaluador),
		Cargo_evaluador:              Utils.CheckUpdatedString(container.Cargo_evaluador, model_container.Cargo_evaluador),
	}

	// update foreign entity
	err = Repositories.GetOneRol(&model_container.Rol_evaluador, Utils.ConvertIntToString(model_container.Id_rol))
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// put query
	err = Repositories.PutOneEvaluador(&model_container, id_evaluador)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.PutMyEvaluadorOutput(model_container))
}

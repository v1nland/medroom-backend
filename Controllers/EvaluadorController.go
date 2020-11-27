package Controllers

import (
	"github.com/gin-gonic/gin"
	"medroom-backend/ApiHelpers"
	"medroom-backend/InputFormats"
	"medroom-backend/Models"
	"medroom-backend/OutputFormats"
	"medroom-backend/Repositories"
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"
)

/*
	*
	*  FUNCIÓN ListEvaluador
	*
    *
	*
	*
    *
*/

// @Summary Lista de evaluadores
// @Description Lista todos los evaluadores
// @Tags Evaluadores
// @Accept  json
// @Produce  json
// @Success 200 {array} SwaggerMessages.ListEvaluadoresSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores [get]
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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetEvaluadoresOutput(container))
}

/*
	*
	*  FUNCIÓN GetOneEvaluador
	*
    *
	*
	*
    *
*/

// @Summary Obtiene un evaluador
// @Description Obtiene un evaluador según su UUID
// @Tags Evaluadores
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a buscar"
// @Success 200 {object} SwaggerMessages.GetOneEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/{uuid_evaluador} [get]
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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetOneEvaluadorOutput(container))
}

/*
	*
	*  FUNCIÓN AddNewEvaluador
	*
    *
	*
	*
    *
*/

// @Summary Agrega un nuevo evaluador
// @Description Genera un nuevo evaluador con los datos entregados
// @Tags Evaluadores
// @Accept  json
// @Produce  json
// @Param   input_evaluador     body    RequestMessages.AddNewEvaluadorPayload     true        "Evaluador a agregar"
// @Success 200 {object} SwaggerMessages.AddNewEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores [post]
func AddNewEvaluador(c *gin.Context) {
	// input container
	var container RequestMessages.AddNewEvaluadorPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.AddNewEvaluadorInput(&container)

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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.AddNewEvaluadorOutput(model_container))
}

/*
	*
	*  FUNCIÓN PutOneEvaluador
	*
    *
	*
	*
    *
*/

// @Summary Modifica un evaluador
// @Description Modifica un evaluador con los datos entregados
// @Tags Evaluadores
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a modificar"
// @Param   input_actualiza_evaluador     body    RequestMessages.PutOneEvaluadorPayload     true        "Evaluador a modificar"
// @Success 200 {object} SwaggerMessages.PutOneEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/{uuid_evaluador} [put]
func PutOneEvaluador(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container RequestMessages.PutOneEvaluadorPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutOneEvaluadorInput(&container)

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
		ID:                           model_container.ID,
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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutOneEvaluadorOutput(model_container))
}

/*
	*
	*  FUNCIÓN DeleteEvaluador
	*
    *
	*
	*
    *
*/

// @Summary Elimina un evaluador
// @Description Elimina un evaluador con los datos entregados
// @Tags Evaluadores
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a eliminar"
// @Success 200 {object} SwaggerMessages.DeleteEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/{uuid_evaluador} [delete]
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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.DeleteEvaluadorOutput(container))
}

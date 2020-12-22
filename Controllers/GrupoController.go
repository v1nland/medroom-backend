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

// @Summary Lista de grupos
// @Description Lista todos los grupos
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} SwaggerMessages.ListGruposSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/grupos [get]
func ListGrupos(c *gin.Context) {
	// model container
	var container []Models.Grupo

	// query
	err := Repositories.GetAllGrupos(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetGruposOutput(container))
}

// @Summary Lista de grupos
// @Description Lista todos los grupos
// @Tags 04 - Administración Academica
// @Accept  json
// @Produce  json
// @Success 200 {array} SwaggerMessages.ListGruposSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos [get]
func swaggerDummyListGrupos(c *gin.Context) {
	ListGrupos(c)
}

// @Summary Obtiene un grupo
// @Description Obtiene un grupo según su Id
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_grupo     path    string     true        "Id del grupo a buscar"
// @Success 200 {object} SwaggerMessages.GetOneGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/grupos/{id_grupo} [get]
func GetOneGrupo(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Grupo

	// query
	err := Repositories.GetOneGrupo(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetOneGrupoOutput(container))
}

// @Summary Obtiene un grupo
// @Description Obtiene un grupo según su Id
// @Tags 04 - Administración Academica
// @Accept  json
// @Produce  json
// @Param   id_grupo     path    string     true        "Id del grupo a buscar"
// @Success 200 {object} SwaggerMessages.GetOneGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos/{id_grupo} [get]
func swaggerDummyGetOneGrupo(c *gin.Context) {
	GetOneGrupo(c)
}

// @Summary Agrega un nuevo grupo
// @Description Genera un nuevo grupo con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_grupo     body    Request.AddNewGrupoPayload     true        "Grupo a agregar"
// @Success 200 {object} SwaggerMessages.AddNewGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/grupos [post]
func AddNewGrupo(c *gin.Context) {
	// input container
	var container Request.AddNewGrupoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.AddNewGrupoInput(&container)

	// generate model entity
	model_container := Models.Grupo{
		Id_curso:     container.Id_curso,
		Id_evaluador: container.Id_evaluador,
		Nombre_grupo: container.Nombre_grupo,
		Sigla_grupo:  container.Sigla_grupo,
	}

	// query
	err := Repositories.AddNewGrupo(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.AddNewGrupoOutput(model_container))
}

// @Summary Agrega un nuevo grupo
// @Description Genera un nuevo grupo con los datos entregados
// @Tags 04 - Administración Academica
// @Accept  json
// @Produce  json
// @Param   input_grupo     body    Request.AddNewGrupoPayload     true        "Grupo a agregar"
// @Success 200 {object} SwaggerMessages.AddNewGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos [post]
func swaggerDummyAddNewGrupo(c *gin.Context) {
	AddNewGrupo(c)
}

// @Summary Modifica un grupo
// @Description Modifica un grupo con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_grupo     path    string     true        "Id del grupo a modificar"
// @Param   input_actualiza_grupo     body    Request.PutOneGrupoPayload     true        "Grupo a modificar"
// @Success 200 {object} SwaggerMessages.PutOneGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/grupos/{id_grupo} [put]
func PutOneGrupo(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container Request.PutOneGrupoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutOneGrupoInput(&container)

	// generate model entity
	var model_container Models.Grupo

	// get query
	err := Repositories.GetOneGrupo(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.Grupo{
		Id:           model_container.Id,
		Id_curso:     Utils.CheckUpdatedInt(container.Id_curso, model_container.Id_curso),
		Id_evaluador: Utils.CheckUpdatedString(container.Id_evaluador, model_container.Id_evaluador),
		Nombre_grupo: Utils.CheckUpdatedString(container.Nombre_grupo, model_container.Nombre_grupo),
		Sigla_grupo:  Utils.CheckUpdatedString(container.Sigla_grupo, model_container.Sigla_grupo),
	}

	// update foreign entity
	err = Repositories.GetOneEvaluador(&model_container.Evaluador_grupo, model_container.Id_evaluador)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// put query
	err = Repositories.PutOneGrupo(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.PutOneGrupoOutput(model_container))
}

// @Summary Modifica un grupo
// @Description Modifica un grupo con los datos entregados
// @Tags 04 - Administración Academica
// @Accept  json
// @Produce  json
// @Param   id_grupo     path    string     true        "Id del grupo a modificar"
// @Param   input_actualiza_grupo     body    Request.PutOneGrupoPayload     true        "Grupo a modificar"
// @Success 200 {object} SwaggerMessages.PutOneGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos/{id_grupo} [put]
func swaggerDummyPutOneGrupo(c *gin.Context) {
	PutOneGrupo(c)
}

// @Summary Elimina un grupo
// @Description Elimina un grupo con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_grupo     path    string     true        "Id del grupo a eliminar"
// @Success 200 {object} SwaggerMessages.DeleteGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/grupos/{id_grupo} [delete]
func DeleteGrupo(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Grupo

	// get query
	err := Repositories.GetOneGrupo(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = Repositories.DeleteGrupo(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.DeleteGrupoOutput(container))
}

// @Summary Elimina un grupo
// @Description Elimina un grupo con los datos entregados
// @Tags 04 - Administración Academica
// @Accept  json
// @Produce  json
// @Param   id_grupo     path    string     true        "Id del grupo a eliminar"
// @Success 200 {object} SwaggerMessages.DeleteGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos/{id_grupo} [delete]
func swaggerDummyDeleteGrupo(c *gin.Context) {
	DeleteGrupo(c)
}

// @Summary Obtiene un grupo de un estudiante
// @Description Obtiene un grupo de un estudiante según su token
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Success 200 {object} SwaggerMessages.GetGrupoEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me/grupo [get]
func GetGrupoEstudiante(c *gin.Context) {
	// params
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	// model container
	var estudiante Models.Estudiante
	if err := Repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// // model container
	// var container Models.Grupo
	// if err := Repositories.GetOneGrupo(&container, Utils.ConvertIntToString(estudiante.Id_grupo)); err != nil {
	// 	ApiHelpers.RespondError(c, 500, "default")
	// 	return
	// }

	// output
	// ApiHelpers.RespondJSON(c, 200, Output.GetGrupoEstudianteOutput(container))
}

// @Summary Obtiene un grupo de un evaluador
// @Description Obtiene un grupo de un evaluador según su token
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Success 200 {object} SwaggerMessages.GetGrupoEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me/grupo [get]
func GetGrupoEvaluador(c *gin.Context) {
	// params
	id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	// model container
	var container Models.Grupo
	if err := Repositories.GetOneGrupoByEvaluadorId(&container, id_evaluador); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetGrupoEvaluadorOutput(container))
}

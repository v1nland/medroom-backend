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
	*  FUNCIÓN ListGrupo
	*
    *
	*
	*
    *
*/

// @Summary Lista de grupos
// @Description Lista todos los grupos
// @Accept  json
// @Produce  json
// @Success 200 {array} ResponseMessages.ListGruposResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /grupos [get]
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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetGruposOutput(container))
}

/*
	*
	*  FUNCIÓN GetOneGrupo
	*
    *
	*
	*
    *
*/

// @Summary Obtiene un grupo
// @Description Obtiene un grupo según su UUID
// @Accept  json
// @Produce  json
// @Param   uuid_grupo     path    string     true        "UUID del grupo a buscar"
// @Success 200 {object} ResponseMessages.GetOneGrupoResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /grupos/{uuid_grupo} [get]
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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetOneGrupoOutput(container))
}

/*
	*
	*  FUNCIÓN AddNewGrupo
	*
    *
	*
	*
    *
*/

// @Summary Agrega un nuevo grupo
// @Description Genera un nuevo grupo con los datos entregados
// @Accept  json
// @Produce  json
// @Param   input_grupo     body    RequestMessages.AddNewGrupoPayload     true        "Grupo a agregar"
// @Success 200 {object} ResponseMessages.AddNewGrupoResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /grupos [post]
func AddNewGrupo(c *gin.Context) {
	// input container
	var container RequestMessages.AddNewGrupoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.AddNewGrupoInput(&container)

	// generate model entity
	model_container := Models.Grupo{
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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.AddNewGrupoOutput(model_container))
}

/*
	*
	*  FUNCIÓN PutOneGrupo
	*
    *
	*
	*
    *
*/

// @Summary Modifica un grupo
// @Description Modifica un grupo con los datos entregados
// @Accept  json
// @Produce  json
// @Param   uuid_grupo     path    string     true        "UUID del grupo a modificar"
// @Param   input_actualiza_grupo     body    RequestMessages.PutOneGrupoPayload     true        "Grupo a modificar"
// @Success 200 {object} ResponseMessages.PutOneGrupoResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /grupos/{uuid_grupo} [put]
func PutOneGrupo(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container RequestMessages.PutOneGrupoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutOneGrupoInput(&container)

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
		ID:           model_container.ID,
		Nombre_grupo: Utils.CheckUpdatedString(container.Nombre_grupo, model_container.Nombre_grupo),
		Sigla_grupo:  Utils.CheckUpdatedString(container.Sigla_grupo, model_container.Sigla_grupo),
	}

	// put query
	err = Repositories.PutOneGrupo(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutOneGrupoOutput(model_container))
}

/*
	*
	*  FUNCIÓN DeleteGrupo
	*
    *
	*
	*
    *
*/

// @Summary Elimina un grupo
// @Description Elimina un grupo con los datos entregados
// @Accept  json
// @Produce  json
// @Param   uuid_grupo     path    string     true        "UUID del grupo a eliminar"
// @Success 200 {object} ResponseMessages.DeleteGrupoResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /grupos/{uuid_grupo} [delete]
func DeleteGrupo(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Grupo

	// query
	err := Repositories.DeleteGrupo(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.DeleteGrupoOutput(container))
}

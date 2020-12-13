package Controllers

import (
	"medroom-backend/ApiHelpers"
	"medroom-backend/InputFormats"
	"medroom-backend/Models"
	"medroom-backend/OutputFormats"
	"medroom-backend/Repositories"
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
)

/*
	*
	*  FUNCIÓN ListCompetencia
	*
    *
	*
	*
    *
*/

// @Summary Lista de competencias
// @Description Lista todos los competencias
// @Tags Competencias
// @Accept  json
// @Produce  json
// @Success 200 {array} SwaggerMessages.ListCompetenciasSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /competencias [get]
func ListCompetencias(c *gin.Context) {
	// model container
	var container []Models.Competencia

	// query
	err := Repositories.GetAllCompetencias(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetCompetenciasOutput(container))
}

/*
	*
	*  FUNCIÓN GetOneCompetencia
	*
    *
	*
	*
    *
*/

// @Summary Obtiene un competencia
// @Description Obtiene un competencia según su ID
// @Tags Competencias
// @Accept  json
// @Produce  json
// @Param   id_competencia     path    string     true        "ID del competencia a buscar"
// @Success 200 {object} SwaggerMessages.GetOneCompetenciaSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /competencias/{id_competencia} [get]
func GetOneCompetencia(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Competencia

	// query
	err := Repositories.GetOneCompetencia(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetOneCompetenciaOutput(container))
}

/*
	*
	*  FUNCIÓN AddNewCompetencia
	*
    *
	*
	*
    *
*/

// @Summary Agrega un nuevo competencia
// @Description Genera un nuevo competencia con los datos entregados
// @Tags Competencias
// @Accept  json
// @Produce  json
// @Param   input_competencia     body    RequestMessages.AddNewCompetenciaPayload     true        "Competencia a agregar"
// @Success 200 {object} SwaggerMessages.AddNewCompetenciaSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /competencias [post]
func AddNewCompetencia(c *gin.Context) {
	// input container
	var container RequestMessages.AddNewCompetenciaPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.AddNewCompetenciaInput(&container)

	// generate model entity
	model_container := Models.Competencia{
		Nombre_competencia: container.Nombre_competencia,
	}

	// query
	err := Repositories.AddNewCompetencia(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.AddNewCompetenciaOutput(model_container))
}

/*
	*
	*  FUNCIÓN PutOneCompetencia
	*
    *
	*
	*
    *
*/

// @Summary Modifica un competencia
// @Description Modifica un competencia con los datos entregados
// @Tags Competencias
// @Accept  json
// @Produce  json
// @Param   id_competencia     path    string     true        "ID del competencia a modificar"
// @Param   input_actualiza_competencia     body    RequestMessages.PutOneCompetenciaPayload     true        "Competencia a modificar"
// @Success 200 {object} SwaggerMessages.PutOneCompetenciaSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /competencias/{id_competencia} [put]
func PutOneCompetencia(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container RequestMessages.PutOneCompetenciaPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutOneCompetenciaInput(&container)

	// generate model entity
	var model_container Models.Competencia

	// get query
	err := Repositories.GetOneCompetencia(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.Competencia{
		ID:                 model_container.ID,
		Nombre_competencia: Utils.CheckUpdatedString(container.Nombre_competencia, model_container.Nombre_competencia),
	}

	// put query
	err = Repositories.PutOneCompetencia(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutOneCompetenciaOutput(model_container))
}

/*
	*
	*  FUNCIÓN DeleteCompetencia
	*
    *
	*
	*
    *
*/

// @Summary Elimina un competencia
// @Description Elimina un competencia con los datos entregados
// @Tags Competencias
// @Accept  json
// @Produce  json
// @Param   id_competencia     path    string     true        "ID del competencia a eliminar"
// @Success 200 {object} SwaggerMessages.DeleteCompetenciaSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /competencias/{id_competencia} [delete]
func DeleteCompetencia(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Competencia

	// get query
	err := Repositories.GetOneCompetencia(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = Repositories.DeleteCompetencia(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.DeleteCompetenciaOutput(container))
}

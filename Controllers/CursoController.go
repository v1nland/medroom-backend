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
	*  FUNCIÓN ListCurso
	*
    *
	*
	*
    *
*/

// @Summary Lista de cursos
// @Description Lista todos los cursos
// @Tags Cursos
// @Accept  json
// @Produce  json
// @Success 200 {array} SwaggerMessages.ListCursosSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /cursos [get]
func ListCursos(c *gin.Context) {
	// model container
	var container []Models.Curso

	// query
	err := Repositories.GetAllCursos(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetCursosOutput(container))
}

/*
	*
	*  FUNCIÓN GetOneCurso
	*
    *
	*
	*
    *
*/

// @Summary Obtiene un curso
// @Description Obtiene un curso según su ID
// @Tags Cursos
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "ID del curso a buscar"
// @Success 200 {object} SwaggerMessages.GetOneCursoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /cursos/{id_curso} [get]
func GetOneCurso(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Curso

	// query
	err := Repositories.GetOneCurso(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetOneCursoOutput(container))
}

/*
	*
	*  FUNCIÓN AddNewCurso
	*
    *
	*
	*
    *
*/

// @Summary Agrega un nuevo curso
// @Description Genera un nuevo curso con los datos entregados
// @Tags Cursos
// @Accept  json
// @Produce  json
// @Param   input_curso     body    RequestMessages.AddNewCursoPayload     true        "Curso a agregar"
// @Success 200 {object} SwaggerMessages.AddNewCursoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /cursos [post]
func AddNewCurso(c *gin.Context) {
	// input container
	var container RequestMessages.AddNewCursoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.AddNewCursoInput(&container)

	// generate model entity
	model_container := Models.Curso{
		Id_periodo:   container.Id_periodo,
		Nombre_curso: container.Nombre_curso,
		Sigla_curso:  container.Sigla_curso,
	}

	// query
	err := Repositories.AddNewCurso(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.AddNewCursoOutput(model_container))
}

/*
	*
	*  FUNCIÓN PutOneCurso
	*
    *
	*
	*
    *
*/

// @Summary Modifica un curso
// @Description Modifica un curso con los datos entregados
// @Tags Cursos
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "ID del curso a modificar"
// @Param   input_actualiza_curso     body    RequestMessages.PutOneCursoPayload     true        "Curso a modificar"
// @Success 200 {object} SwaggerMessages.PutOneCursoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /cursos/{id_curso} [put]
func PutOneCurso(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container RequestMessages.PutOneCursoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutOneCursoInput(&container)

	// generate model entity
	var model_container Models.Curso

	// get query
	err := Repositories.GetOneCurso(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.Curso{
		ID:           model_container.ID,
		Id_periodo:   Utils.CheckUpdatedInt(container.Id_periodo, model_container.Id_periodo),
		Nombre_curso: Utils.CheckUpdatedString(container.Nombre_curso, model_container.Nombre_curso),
		Sigla_curso:  Utils.CheckUpdatedString(container.Sigla_curso, model_container.Sigla_curso),
	}

	// update foreign entity
	err = Repositories.GetOnePeriodo(&model_container.Periodo_curso, Utils.ConvertIntToString(model_container.Id_periodo))
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// put query
	err = Repositories.PutOneCurso(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutOneCursoOutput(model_container))
}

/*
	*
	*  FUNCIÓN DeleteCurso
	*
    *
	*
	*
    *
*/

// @Summary Elimina un curso
// @Description Elimina un curso con los datos entregados
// @Tags Cursos
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "ID del curso a eliminar"
// @Success 200 {object} SwaggerMessages.DeleteCursoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /cursos/{id_curso} [delete]
func DeleteCurso(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Curso

	// get query
	err := Repositories.GetOneCurso(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = Repositories.DeleteCurso(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.DeleteCursoOutput(container))
}

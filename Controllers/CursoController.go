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

// @Summary Lista de cursos
// @Description Lista todos los cursos
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListCursosSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos [get]
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
	ApiHelpers.RespondJSON(c, 200, Output.ListCursosOutput(container))
}

// @Summary Obtiene un curso
// @Description Obtiene un curso según su Id
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a buscar"
// @Success 200 {object} Swagger.GetOneCursoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos/{id_curso} [get]
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
	ApiHelpers.RespondJSON(c, 200, Output.GetOneCursoOutput(container))
}

// @Summary Agrega un nuevo curso
// @Description Genera un nuevo curso con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_curso     body    Request.AddNewCursoPayload     true        "Curso a agregar"
// @Success 200 {object} Swagger.AddNewCursoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos [post]
func AddNewCurso(c *gin.Context) {
	// input container
	var container Request.AddNewCursoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.AddNewCursoInput(&container)

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
	ApiHelpers.RespondJSON(c, 200, Output.AddNewCursoOutput(model_container))
}

// @Summary Modifica un curso
// @Description Modifica un curso con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a modificar"
// @Param   input_actualiza_curso     body    Request.PutOneCursoPayload     true        "Curso a modificar"
// @Success 200 {object} Swagger.PutOneCursoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos/{id_curso} [put]
func PutOneCurso(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container Request.PutOneCursoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutOneCursoInput(&container)

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
		Id:           model_container.Id,
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
	ApiHelpers.RespondJSON(c, 200, Output.PutOneCursoOutput(model_container))
}

// @Summary Elimina un curso
// @Description Elimina un curso con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a eliminar"
// @Success 200 {object} Swagger.DeleteCursoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos/{id_curso} [delete]
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
	ApiHelpers.RespondJSON(c, 200, Output.DeleteCursoOutput(container))
}

// @Summary Obtiene un curso de un estudiante
// @Description Obtiene un curso de un estudiante según su token
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetCursoEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me/curso [get]
func GetCursoEstudiante(c *gin.Context) {
	// params
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	// model container
	var estudiante Models.Estudiante
	if err := Repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// model container
	// var container Models.Grupo
	// if err := Repositories.GetOneGrupo(&container, Utils.ConvertIntToString(estudiante.Id_grupo)); err != nil {
	// 	ApiHelpers.RespondError(c, 500, "default")
	// 	return
	// }

	// // model container
	// var container_curso Models.Curso
	// if err := Repositories.GetOneCurso(&container_curso, Utils.ConvertIntToString(estudiante.Id_grupo)); err != nil {
	// 	ApiHelpers.RespondError(c, 500, "default")
	// 	return
	// }
	// output
	// ApiHelpers.RespondJSON(c, 200, Output.GetCursoEstudianteOutput(container_curso))
}

// @Summary Obtiene un curso de un evaluador
// @Description Obtiene un curso de un evaluador según su token
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetCursoEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me/curso [get]
func GetCursoEvaluador(c *gin.Context) {
	// params
	id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	// model container
	var container Models.Grupo
	if err := Repositories.GetOneGrupoByEvaluadorId(&container, id_evaluador); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// model container
	var curso_container Models.Curso
	if err := Repositories.GetOneCurso(&curso_container, Utils.ConvertIntToString(container.Id_curso)); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetCursoEvaluadorOutput(curso_container))
}

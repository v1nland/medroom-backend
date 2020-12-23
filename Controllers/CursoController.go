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
	if err := Repositories.GetAllCursos(&container); err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			ApiHelpers.RespondJSON(c, 200, container)
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// output
	// ApiHelpers.RespondJSON(c, 200, Output.ListCursosOutput(container))
	ApiHelpers.RespondJSON(c, 200, container)
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
	if err := Repositories.GetOneCurso(&container, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Curso not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

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
	var input Request.AddNewCursoPayload
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.AddNewCursoInput(&input)

	model := Models.Curso{
		Id_periodo:   *input.Id_periodo,
		Nombre_curso: *input.Nombre_curso,
		Sigla_curso:  *input.Sigla_curso,
		Estado_curso: true,
	}

	if err := Repositories.AddNewCurso(&model); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	// ApiHelpers.RespondJSON(c, 200, Output.AddNewCursoOutput(model_container))
	ApiHelpers.RespondJSON(c, 200, model)
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

	// input input
	var input Request.PutOneCursoPayload
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutOneCursoInput(&input)

	// generate model entity
	var model Models.Curso
	if err := Repositories.GetOneCurso(&model, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Curso not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// replace data in model entity
	model = Models.Curso{
		Id:           model.Id,
		Id_periodo:   Utils.CheckNullInt(input.Id_periodo, model.Id_periodo),
		Nombre_curso: Utils.CheckNullString(input.Nombre_curso, model.Nombre_curso),
		Sigla_curso:  Utils.CheckNullString(input.Sigla_curso, model.Sigla_curso),
	}

	if err := Repositories.PutOneCurso(&model, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Curso not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.PutOneCursoOutput(model))
	ApiHelpers.RespondJSON(c, 200, model)
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

	var container Models.Curso
	if err := Repositories.GetOneCurso(&container, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Curso not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	if err := Repositories.DeleteCurso(&container, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Curso not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.DeleteCursoOutput(container))
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Obtiene los cursos de un estudiante
// @Description Obtiene los cursos de un estudiante según su token
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetCursosEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me/cursos [get]
func GetCursosEstudiante(c *gin.Context) {
	id := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var estudiante Models.Estudiante
	if err := Repositories.GetOneEstudiante(&estudiante, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	var cursos []Models.Curso
	for i := 0; i < len(estudiante.Grupos_estudiante); i++ {
		var curso Models.Curso
		if err := Repositories.GetOneCurso(&curso, Utils.ConvertIntToString(estudiante.Grupos_estudiante[i].Id_curso)); err != nil {
			ApiHelpers.RespondError(c, 500, "default")
			return
		}

		cursos = append(cursos, curso)
	}

	// ApiHelpers.RespondJSON(c, 200, Output.GetCursoEstudianteOutput(cursos))
	ApiHelpers.RespondJSON(c, 200, cursos)
}

// @Summary Obtiene un curso de un evaluador
// @Description Obtiene un curso de un evaluador según su token
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetCursoEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos [get]
func GetCursosEvaluador(c *gin.Context) {
	// params
	id := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var grupos []Models.Grupo
	if err := Repositories.GetAllGruposByEvaluadorId(&grupos, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	var cursos []Models.Curso
	for i := 0; i < len(grupos); i++ {
		var curso Models.Curso
		if err := Repositories.GetOneCurso(&curso, Utils.ConvertIntToString(grupos[i].Id_curso)); err != nil {
			ApiHelpers.RespondError(c, 500, "default")
			return
		}

		cursos = append(cursos, curso)
	}

	// ApiHelpers.RespondJSON(c, 200, Output.GetCursoEvaluadorOutput(cursos))
	ApiHelpers.RespondJSON(c, 200, cursos)
}

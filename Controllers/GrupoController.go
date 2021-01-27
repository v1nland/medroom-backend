package Controllers

import (
	"errors"
	"medroom-backend/ApiHelpers"
	"medroom-backend/Config"
	"medroom-backend/Formats/Input"
	"medroom-backend/Messages/Request"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Lista de grupos
// @Description Lista todos los grupos
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListGruposSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos [get]
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
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Obtiene un grupo
// @Description Obtiene un grupo según su Id
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_grupo     path    string     true        "Id del grupo a buscar"
// @Success 200 {object} Swagger.GetOneGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos/{id_grupo} [get]
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
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Agrega un nuevo grupo
// @Description Genera un nuevo grupo con los datos entregados
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   input_grupo     body    Request.AddNewGrupoPayload     true        "Grupo a agregar"
// @Success 200 {object} Swagger.AddNewGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos [post]
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
		Id_curso:     *container.Id_curso,
		Nombre_grupo: *container.Nombre_grupo,
		Sigla_grupo:  *container.Sigla_grupo,
	}

	// query
	err := Repositories.AddNewGrupo(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, model_container)
}

// @Summary Modifica un grupo
// @Description Modifica un grupo con los datos entregados
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_grupo     path    string     true        "Id del grupo a modificar"
// @Param   input_actualiza_grupo     body    Request.PutOneGrupoPayload     true        "Grupo a modificar"
// @Success 200 {object} Swagger.PutOneGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos/{id_grupo} [put]
func PutOneGrupo(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input input
	var input Request.PutOneGrupoPayload
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutOneGrupoInput(&input)

	// generate model entity
	var model Models.Grupo
	if err := Repositories.GetOneGrupo(&model, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model = Models.Grupo{
		Id:           model.Id,
		Id_curso:     Utils.CheckNullInt(input.Id_curso, model.Id_curso),
		Nombre_grupo: Utils.CheckNullString(input.Nombre_grupo, model.Nombre_grupo),
		Sigla_grupo:  Utils.CheckNullString(input.Sigla_grupo, model.Sigla_grupo),
	}

	// put query
	if err := Repositories.PutOneGrupo(&model, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, model)
}

// @Summary Elimina un grupo
// @Description Elimina un grupo con los datos entregados
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_grupo     path    string     true        "Id del grupo a eliminar"
// @Success 200 {object} Swagger.DeleteGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos/{id_grupo} [delete]
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
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Obtiene los grupos de un estudiante
// @Description Obtiene los grupos de un estudiante según su token
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a buscar"
// @Success 200 {object} Swagger.GetGruposEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me/cursos/{id_curso}/grupos [get]
func GetGruposEstudiante(c *gin.Context) {
	id_curso := c.Params.ByName("id_curso")
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var grupos []Models.Grupo
	if err := Repositories.GetGruposEstudiante(&grupos, id_curso, id_estudiante); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Grupo not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	ApiHelpers.RespondJSON(c, 200, grupos)
}

// @Summary Obtiene un grupo de un estudiante
// @Description Obtiene un grupo de un estudiante según su token
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a buscar"
// @Param   id_grupo     path    string     true        "Id del grupo a buscar"
// @Success 200 {object} Swagger.GetOneGrupoEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me/cursos/{id_curso}/grupos/{id_grupo} [get]
func GetOneGrupoEstudiante(c *gin.Context) {
	id_grupo := c.Params.ByName("id_grupo")
	id_curso := c.Params.ByName("id_curso")
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var grupo Models.Grupo
	if err := Repositories.GetOneGrupoEstudiante(&grupo, id_grupo, id_curso, id_estudiante); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Grupo not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	ApiHelpers.RespondJSON(c, 200, grupo)
}

// @Summary Obtiene los grupos de un evaluador
// @Description Obtiene los grupos de un evaluador según su token
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a buscar"
// @Success 200 {object} Swagger.GetGruposEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos/{id_curso}/grupos [get]
func GetGruposEvaluador(c *gin.Context) {
	id_curso := c.Params.ByName("id_curso")
	id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var grupos []Models.Grupo
	if err := Repositories.GetGruposEvaluador(&grupos, id_curso, id_evaluador); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Grupo not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	ApiHelpers.RespondJSON(c, 200, grupos)
}

// @Summary Obtiene un grupo de un evaluador
// @Description Obtiene un grupo de un evaluador según su token
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a buscar"
// @Param   id_grupo     path    string     true        "Id del grupo a buscar"
// @Success 200 {object} Swagger.GetOneGrupoEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos/{id_curso}/grupos/{id_grupo} [get]
func GetOneGrupoEvaluador(c *gin.Context) {
	id_grupo := c.Params.ByName("id_grupo")
	id_curso := c.Params.ByName("id_curso")
	id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var grupo Models.Grupo
	if err := Repositories.GetOneGrupoEvaluador(&grupo, id_grupo, id_curso, id_evaluador); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Grupo not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	ApiHelpers.RespondJSON(c, 200, grupo)
}

// @Summary Modifica los grupos de un estudiante
// @Description Modifica los grupos de un estudiante con los datos entregados
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "ID del curso a modificar"
// @Param   id_grupo     path    string     true        "ID del grupo a modificar"
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a asociar"
// @Success 200 {object} Swagger.AddEstudianteToGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/cursos/{id_curso}/grupos/{id_grupo}/estudiantes/{uuid_estudiante} [put]
func AddEstudianteToGrupo(c *gin.Context) {
	// params
	id_curso := c.Params.ByName("id")
	id_grupo := c.Params.ByName("id_grupo")
	id_estudiante := c.Params.ByName("id_estudiante")

	var curso Models.Curso
	if err := Repositories.GetOneCurso(&curso, id_curso); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Curso not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	var grupo Models.Grupo
	if err := Repositories.GetOneGrupo(&grupo, id_grupo); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Grupo not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	var estudiante Models.Estudiante
	if err := Repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	Config.DB.Model(&grupo).Association("Estudiantes_grupo").Append([]Models.Estudiante{estudiante})

	found, id_grupo_sg := Utils.SearchIdGrupoBySigla(curso.Grupos_curso, "SG")
	if found {
		Repositories.DeleteEstudianteGrupo(Utils.ConvertIntToString(id_grupo_sg), id_estudiante)
	}

	ApiHelpers.RespondJSON(c, 200, grupo)
}

// @Summary Modifica los grupos de un evaluador
// @Description Modifica los grupos de un evaluador con los datos entregados
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "ID del curso a modificar"
// @Param   id_grupo     path    string     true        "ID del grupo a modificar"
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a asociar"
// @Success 200 {object} Swagger.AddEvaluadorToGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/cursos/{id_curso}/grupos/{id_grupo}/evaluadores/{uuid_evaluador} [put]
func AddEvaluadorToGrupo(c *gin.Context) {
	// params
	id_curso := c.Params.ByName("id")
	id_grupo := c.Params.ByName("id_grupo")
	id_evaluador := c.Params.ByName("id_evaluador")

	var curso Models.Curso
	if err := Repositories.GetOneCurso(&curso, id_curso); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Curso not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	var grupo Models.Grupo
	if err := Repositories.GetOneGrupo(&grupo, id_grupo); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Grupo not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	var evaluador Models.Evaluador
	if err := Repositories.GetOneEvaluador(&evaluador, id_evaluador); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Evaluador not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	Config.DB.Model(&grupo).Association("Evaluadores_grupo").Append([]Models.Evaluador{evaluador})

	found, id_grupo_sg := Utils.SearchIdGrupoBySigla(curso.Grupos_curso, "SG")
	if found {
		Repositories.DeleteEvaluadorGrupo(Utils.ConvertIntToString(id_grupo_sg), id_evaluador)
	}

	ApiHelpers.RespondJSON(c, 200, grupo)
}

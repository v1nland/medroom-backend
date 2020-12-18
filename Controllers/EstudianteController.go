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

// @Summary Lista de estudiantes
// @Description Lista todos los estudiantes existentes
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} SwaggerMessages.ListEstudiantesSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes [get]
func ListEstudiantes(c *gin.Context) {
	// model container
	var container []Models.Estudiante

	// query
	err := Repositories.GetAllEstudiantes(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetEstudiantesOutput(container))
}

// @Summary Obtiene un estudiante
// @Description Obtiene un estudiante según su UUID
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a buscar"
// @Success 200 {object} SwaggerMessages.GetOneEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes/{uuid_estudiante} [get]
func GetOneEstudiante(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Estudiante

	// query
	err := Repositories.GetOneEstudiante(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetOneEstudianteOutput(container))
}

// @Summary Agrega un nuevo estudiante
// @Description Genera un nuevo estudiante con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_estudiante     body    RequestMessages.AddNewEstudiantePayload     true        "Estudiante a agregar"
// @Success 200 {object} SwaggerMessages.AddNewEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes [post]
func AddNewEstudiante(c *gin.Context) {
	// input container
	var container RequestMessages.AddNewEstudiantePayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.AddNewEstudianteInput(&container)

	// generate model entity
	model_container := Models.Estudiante{
		Id_rol:                        container.Id_rol,
		Id_grupo:                      container.Id_grupo,
		Rut_estudiante:                container.Rut_estudiante,
		Nombres_estudiante:            container.Nombres_estudiante,
		Apellidos_estudiante:          container.Apellidos_estudiante,
		Hash_contrasena_estudiante:    container.Hash_contrasena_estudiante,
		Correo_electronico_estudiante: container.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      container.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   container.Telefono_celular_estudiante,
	}

	// query
	err := Repositories.AddNewEstudiante(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.AddNewEstudianteOutput(model_container))
}

// @Summary Modifica un estudiante
// @Description Modifica un estudiante con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a modificar"
// @Param   input_actualiza_estudiante     body    RequestMessages.PutOneEstudiantePayload     true        "Estudiante a modificar"
// @Success 200 {object} SwaggerMessages.PutOneEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes/{uuid_estudiante} [put]
func PutOneEstudiante(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container RequestMessages.PutMyEstudiantePayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutMyEstudianteInput(&container)

	// generate model entity
	var model_container Models.Estudiante

	// get query
	if err := Repositories.GetOneEstudiante(&model_container, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.Estudiante{
		ID:                            model_container.ID,
		Id_rol:                        Utils.CheckUpdatedInt(container.Id_rol, model_container.Id_rol),
		Id_grupo:                      Utils.CheckUpdatedInt(container.Id_grupo, model_container.Id_grupo),
		Rut_estudiante:                Utils.CheckUpdatedString(container.Rut_estudiante, model_container.Rut_estudiante),
		Nombres_estudiante:            Utils.CheckUpdatedString(container.Nombres_estudiante, model_container.Nombres_estudiante),
		Apellidos_estudiante:          Utils.CheckUpdatedString(container.Apellidos_estudiante, model_container.Apellidos_estudiante),
		Hash_contrasena_estudiante:    Utils.CheckUpdatedString(container.Hash_contrasena_estudiante, model_container.Hash_contrasena_estudiante),
		Correo_electronico_estudiante: Utils.CheckUpdatedString(container.Correo_electronico_estudiante, model_container.Correo_electronico_estudiante),
		Telefono_fijo_estudiante:      Utils.CheckUpdatedString(container.Telefono_fijo_estudiante, model_container.Telefono_fijo_estudiante),
		Telefono_celular_estudiante:   Utils.CheckUpdatedString(container.Telefono_celular_estudiante, model_container.Telefono_celular_estudiante),
	}

	// update foreign entity
	if err := Repositories.GetOneRol(&model_container.Rol_estudiante, Utils.ConvertIntToString(model_container.Id_rol)); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// put query
	if err := Repositories.PutOneEstudiante(&model_container, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutMyEstudianteOutput(model_container))
}

// @Summary Elimina un estudiante
// @Description Elimina un estudiante con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a eliminar"
// @Success 200 {object} SwaggerMessages.DeleteEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes/{uuid_estudiante} [delete]
func DeleteEstudiante(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Estudiante

	// get query
	err := Repositories.GetOneEstudiante(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = Repositories.DeleteEstudiante(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.DeleteEstudianteOutput(container))
}

// @Summary Obtiene el perfil del estudiante
// @Description Obtiene el perfil del estudiante según su token
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Success 200 {object} SwaggerMessages.GetMyEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me [get]
func GetMyEstudiante(c *gin.Context) {
	// params
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	// model container
	var container Models.Estudiante

	// query
	err := Repositories.GetOneEstudiante(&container, id_estudiante)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetMyEstudianteOutput(container))
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil del propio estudiante con los datos entregados
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   input_actualiza_estudiante     body    RequestMessages.PutMyEstudiantePayload     true        "Nuevos datos del estudiante a modificar"
// @Success 200 {object} SwaggerMessages.PutMyEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me [put]
func PutMyEstudiante(c *gin.Context) {
	// params
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	// input container
	var container RequestMessages.PutMyEstudiantePayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutMyEstudianteInput(&container)

	// generate model entity
	var model_container Models.Estudiante

	// get query
	err := Repositories.GetOneEstudiante(&model_container, id_estudiante)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.Estudiante{
		ID:                            model_container.ID,
		Id_rol:                        Utils.CheckUpdatedInt(container.Id_rol, model_container.Id_rol),
		Id_grupo:                      Utils.CheckUpdatedInt(container.Id_grupo, model_container.Id_grupo),
		Rut_estudiante:                Utils.CheckUpdatedString(container.Rut_estudiante, model_container.Rut_estudiante),
		Nombres_estudiante:            Utils.CheckUpdatedString(container.Nombres_estudiante, model_container.Nombres_estudiante),
		Apellidos_estudiante:          Utils.CheckUpdatedString(container.Apellidos_estudiante, model_container.Apellidos_estudiante),
		Hash_contrasena_estudiante:    Utils.CheckUpdatedString(container.Hash_contrasena_estudiante, model_container.Hash_contrasena_estudiante),
		Correo_electronico_estudiante: Utils.CheckUpdatedString(container.Correo_electronico_estudiante, model_container.Correo_electronico_estudiante),
		Telefono_fijo_estudiante:      Utils.CheckUpdatedString(container.Telefono_fijo_estudiante, model_container.Telefono_fijo_estudiante),
		Telefono_celular_estudiante:   Utils.CheckUpdatedString(container.Telefono_celular_estudiante, model_container.Telefono_celular_estudiante),
	}

	// update foreign entity
	err = Repositories.GetOneRol(&model_container.Rol_estudiante, Utils.ConvertIntToString(model_container.Id_rol))
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// put query
	err = Repositories.PutOneEstudiante(&model_container, id_estudiante)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutOneEstudianteOutput(model_container))
}

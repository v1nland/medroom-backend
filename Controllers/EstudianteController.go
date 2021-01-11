package Controllers

import (
	"errors"
	"medroom-backend/ApiHelpers"
	"medroom-backend/Formats/Input"
	"medroom-backend/Messages/Request"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Lista de estudiantes
// @Description Lista todos los estudiantes existentes
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListEstudiantesSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes [get]
func ListEstudiantes(c *gin.Context) {
	// model container
	var container []Models.Estudiante
	if err := Repositories.GetAllEstudiantes(&container); err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			ApiHelpers.RespondJSON(c, 200, container)
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.ListEstudiantesOutput(container))
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Obtiene un estudiante
// @Description Obtiene un estudiante según su UUID
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a buscar"
// @Success 200 {object} Swagger.GetOneEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes/{uuid_estudiante} [get]
func GetOneEstudiante(c *gin.Context) {
	id := c.Params.ByName("id")

	var container Models.Estudiante
	if err := Repositories.GetOneEstudiante(&container, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.GetOneEstudianteOutput(container))
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Agrega un nuevo estudiante
// @Description Genera un nuevo estudiante con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_estudiante     body    Request.AddNewEstudiantePayload     true        "Estudiante a agregar"
// @Success 200 {object} Swagger.AddNewEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes [post]
func AddNewEstudiante(c *gin.Context) {
	var input Request.AddNewEstudiantePayload
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.AddNewEstudianteInput(&input)

	model := Models.Estudiante{
		Id_rol:                        *input.Id_rol,
		Rut_estudiante:                *input.Rut_estudiante,
		Nombres_estudiante:            *input.Nombres_estudiante,
		Apellidos_estudiante:          *input.Apellidos_estudiante,
		Hash_contrasena_estudiante:    *input.Hash_contrasena_estudiante,
		Correo_electronico_estudiante: *input.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      *input.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   *input.Telefono_celular_estudiante,
	}

	if err := Repositories.AddNewEstudiante(&model); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.AddNewEstudianteOutput(model))
	ApiHelpers.RespondJSON(c, 200, model)
}

// @Summary Modifica un estudiante
// @Description Modifica un estudiante con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a modificar"
// @Param   input_actualiza_estudiante     body    Request.PutOneEstudiantePayload     true        "Estudiante a modificar"
// @Success 200 {object} Swagger.PutOneEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes/{uuid_estudiante} [put]
func PutOneEstudiante(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input bind
	var input Request.PutOneEstudiantePayload
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutOneEstudianteInput(&input)

	// get model entity
	var model Models.Estudiante
	if err := Repositories.GetOneEstudiante(&model, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// replace new data
	model = Models.Estudiante{
		Id:                            model.Id,
		Id_rol:                        model.Id_rol,
		Rol_estudiante:                model.Rol_estudiante,
		Rut_estudiante:                Utils.CheckNullString(input.Rut_estudiante, model.Rut_estudiante),
		Nombres_estudiante:            Utils.CheckNullString(input.Nombres_estudiante, model.Nombres_estudiante),
		Apellidos_estudiante:          Utils.CheckNullString(input.Apellidos_estudiante, model.Apellidos_estudiante),
		Hash_contrasena_estudiante:    Utils.CheckNullString(input.Hash_contrasena_estudiante, model.Hash_contrasena_estudiante),
		Correo_electronico_estudiante: Utils.CheckNullString(input.Correo_electronico_estudiante, model.Correo_electronico_estudiante),
		Telefono_fijo_estudiante:      Utils.CheckNullString(input.Telefono_fijo_estudiante, model.Telefono_fijo_estudiante),
		Telefono_celular_estudiante:   Utils.CheckNullString(input.Telefono_celular_estudiante, model.Telefono_celular_estudiante),
	}

	// put query
	if err := Repositories.PutOneEstudiante(&model, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.PutOneEstudianteOutput(model))
	ApiHelpers.RespondJSON(c, 200, model)
}

// @Summary Elimina un estudiante
// @Description Elimina un estudiante con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a eliminar"
// @Success 200 {object} Swagger.DeleteEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes/{uuid_estudiante} [delete]
func DeleteEstudiante(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	var container Models.Estudiante
	if err := Repositories.GetOneEstudiante(&container, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	if err := Repositories.DeleteEstudiante(&container, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.DeleteEstudianteOutput(container))
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Obtiene el perfil del estudiante
// @Description Obtiene el perfil del estudiante según su token
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetMyEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me [get]
func GetMyEstudiante(c *gin.Context) {
	// params
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	// query estudiante
	var container Models.Estudiante
	if err := Repositories.GetOneEstudiante(&container, id_estudiante); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.GetMyEstudianteOutput(container))
	ApiHelpers.RespondJSON(c, 200, container)
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil del propio estudiante con los datos entregados
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   input_actualiza_estudiante     body    Request.PutMyEstudiantePayload     true        "Nuevos datos del estudiante a modificar"
// @Success 200 {object} Swagger.PutMyEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me [put]
func PutMyEstudiante(c *gin.Context) {
	// params
	id := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	// input bind
	var input Request.PutMyEstudiantePayload
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutMyEstudianteInput(&input)

	// get model entity
	var model Models.Estudiante
	if err := Repositories.GetOneEstudiante(&model, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// replace new data
	model = Models.Estudiante{
		Id:                            model.Id,
		Id_rol:                        model.Id_rol,
		Rol_estudiante:                model.Rol_estudiante,
		Rut_estudiante:                model.Rut_estudiante,
		Nombres_estudiante:            model.Nombres_estudiante,
		Apellidos_estudiante:          model.Apellidos_estudiante,
		Hash_contrasena_estudiante:    Utils.CheckNullString(input.Hash_contrasena_estudiante, model.Hash_contrasena_estudiante),
		Correo_electronico_estudiante: model.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      Utils.CheckNullString(input.Telefono_fijo_estudiante, model.Telefono_fijo_estudiante),
		Telefono_celular_estudiante:   Utils.CheckNullString(input.Telefono_celular_estudiante, model.Telefono_celular_estudiante),
	}

	// put query
	if err := Repositories.PutOneEstudiante(&model, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.PutMyEstudianteOutput(model))
	ApiHelpers.RespondJSON(c, 200, model)
}

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

// @Summary Lista de evaluadores
// @Description Lista todos los evaluadores
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListEvaluadoresSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores [get]
func ListEvaluadores(c *gin.Context) {
	var evaluadores []Models.Evaluador
	if err := Repositories.GetAllEvaluadores(&evaluadores); err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			ApiHelpers.RespondJSON(c, 200, evaluadores)
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.ListEvaluadoresOutput(container))
	ApiHelpers.RespondJSON(c, 200, evaluadores)
}

// @Summary Obtiene un evaluador
// @Description Obtiene un evaluador según su UUID
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a buscar"
// @Success 200 {object} Swagger.GetOneEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores/{uuid_evaluador} [get]
func GetOneEvaluador(c *gin.Context) {
	id := c.Params.ByName("id")

	var evaluador Models.Evaluador
	if err := Repositories.GetOneEvaluador(&evaluador, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.GetOneEvaluadorOutput(container))
	ApiHelpers.RespondJSON(c, 200, evaluador)
}

// @Summary Agrega un nuevo evaluador
// @Description Genera un nuevo evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_evaluador     body    Request.AddNewEvaluador     true        "Evaluador a agregar"
// @Success 200 {object} Swagger.AddNewEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores [post]
func AddNewEvaluador(c *gin.Context) {
	var input Request.AddNewEvaluador
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.AddNewEvaluador(&input)

	model := Models.Evaluador{
		Id_rol:                       *input.Id_rol,
		Rut_evaluador:                *input.Rut_evaluador,
		Nombres_evaluador:            *input.Nombres_evaluador,
		Apellidos_evaluador:          *input.Apellidos_evaluador,
		Hash_contrasena_evaluador:    *input.Hash_contrasena_evaluador,
		Correo_electronico_evaluador: *input.Correo_electronico_evaluador,
		Telefono_fijo_evaluador:      *input.Telefono_fijo_evaluador,
		Telefono_celular_evaluador:   *input.Telefono_celular_evaluador,
		Recinto_evaluador:            *input.Recinto_evaluador,
		Cargo_evaluador:              *input.Cargo_evaluador,
	}

	err := Repositories.AddNewEvaluador(&model)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.AddNewEvaluadorOutput(model))
	ApiHelpers.RespondJSON(c, 200, model)
}

// @Summary Modifica un evaluador
// @Description Modifica un evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a modificar"
// @Param   input_actualiza_evaluador     body    Request.PutOneEvaluador     true        "Evaluador a modificar"
// @Success 200 {object} Swagger.PutOneEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores/{uuid_evaluador} [put]
func PutOneEvaluador(c *gin.Context) {
	id := c.Params.ByName("id")

	var input Request.PutOneEvaluador
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.PutOneEvaluador(&input)

	var evaluador Models.Evaluador
	if err := Repositories.GetOneEvaluador(&evaluador, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Evaluador not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	evaluador = Models.Evaluador{
		Id:                           evaluador.Id,
		Id_rol:                       evaluador.Id_rol,
		Rol_evaluador:                evaluador.Rol_evaluador,
		Rut_evaluador:                Utils.CheckNullString(input.Rut_evaluador, evaluador.Rut_evaluador),
		Nombres_evaluador:            Utils.CheckNullString(input.Nombres_evaluador, evaluador.Nombres_evaluador),
		Apellidos_evaluador:          Utils.CheckNullString(input.Apellidos_evaluador, evaluador.Apellidos_evaluador),
		Hash_contrasena_evaluador:    Utils.CheckNullString(input.Hash_contrasena_evaluador, evaluador.Hash_contrasena_evaluador),
		Correo_electronico_evaluador: Utils.CheckNullString(input.Correo_electronico_evaluador, evaluador.Correo_electronico_evaluador),
		Telefono_fijo_evaluador:      Utils.CheckNullString(input.Telefono_fijo_evaluador, evaluador.Telefono_fijo_evaluador),
		Telefono_celular_evaluador:   Utils.CheckNullString(input.Telefono_celular_evaluador, evaluador.Telefono_celular_evaluador),
		Recinto_evaluador:            Utils.CheckNullString(input.Recinto_evaluador, evaluador.Recinto_evaluador),
		Cargo_evaluador:              Utils.CheckNullString(input.Cargo_evaluador, evaluador.Cargo_evaluador),
	}

	if err := Repositories.PutOneEvaluador(&evaluador, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.PutOneEvaluadorOutput(model))
	ApiHelpers.RespondJSON(c, 200, evaluador)
}

// @Summary Elimina un evaluador
// @Description Elimina un evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a eliminar"
// @Success 200 {object} Swagger.DeleteEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/evaluadores/{uuid_evaluador} [delete]
func DeleteEvaluador(c *gin.Context) {
	id := c.Params.ByName("id")

	var evaluador Models.Evaluador
	if err := Repositories.GetOneEvaluador(&evaluador, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	if err := Repositories.DeleteEvaluador(&evaluador, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.DeleteEvaluadorOutput(container))
	ApiHelpers.RespondJSON(c, 200, evaluador)
}

// @Summary Obtiene el perfil del evaluador
// @Description Obtiene el perfil del evaluador según su token
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetMyEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me [get]
func GetMyEvaluador(c *gin.Context) {
	id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var evaluador Models.Evaluador
	if err := Repositories.GetOneEvaluador(&evaluador, id_evaluador); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.GetMyEvaluadorOutput(container))
	ApiHelpers.RespondJSON(c, 200, evaluador)
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil del propio evaluador con los datos entregados
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   input_actualiza_evaluador     body    Request.PutMyEvaluador     true        "Evaluador a modificar"
// @Success 200 {object} Swagger.PutMyEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me [put]
func PutMyEvaluador(c *gin.Context) {
	id := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var input Request.PutMyEvaluador
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.PutMyEvaluador(&input)

	var evaluador Models.Evaluador
	if err := Repositories.GetOneEvaluador(&evaluador, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	evaluador = Models.Evaluador{
		Id:                           evaluador.Id,
		Id_rol:                       evaluador.Id_rol,
		Rol_evaluador:                evaluador.Rol_evaluador,
		Rut_evaluador:                evaluador.Rut_evaluador,
		Nombres_evaluador:            evaluador.Nombres_evaluador,
		Apellidos_evaluador:          evaluador.Apellidos_evaluador,
		Hash_contrasena_evaluador:    Utils.CheckNullString(input.Hash_contrasena_evaluador, evaluador.Hash_contrasena_evaluador),
		Correo_electronico_evaluador: evaluador.Correo_electronico_evaluador,
		Telefono_fijo_evaluador:      Utils.CheckNullString(input.Telefono_fijo_evaluador, evaluador.Telefono_fijo_evaluador),
		Telefono_celular_evaluador:   Utils.CheckNullString(input.Telefono_celular_evaluador, evaluador.Telefono_celular_evaluador),
		Recinto_evaluador:            evaluador.Recinto_evaluador,
		Cargo_evaluador:              Utils.CheckNullString(input.Cargo_evaluador, evaluador.Cargo_evaluador),
	}

	if err := Repositories.PutOneEvaluador(&evaluador, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.PutMyEvaluadorOutput(model))
	ApiHelpers.RespondJSON(c, 200, evaluador)
}

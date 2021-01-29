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
	var estudiantes []Models.Estudiante
	if err := Repositories.GetAllEstudiantes(&estudiantes); err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			ApiHelpers.RespondJSON(c, 200, estudiantes)
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.ListEstudiantesOutput(container))
	ApiHelpers.RespondJSON(c, 200, estudiantes)
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

	var estudiante Models.Estudiante
	if err := Repositories.GetOneEstudiante(&estudiante, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.GetOneEstudianteOutput(container))
	ApiHelpers.RespondJSON(c, 200, estudiante)
}

// @Summary Agrega un nuevo estudiante
// @Description Genera un nuevo estudiante con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_estudiante     body    Request.AddNewEstudiante     true        "Estudiante a agregar"
// @Success 200 {object} Swagger.AddNewEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes [post]
func AddNewEstudiante(c *gin.Context) {
	var input Request.AddNewEstudiante
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.AddNewEstudiante(&input)

	estudiante := Models.Estudiante{
		Id_rol:                        *input.Id_rol,
		Rut_estudiante:                *input.Rut_estudiante,
		Nombres_estudiante:            *input.Nombres_estudiante,
		Apellidos_estudiante:          *input.Apellidos_estudiante,
		Hash_contrasena_estudiante:    *input.Hash_contrasena_estudiante,
		Correo_electronico_estudiante: *input.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      *input.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   *input.Telefono_celular_estudiante,
	}

	if err := Repositories.AddNewEstudiante(&estudiante); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.AddNewEstudianteOutput(model))
	ApiHelpers.RespondJSON(c, 200, estudiante)
}

// @Summary Modifica un estudiante
// @Description Modifica un estudiante con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a modificar"
// @Param   input_actualiza_estudiante     body    Request.PutOneEstudiante     true        "Estudiante a modificar"
// @Success 200 {object} Swagger.PutOneEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes/{uuid_estudiante} [put]
func PutOneEstudiante(c *gin.Context) {
	id := c.Params.ByName("id")

	var input Request.PutOneEstudiante
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.PutOneEstudiante(&input)

	var estudiante Models.Estudiante
	if err := Repositories.GetOneEstudiante(&estudiante, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	estudiante = Models.Estudiante{
		Id:                            estudiante.Id,
		Id_rol:                        estudiante.Id_rol,
		Rol_estudiante:                estudiante.Rol_estudiante,
		Rut_estudiante:                Utils.CheckNullString(input.Rut_estudiante, estudiante.Rut_estudiante),
		Nombres_estudiante:            Utils.CheckNullString(input.Nombres_estudiante, estudiante.Nombres_estudiante),
		Apellidos_estudiante:          Utils.CheckNullString(input.Apellidos_estudiante, estudiante.Apellidos_estudiante),
		Hash_contrasena_estudiante:    Utils.CheckNullString(input.Hash_contrasena_estudiante, estudiante.Hash_contrasena_estudiante),
		Correo_electronico_estudiante: Utils.CheckNullString(input.Correo_electronico_estudiante, estudiante.Correo_electronico_estudiante),
		Telefono_fijo_estudiante:      Utils.CheckNullString(input.Telefono_fijo_estudiante, estudiante.Telefono_fijo_estudiante),
		Telefono_celular_estudiante:   Utils.CheckNullString(input.Telefono_celular_estudiante, estudiante.Telefono_celular_estudiante),
	}

	if err := Repositories.PutOneEstudiante(&estudiante, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.PutOneEstudianteOutput(model))
	ApiHelpers.RespondJSON(c, 200, estudiante)
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
	id := c.Params.ByName("id")

	var estudiante Models.Estudiante
	if err := Repositories.GetOneEstudiante(&estudiante, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	if err := Repositories.DeleteEstudiante(&estudiante, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.DeleteEstudianteOutput(container))
	ApiHelpers.RespondJSON(c, 200, estudiante)
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
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var estudiante Models.Estudiante
	if err := Repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.GetMyEstudianteOutput(container))
	ApiHelpers.RespondJSON(c, 200, estudiante)
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil del propio estudiante con los datos entregados
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   input_actualiza_estudiante     body    Request.PutMyEstudiante     true        "Nuevos datos del estudiante a modificar"
// @Success 200 {object} Swagger.PutMyEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me [put]
func PutMyEstudiante(c *gin.Context) {
	id := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var input Request.PutMyEstudiante
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.PutMyEstudiante(&input)

	var estudiante Models.Estudiante
	if err := Repositories.GetOneEstudiante(&estudiante, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	estudiante = Models.Estudiante{
		Id:                            estudiante.Id,
		Id_rol:                        estudiante.Id_rol,
		Rol_estudiante:                estudiante.Rol_estudiante,
		Rut_estudiante:                estudiante.Rut_estudiante,
		Nombres_estudiante:            estudiante.Nombres_estudiante,
		Apellidos_estudiante:          estudiante.Apellidos_estudiante,
		Hash_contrasena_estudiante:    Utils.CheckNullString(input.Hash_contrasena_estudiante, estudiante.Hash_contrasena_estudiante),
		Correo_electronico_estudiante: estudiante.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      Utils.CheckNullString(input.Telefono_fijo_estudiante, estudiante.Telefono_fijo_estudiante),
		Telefono_celular_estudiante:   Utils.CheckNullString(input.Telefono_celular_estudiante, estudiante.Telefono_celular_estudiante),
	}

	if err := Repositories.PutOneEstudiante(&estudiante, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.PutMyEstudianteOutput(model))
	ApiHelpers.RespondJSON(c, 200, estudiante)
}

// @Summary Lista de estudiantes de un curso
// @Description Lista todos los estudiantes existentes en un curso
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListEstudiantesCursoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/cursos/:id_curso/estudiantes [get]
func ListEstudiantesCurso(c *gin.Context) {
	// id := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")
	id_curso := c.Params.ByName("id")

	var estudiantes []Models.Estudiante
	if err := Repositories.GetAllEstudiantesCurso(&estudiantes, id_curso); err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			ApiHelpers.RespondJSON(c, 200, estudiantes)
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	ApiHelpers.RespondJSON(c, 200, estudiantes)
}

// @Summary Lista de estudiantes de un curso sin grupo
// @Description Lista todos los estudiantes existentes en un curso sin grupo
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListEstudiantesCursoSinGrupoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/cursos/:id_curso/estudiantes/sin-grupo [get]
func ListEstudiantesCursoSinGrupo(c *gin.Context) {
	// id := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")
	id_curso := c.Params.ByName("id")

	var estudiantes []Models.Estudiante
	if err := Repositories.GetAllEstudiantesCursoSinGrupo(&estudiantes, id_curso); err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			ApiHelpers.RespondJSON(c, 200, estudiantes)
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	ApiHelpers.RespondJSON(c, 200, estudiantes)
}

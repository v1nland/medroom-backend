package rol

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/api_helpers"
	"medroom-backend/formats/Output"
	"medroom-backend/formats/f_input"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Agrega un nuevo rol
// @Description Genera un nuevo rol con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_rol     body    Request.AddNewRol     true        "Rol a agregar"
// @Success 200 {object} Swagger.AddNewRolSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/roles [post]
func AddNewRol(c *gin.Context) {
	var input Request.AddNewRol

	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.AddNewRol(&input)

	rol := models.Rol{
		Nombre_rol: *input.Nombre_rol,
	}

	if err := repositories.AddNewRol(&rol); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, Output.AddNewRolOutput(rol))
}

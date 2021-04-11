package rol

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type addRequest struct {
	Nombre_rol *string `json:"nombre_rol"`
}

func addRequestParse(u *addRequest) {
	if u.Nombre_rol != nil {
		*u.Nombre_rol = strings.TrimSpace(*u.Nombre_rol)
		*u.Nombre_rol = strings.ToUpper(*u.Nombre_rol)
		*u.Nombre_rol = utils.RemoveAccents(*u.Nombre_rol)
	}
}

// @Summary Agrega un nuevo rol
// @Description Genera un nuevo rol con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_rol     body    Request.AddNewRol     true        "Rol a agregar"
// @Success 200 {object} Swagger.AddNewRolSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/roles [post]
func Add(c *gin.Context) {
	var input addRequest

	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	addRequestParse(&input)

	rol := models.Rol{
		Nombre_rol: *input.Nombre_rol,
	}

	if err := repositories.AddNewRol(&rol); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, rol)
}

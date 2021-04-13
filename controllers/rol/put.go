package rol

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type putRequest struct {
	Nombre_rol *string `json:"nombre_rol"`
}

func putRequestParse(u *putRequest) {
	if u.Nombre_rol != nil {
		*u.Nombre_rol = strings.TrimSpace(*u.Nombre_rol)
		*u.Nombre_rol = strings.ToUpper(*u.Nombre_rol)
		*u.Nombre_rol = utils.RemoveAccents(*u.Nombre_rol)
	}
}

// @Summary Modifica un rol
// @Description Modifica un rol con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_rol     path    string     true        "Id del rol a modificar"
// @Param   input_actualiza_rol     body    putRequest     true        "Rol a modificar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/roles/{id_rol} [put]
func PutOneRol(c *gin.Context) {
	id_rol := c.Params.ByName("id_rol")

	var input putRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	putRequestParse(&input)

	var rol models.Rol
	if err := repositories.GetOneRol(&rol, id_rol); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	rol = models.Rol{
		Id:         rol.Id,
		Nombre_rol: utils.CheckNullString(input.Nombre_rol, rol.Nombre_rol),
	}

	if err := repositories.PutOneRol(&rol, id_rol); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, rol)
}

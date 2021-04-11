package periodo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type addRequest struct {
	Nombre_periodo *string `json:"nombre_periodo"`
}

func addRequestParse(u *addRequest) {
	if u.Nombre_periodo != nil {
		*u.Nombre_periodo = strings.TrimSpace(*u.Nombre_periodo)
		*u.Nombre_periodo = strings.ToUpper(*u.Nombre_periodo)
		*u.Nombre_periodo = utils.RemoveAccents(*u.Nombre_periodo)
	}
}

// @Summary Agrega un nuevo periodo
// @Description Genera un nuevo periodo con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_periodo     body    Request.Add     true        "Periodo a agregar"
// @Success 200 {object} Swagger.AddNewPeriodoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/periodos [post]
func Add(c *gin.Context) {
	var input addRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	addRequestParse(&input)

	periodo := models.Periodo{
		Id: *input.Nombre_periodo,
	}

	if err := repositories.AddNewPeriodo(&periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, periodo)
}

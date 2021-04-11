package periodo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type putRequest struct {
	Nombre_periodo *string `json:"nombre_periodo"`
}

func putRequestParse(u *putRequest) {
	if u.Nombre_periodo != nil {
		*u.Nombre_periodo = strings.TrimSpace(*u.Nombre_periodo)
		*u.Nombre_periodo = strings.ToUpper(*u.Nombre_periodo)
		*u.Nombre_periodo = utils.RemoveAccents(*u.Nombre_periodo)
	}
}

// @Summary Modifica un periodo
// @Description Modifica un periodo con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo a modificar"
// @Param   input_actualiza_periodo     body    Request.Put     true        "Periodo a modificar"
// @Success 200 {object} Swagger.PutOnePeriodoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/periodos/{id_periodo} [put]
func Put(c *gin.Context) {
	id := c.Params.ByName("id")

	var input putRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	putRequestParse(&input)

	var periodo models.Periodo
	if err := repositories.GetOnePeriodo(&periodo, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	periodo = models.Periodo{
		Id: periodo.Id,
		// Id: utils.CheckNullString(input.Id, periodo.Id),
	}

	if err := repositories.PutOnePeriodo(&periodo, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, periodo)
}

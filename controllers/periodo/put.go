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
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo a modificar"
// @Param   input_actualiza_periodo     body    putRequest     true        "Periodo a modificar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/periodos/{id_periodo} [put]
func Put(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")

	var input putRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	putRequestParse(&input)

	var periodo models.Periodo
	if err := repositories.GetOnePeriodo(&periodo, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	periodo = models.Periodo{
		Id: periodo.Id,
		// Id: utils.CheckNullString(input.Id, periodo.Id),
	}

	if err := repositories.PutOnePeriodo(&periodo, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, periodo)
}

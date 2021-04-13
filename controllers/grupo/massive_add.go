package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type massiveAddRequest struct {
	Grupos []struct {
		Id           *int    `json:"id"`
		Id_curso     *int    `json:"id_curso"`
		Nombre_grupo *string `json:"nombre_grupo"`
		Sigla_grupo  *string `json:"sigla_grupo"`
	} `json:"grupos"`
}

func massiveAddRequestParse(u *massiveAddRequest) {
	for i := 0; i < len(u.Grupos); i++ {
		if u.Grupos[i].Nombre_grupo != nil {
			*u.Grupos[i].Nombre_grupo = strings.TrimSpace(*u.Grupos[i].Nombre_grupo)
			*u.Grupos[i].Nombre_grupo = strings.ToUpper(*u.Grupos[i].Nombre_grupo)
			*u.Grupos[i].Nombre_grupo = utils.RemoveAccents(*u.Grupos[i].Nombre_grupo)
		}

		if u.Grupos[i].Sigla_grupo != nil {
			*u.Grupos[i].Sigla_grupo = strings.TrimSpace(*u.Grupos[i].Sigla_grupo)
			*u.Grupos[i].Sigla_grupo = strings.ToUpper(*u.Grupos[i].Sigla_grupo)
			*u.Grupos[i].Sigla_grupo = utils.RemoveAccents(*u.Grupos[i].Sigla_grupo)
		}
	}
}

// @Summary Agrega nuevos grupos de forma masiva
// @Description Genera nuevos grupos con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_grupo     body    massiveAddRequest     true        "Grupo a agregar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/grupos/carga-masiva [post]
func MassiveAdd(c *gin.Context) {
	var payload massiveAddRequest
	if err := c.ShouldBind(&payload); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	massiveAddRequestParse(&payload)

	var grupos_error []int
	for i := 0; i < len(payload.Grupos); i++ {
		grupo := models.Grupo{
			// Id:           *payload.Grupos[i].Sigla_grupo,
			Sigla_curso:  utils.IntToString(*payload.Grupos[i].Id_curso),
			Nombre_grupo: *payload.Grupos[i].Nombre_grupo,
			Sigla_grupo:  *payload.Grupos[i].Sigla_grupo,
		}

		if err := repositories.AddNewGrupo(&grupo); err != nil {
			grupos_error = append(grupos_error, *payload.Grupos[i].Id)
		}
	}

	if len(grupos_error) > 0 {
		api_helpers.RespondJson(c, 201, grupos_error)
	} else {
		api_helpers.RespondJson(c, 200, "ok")
	}
}

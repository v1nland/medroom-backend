package grupo

import (
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
	"medroom-backend/app/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type massive_add_input struct {
	Grupos []struct {
		Id           *int    `json:"id"`
		Id_curso     *int    `json:"id_curso"`
		Nombre_grupo *string `json:"nombre_grupo"`
		Sigla_grupo  *string `json:"sigla_grupo"`
	} `json:"grupos"`
}

func massive_add_format(u *massive_add_input) {
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
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_grupo     body    massive_add_input     true        "Grupo a agregar"
// @Success 200 {object} Swagger.AddNewGrupoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/grupos/carga-masiva [post]
func AddNewGrupos(c *gin.Context) {
	var payload massive_add_input
	if err := c.ShouldBind(&payload); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	massive_add_format(&payload)

	var grupos_error []int
	for i := 0; i < len(payload.Grupos); i++ {
		grupo := models.Grupo{
			Id:           *payload.Grupos[i].Id,
			Id_curso:     *payload.Grupos[i].Id_curso,
			Nombre_grupo: *payload.Grupos[i].Nombre_grupo,
			Sigla_grupo:  *payload.Grupos[i].Sigla_grupo,
		}

		if err := repositories.AddNewGrupo(&grupo); err != nil {
			grupos_error = append(grupos_error, *payload.Grupos[i].Id)
		}
	}

	if len(grupos_error) > 0 {
		api_helpers.RespondJSON(c, 201, grupos_error)
	} else {
		api_helpers.RespondJSON(c, 200, "ok")
	}
}

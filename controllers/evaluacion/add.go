package evaluacion

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type addRequest struct {
	Nombre_evaluacion *string `json:"nombre_evaluacion"`
}

func addRequestParse(u *addRequest) {
	if u.Nombre_evaluacion != nil {
		*u.Nombre_evaluacion = strings.TrimSpace(*u.Nombre_evaluacion)
		*u.Nombre_evaluacion = strings.ToUpper(*u.Nombre_evaluacion)
		*u.Nombre_evaluacion = utils.RemoveAccents(*u.Nombre_evaluacion)
	}
}

// @Summary Agrega una evaluación
// @Description Genera una evaluación para un grupo
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso"
// @Param   id_grupo     path    string     true        "Id del grupo"
// @Param   input_evaluacion     body    Request.Add     true        "Evaluacion a agregar"
// @Success 200 {array} Swagger.AddNewEvaluacionSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos/{id_curso}/grupos/{id_grupo}/evaluaciones [post]
func Add(c *gin.Context) {
	// id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")
	// id_curso := c.Params.ByName("id_curso")
	sigla_grupo := c.Params.ByName("id_grupo")

	var input addRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	addRequestParse(&input)

	evaluacion := models.Evaluacion{
		Sigla_grupo:       sigla_grupo,
		Nombre_evaluacion: *input.Nombre_evaluacion,
	}

	if err := repositories.AddNewEvaluacion(&evaluacion); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEvaluacionesEstudiante(model))
	api_helpers.RespondJSON(c, 200, evaluacion)
}

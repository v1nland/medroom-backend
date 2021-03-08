package evaluacion

import (
	"medroom-backend/app/Messages/Request"
	"medroom-backend/app/Utils"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/formats/f_input"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Agrega una evaluación
// @Description Genera una evaluación para un grupo
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso"
// @Param   id_grupo     path    string     true        "Id del grupo"
// @Param   input_evaluacion     body    Request.AddNewEvaluacion     true        "Evaluacion a agregar"
// @Success 200 {array} Swagger.AddNewEvaluacionSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos/{id_curso}/grupos/{id_grupo}/evaluaciones [post]
func AddNewEvaluacion(c *gin.Context) {
	// id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")
	// id_curso := c.Params.ByName("id_curso")
	id_grupo := c.Params.ByName("id_grupo")

	var input Request.AddNewEvaluacion
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.AddNewEvaluacion(&input)

	evaluacion := models.Evaluacion{
		Id_grupo:          Utils.ConvertStringToInt(id_grupo),
		Nombre_evaluacion: *input.Nombre_evaluacion,
	}

	if err := repositories.AddNewEvaluacion(&evaluacion); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEvaluacionesEstudiante(model))
	api_helpers.RespondJSON(c, 200, evaluacion)
}

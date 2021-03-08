package puntaje

/*import (
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/formats/f_input"
	"medroom-backend/app/models"
	"medroom-backend/app/formats/f_output"
	"medroom-backend/app/repositories"
	"medroom-backend/app/RequestMessages"
	"medroom-backend/app/Utils"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de puntajes
// @Description Lista todos los puntajes
// @Tags Puntajes
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListPuntajesSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /puntajes [get]

func ListPuntajes(c *gin.Context) {
	// model container
	var container []models.Puntaje

	// query
	err := repositories.GetAllPuntajes(&container)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// output
	api_helpers.RespondJSON(c, 200, f_output.GetPuntajes(container))
}

// @Summary Obtiene un puntaje
// @Description Obtiene un puntaje según su Id
// @Tags Puntajes
// @Accept  json
// @Produce  json
// @Param   id_puntaje     path    string     true        "Id del puntaje a buscar"
// @Success 200 {object} Swagger.GetOnePuntajeSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /puntajes/{id_puntaje} [get]

func GetOnePuntaje(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container models.Puntaje

	// query
	err := repositories.GetOnePuntaje(&container, id)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// output
	api_helpers.RespondJSON(c, 200, f_output.GetOnePuntaje(container))
}

// @Summary Agrega un nuevo puntaje
// @Description Genera un nuevo puntaje con los datos entregados
// @Tags Puntajes
// @Accept  json
// @Produce  json
// @Param   input_puntaje     body    Request.AddNewPuntaje     true        "Puntaje a agregar"
// @Success 200 {object} Swagger.AddNewPuntajeSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /puntajes [post]

func AddNewPuntaje(c *gin.Context) {
	// input container
	var container Request.AddNewPuntaje

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	// format input
	input.AddNewPuntajeinput(&container)

	// generate model entity
	model_container := models.Puntaje{
		Id_evaluacion:              container.Id_evaluacion,
		Nombre_competencia_puntaje: container.Nombre_competencia_puntaje,
		Calificacion_puntaje:       container.Calificacion_puntaje,
		Feedback_puntaje:           container.Feedback_puntaje,
	}

	// query
	err := repositories.AddNewPuntaje(&model_container)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// output
	api_helpers.RespondJSON(c, 200, f_output.AddNewPuntaje(model_container))
}

// @Summary Modifica un puntaje
// @Description Modifica un puntaje con los datos entregados
// @Tags Puntajes
// @Accept  json
// @Produce  json
// @Param   id_puntaje     path    string     true        "Id del puntaje a modificar"
// @Param   input_actualiza_puntaje     body    Request.PutOnePuntaje     true        "Puntaje a modificar"
// @Success 200 {object} Swagger.PutOnePuntajeSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /puntajes/{id_puntaje} [put]

func PutOnePuntaje(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container Request.PutOnePuntaje

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	// format input
	input.PutOnePuntajeinput(&container)

	// generate model entity
	var model_container models.Puntaje

	// get query
	err := repositories.GetOnePuntaje(&model_container, id)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = models.Puntaje{
		Id_evaluacion:              Utils.CheckUpdatedInt(container.Id_evaluacion, model_container.Id_evaluacion),
		Nombre_competencia_puntaje: Utils.CheckUpdatedString(container.Nombre_competencia_puntaje, model_container.Nombre_competencia_puntaje),
		Calificacion_puntaje:       Utils.CheckUpdatedInt(container.Calificacion_puntaje, model_container.Calificacion_puntaje),
		Feedback_puntaje:           Utils.CheckUpdatedString(container.Feedback_puntaje, model_container.Feedback_puntaje),
	}

	// update foreign entity
	// err = repositories.GetOneEvaluacion(&model_container.Evaluacion_puntaje, Utils.ConvertIntToString(model_container.Id_evaluacion))
	// if err != nil {
	// 	api_helpers.RespondError(c, 500, "default")
	// 	return
	// }

	// put query
	err = repositories.PutOnePuntaje(&model_container, id)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// output
	api_helpers.RespondJSON(c, 200, f_output.PutOnePuntaje(model_container))
}

// @Summary Elimina un puntaje
// @Description Elimina un puntaje con los datos entregados
// @Tags Puntajes
// @Accept  json
// @Produce  json
// @Param   id_puntaje     path    string     true        "Id del puntaje a eliminar"
// @Success 200 {object} Swagger.DeletePuntajeSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /puntajes/{id_puntaje} [delete]

func DeletePuntaje(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container models.Puntaje

	// get query
	err := repositories.GetOnePuntaje(&container, id)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = repositories.DeletePuntaje(&container, id)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// output
	api_helpers.RespondJSON(c, 200, f_output.DeletePuntaje(container))
} */

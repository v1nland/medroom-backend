package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

type getOneGrupoEstudianteResponse struct {
	// Sigla_curso        string       `json:"sigla_curso"`
	// Id_periodo_curso   string       `json:"id_periodo_curso"`
	// Sigla_grupo        string       `json:"sigla_grupo"`
	// Evaluaciones_grupo []Evaluacion `json:"evaluaciones_grupo"`
	// Evaluadores_grupo  []Evaluador  `json:"evaluadores_grupo"`
	// Estudiantes_grupo  []Estudiante `json:"estudiantes_grupo"`
	// Nombre_grupo       string       `json:"nombre_grupo"`
	// CreatedAt          time.Time    `json:"created_at"`
	// UpdatedAt          time.Time    `json:"updated_at"`
}

// @Summary Obtiene un grupo de un estudiante
// @Description Obtiene un grupo de un estudiante según su token
// @Tags Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /estudiantes/me/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo} [get]
func GetOneGrupoEstudiante(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")
	sigla_grupo := c.Params.ByName("sigla_grupo")
	id_estudiante := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var grupo models.Grupo
	if err := repositories.GetOneGrupoEstudiante(&grupo, sigla_grupo, sigla_curso, id_periodo, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, grupo)
}

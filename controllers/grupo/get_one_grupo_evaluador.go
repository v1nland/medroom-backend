package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type estudianteGetOneGrupoEvaluador struct {
	Nombres_estudiante            string `json:"nombres_estudiante"`
	Apellidos_estudiante          string `json:"apellidos_estudiante"`
	Correo_electronico_estudiante string `json:"correo_electronico_estudiante" gorm:"unique;not null"`
}

type evaluacionGetOneGrupoEvaluador struct {
	Id                int       `json:"id"`
	Nombre_evaluacion string    `json:"nombre_evaluacion"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type getOneGrupoEvaluadorResponse struct {
	Sigla_curso        string                           `json:"sigla_curso"`
	Id_periodo_curso   string                           `json:"id_periodo_curso"`
	Nombre_curso       string                           `json:"nombre_curso"`
	Sigla_grupo        string                           `json:"sigla_grupo"`
	Nombre_grupo       string                           `json:"nombre_grupo"`
	Evaluaciones_grupo []evaluacionGetOneGrupoEvaluador `json:"evaluaciones_grupo"` //id, nombre, created_at, updated_at
	Evaluadores_grupo  []estudianteGetOneGrupoEvaluador `json:"estudiantes_grupo"`  //id,nombre,apellido,correo
}

// @Summary Obtiene un grupo de un evaluador
// @Description Obtiene un grupo de un evaluador según su token
// @Tags Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /evaluadores/me/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo} [get]
func GetOneGrupoEvaluador(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")
	sigla_grupo := c.Params.ByName("sigla_grupo")
	id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var grupo models.Grupo
	if err := repositories.GetOneGrupoEvaluador(&grupo, sigla_grupo, sigla_curso, id_periodo, id_evaluador); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, grupo.Sigla_curso, grupo.Id_periodo_curso); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, grupo)
}

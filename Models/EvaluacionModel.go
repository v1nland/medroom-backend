package Models

import (
	"time"
)

type Evaluacion struct {
	Id                                      int       `json:"id"`
	Id_estudiante                           string    `json:"id_estudiante" sql:"type:uuid REFERENCES public.estudiantes(id)"`
	Id_evaluador                            string    `json:"id_evaluador" sql:"type:uuid REFERENCES public.evaluadores(id)"`
	Evaluador_evaluacion                    Evaluador `json:"evaluador_evaluacion" gorm:"foreignKey:Id_evaluador"`
	Id_periodo                              int       `json:"id_periodo" sql:"type:int REFERENCES public.periodos(id)"`
	Periodo_evaluacion                      Periodo   `json:"periodo_evaluacion" gorm:"foreignKey:Id_periodo"`
	Puntajes_evaluacion                     []Puntaje `json:"puntajes_evaluacion" gorm:"foreignKey:Id_evaluacion"`
	Nombre_evaluacion                       string    `json:"nombre_evaluacion"`
	Entorno_clinico_evaluacion              string    `json:"entorno_clinico_evaluacion"`
	Paciente_evaluacion                     string    `json:"paciente_evaluacion"`
	Asunto_principal_consulta_evaluacion    string    `json:"asunto_principal_consulta_evaluacion"`
	Complejidad_caso_evaluacion             string    `json:"complejidad_caso_evaluacion"`
	Numero_observaciones_previas_evaluacion string    `json:"numero_observaciones_previas_evaluacion"`
	Categoria_observador_evaluacion         string    `json:"categoria_observador_evaluacion"`
	Observacion_calificacion_evaluacion     string    `json:"observacion_calificacion_evaluacion"`
	Tiempo_utilizado_evaluacion             int       `json:"tiempo_utilizado_evaluacion"`
	CreatedAt                               time.Time `json:"created_at"`
	UpdatedAt                               time.Time `json:"updated_at"`
}

func (u *Evaluacion) TableName() string {
	return "public.evaluaciones"
}

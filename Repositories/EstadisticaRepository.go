package Repositories

import (
	"medroom-backend/Config"
	"medroom-backend/Messages/Query"
)

func PromedioCalificacionEvaluacion(u *[]Query.PromedioCalificacionEvaluacion, id_evaluacion string, id_grupo string) (err error) {
	if err := Config.DB.Raw(`select 
														es.id as id_estudiante,
														es.nombres_estudiante,
														es.apellidos_estudiante,
														ev.nombre_evaluacion,
														avg(p.calificacion_puntaje) as promedio_estudiante
													from 
														puntajes p,
														calificaciones_estudiantes ce,
														evaluaciones ev,
														estudiantes es
													where
														p.id_calificacion_estudiante = ce.id and
														ce.id_evaluacion = ev.id and 
														ce.id_estudiante = es.id and 
														ev.id = ? and
														ev.id_grupo = ?
													group by 
														p.id_calificacion_estudiante,
														ev.id,
														ev.nombre_evaluacion,
														es.id,
														es.nombres_estudiante,
														es.apellidos_estudiante`, id_evaluacion, id_grupo).First(u).Error; err != nil {
		return err
	}
	return nil
}

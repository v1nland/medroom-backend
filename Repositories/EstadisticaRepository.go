package Repositories

import (
	"medroom-backend/Config"
	"medroom-backend/Messages/Query"
)

func CalificacionesEstudiantePorCompetencia(u *[]Query.CalificacionesEstudiantePorCompetencia, id_grupo string, id_estudiante string) (err error) {
	if err := Config.DB.Raw(`select t1.*, coalesce(t2.calificacion_puntaje, 0) as calificacion_puntaje_estudiante
													from 
														(select 
															ev.id as id_evaluacion,
															ev.nombre_evaluacion,
															p.id_competencia,
															avg(p.calificacion_puntaje) as promedio_calificacion_puntaje
														from 
															evaluaciones ev
															left outer join calificaciones_estudiantes ce on ev.id = ce.id_evaluacion 
															left outer join puntajes p on ce.id = p.id_calificacion_estudiante
														where 
															ev.id_grupo = ?
														group by ev.id, ev.nombre_evaluacion, p.id_competencia
														order by ev.id) as t1
														left outer join (select 
																			ev.id as id_evaluacion,
																			ev.nombre_evaluacion,
																			p.id_competencia,
																			p.calificacion_puntaje 
																		from 
																			evaluaciones ev
																			left outer join calificaciones_estudiantes ce on ev.id = ce.id_evaluacion 
																			left outer join puntajes p on ce.id = p.id_calificacion_estudiante
																			join estudiantes e on ce.id_estudiante = e.id 
																		where 
																			ev.id_grupo = ?
																			and ce.id_estudiante = ?
																		order by ev.id) as t2
														on (t1.id_evaluacion = t2.id_evaluacion and t1.id_competencia = t2.id_competencia)
													order by id_evaluacion`, id_grupo, id_grupo, id_estudiante).First(u).Error; err != nil {
		return err
	}

	return nil
}

func CalificacionesEstudiantePorEvaluacion(u *[]Query.CalificacionesEstudiantePorEvaluacion, id_grupo string, id_estudiante string) (err error) {
	if err := Config.DB.Raw(`select t1.*, coalesce(t2.calificacion_puntaje, 0) as promedio_calificacion_puntaje_estudiante
													from 
														(select 
															ev.id as id_evaluacion,
															ev.nombre_evaluacion,
															avg(p.calificacion_puntaje) as promedio_calificacion_puntaje_grupo
														from 
															evaluaciones ev
															left outer join calificaciones_estudiantes ce on ev.id = ce.id_evaluacion 
															left outer join puntajes p on ce.id = p.id_calificacion_estudiante
														where 
															ev.id_grupo = ?
														group by ev.id, ev.nombre_evaluacion
														order by ev.id) as t1
														left outer join (select 
																			ev.id as id_evaluacion,
																			ev.nombre_evaluacion,
																			avg(p.calificacion_puntaje) as calificacion_puntaje
																		from 
																			evaluaciones ev
																			left outer join calificaciones_estudiantes ce on ev.id = ce.id_evaluacion 
																			left outer join puntajes p on ce.id = p.id_calificacion_estudiante
																			join estudiantes e on ce.id_estudiante = e.id 
																		where 
																			ev.id_grupo = ?
																			and ce.id_estudiante = ?
																		group by ev.id, ev.nombre_evaluacion 
																		order by ev.id) as t2
														on (t1.id_evaluacion = t2.id_evaluacion)
													order by id_evaluacion`, id_grupo, id_grupo, id_estudiante).First(u).Error; err != nil {
		return err
	}

	return nil
}

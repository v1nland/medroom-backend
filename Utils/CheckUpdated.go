package Utils

/*
	*
	*  FUNCIONES CheckUpdatedString y CheckUpdatedInt
	*
    *  ESTAS FUNCIONES PERMITEN SABER SI UNA VARIABLE
	*  NECESITA SER ACTUALIZADA EN UNA REQUEST PUT
	*  COMPARA SI EL NUEVO VALOR ES DISTINTO A VACÍO Y RETORNA EL NUEVO VALOR
    *
*/
func CheckUpdatedString(new_value string, old_value string) (updated string) {
	if new_value != old_value && new_value != "" {
		return new_value
	} else {
		return old_value
	}
}

func CheckUpdatedInt(new_value int, old_value int) (updated int) {
	if new_value != old_value && new_value != 0 {
		return new_value
	} else {
		return old_value
	}
}

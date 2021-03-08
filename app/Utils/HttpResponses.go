package utils

/*
	*
	*  FUNCIÓN StatusMessage
	*
    *  ESTA FUNCIÓN RETORNA EL MENSAJE DE RESPUESTA CORRESPONDIENTE A
	*  UN HTTP STATUS CODE.
	*  ESTÁ DISEÑADA PARA USO DE LAS RESPUESTAS DE LOS SERVICIOS (APIHELPERS/RESPONSE.GO).
    *
*/
func StatusMessage(status int) string {
	switch status {
	case 200:
		return "OK"
	case 301:
		return "Moved permanently"
	case 308:
		return "Permanent redirect"
	case 400:
		return "Bad request"
	case 401:
		return "Unauthorized"
	case 404:
		return "Not found"
	case 500:
		return "Internal server error"
	default:
		return "Unhandled error"
	}
}

/*
	*
	*  FUNCIÓN StatusDescription
	*
    *  ESTA FUNCIÓN RETORNA LA DESCRIPCIÓN DEL ERROR
	*
	*  ESTÁ DISEÑADA PARA USO DE LAS RESPUESTAS DE LOS SERVICIOS (APIHELPERS/RESPONSE.GO).
    *  ESTA FUNCIÓN SOLO SE LLAMADA SI NO SE ESPECIFICA UN MENSAJE CUSTOMIZADO
*/
func StatusDescription(status int) string {
	switch status {
	case 200:
		return "Resource OK"
	case 301:
		return "Redirected as GET method"
	case 308:
		return "Redirected keeping original method"
	case 400:
		return "Invalid or missing JSON input"
	case 401:
		return "Invalid or missing authorization header"
	case 404:
		return "The requested resource was not found"
	case 500:
		return "There was an error communicating with the database"
	default:
		return "Invalid status code"
	}
}

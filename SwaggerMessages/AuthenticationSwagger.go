package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type AuthenticationSwagger struct {
	Status bool                            `json:"status"`
	Meta   string                          `json:"meta"`
	Data   ResponseMessages.Authentication `json:"data"`
}

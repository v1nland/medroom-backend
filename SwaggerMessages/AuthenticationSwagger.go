package SwaggerMessages

import "medroom-backend/Messages/Response"

type AuthenticationSwagger struct {
	Status bool                    `json:"status"`
	Meta   string                  `json:"meta"`
	Data   Response.Authentication `json:"data"`
}

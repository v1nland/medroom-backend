package Response

type ListRolesResponse struct {
	Id         int    `json:"id"`
	Nombre_rol string `json:"nombre_rol"`
}

type GetOneRolResponse struct {
	Id         int    `json:"id"`
	Nombre_rol string `json:"nombre_rol"`
}

type AddNewRolResponse struct {
	Id int `json:"id"`
}

type PutOneRolResponse struct {
	Id int `json:"id"`
}

type DeleteRolResponse struct {
	Id int `json:"id"`
}

package Request

type ListRoles struct {
}

type GetOneRol struct {
}

type AddNewRol struct {
	Nombre_rol *string `json:"nombre_rol"`
}

type PutOneRol struct {
	Nombre_rol *string `json:"nombre_rol"`
}

type DeleteRol struct {
}

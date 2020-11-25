package RequestMessages

type ListRolsPayload struct {
}

type GetOneRolPayload struct {
}

type AddNewRolPayload struct {
	Nombre_rol string `json:"nombre_rol"`
}

type PutOneRolPayload struct {
	Nombre_rol string `json:"nombre_rol"`
}

type DeleteRolPayload struct {
}

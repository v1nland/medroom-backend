package Models

type Competencia struct {
	Id                 string `json:"id" sql:"primary_key" gorm:"primaryKey:unique;not null"`
	Nombre_competencia string `json:"nombre_competencia" gorm:"unique;not null"`
}

func (u *Competencia) TableName() string {
	return "public.competencias"
}

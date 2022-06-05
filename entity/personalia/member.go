package personalia

type Member struct {
	ID   string `json:"id" db:"id" validate:"required,max=20"`
	Name string `json:"name" db:"name" validate:"max=255"`
}

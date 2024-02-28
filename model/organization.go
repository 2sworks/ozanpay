package model

type Organization struct {
	BaseModel
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	Website       string `json:"website"`
	ContactPerson string `json:"contact_person"`
}

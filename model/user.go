package model

type UserRole int

const (
	UserModelUser UserRole = iota + 1
	UserModelSeller
	UserModelOrganization
	UserModelAdmin
	UserModelSuperAdmin
)

type User struct {
	BaseModel
	Name     string   `json:"name"`
	Surname  string   `json:"surname"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

func (s User) String() string {
	return s.Name + " " + s.Surname
}

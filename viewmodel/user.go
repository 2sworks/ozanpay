package viewmodel

import "ozanpay/model"

type UserCreateVM struct {
	Name     string         `json:"name"`
	Surname  string         `json:"surname"`
	Email    string         `json:"email"`
	Phone    string         `json:"phone"`
	Password string         `json:"password"`
	Role     model.UserRole `json:"role"`
}

func (vm UserCreateVM) Validate() error {
	return nil
}

func (vm UserCreateVM) ToModel(m model.User) model.User {
	m.Name = vm.Name
	m.Surname = vm.Surname
	m.Email = vm.Email
	m.Phone = vm.Phone
	m.Password = vm.Password
	m.Role = vm.Role

	return m
}

type UserUpdateVM struct {
	ID       uint           `json:"id"`
	Name     string         `json:"name"`
	Surname  string         `json:"surname"`
	Email    string         `json:"email"`
	Phone    string         `json:"phone"`
	Password string         `json:"password"`
	Role     model.UserRole `json:"role"`
}

func (vm UserUpdateVM) Validate() error {
	return nil
}

func (vm UserUpdateVM) ToModel(m model.User) model.User {
	m.ID = vm.ID
	m.Name = vm.Name
	m.Surname = vm.Surname
	m.Email = vm.Email
	m.Phone = vm.Phone
	m.Password = vm.Password
	m.Role = vm.Role

	return m
}

type UserListVM struct {
	ID      uint           `json:"id"`
	Name    string         `json:"name"`
	Surname string         `json:"surname"`
	Email   string         `json:"email"`
	Phone   string         `json:"phone"`
	Role    model.UserRole `json:"role"`
}

func (vm UserListVM) ToVM(m model.User) UserListVM {
	vm.ID = m.ID
	vm.Name = m.Name
	vm.Surname = m.Surname
	vm.Email = m.Email
	vm.Phone = m.Phone
	vm.Role = m.Role

	return vm
}

type UserDetailVM struct {
	ID      uint           `json:"id"`
	Name    string         `json:"name"`
	Surname string         `json:"surname"`
	Email   string         `json:"email"`
	Phone   string         `json:"phone"`
	Role    model.UserRole `json:"role"`
}

func (vm UserDetailVM) ToVM(m model.User) UserDetailVM {
	vm.ID = m.ID
	vm.Name = m.Name
	vm.Surname = m.Surname
	vm.Email = m.Email
	vm.Phone = m.Phone
	vm.Role = m.Role

	return vm
}

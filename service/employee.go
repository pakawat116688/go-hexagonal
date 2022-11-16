package service

type EmployeeResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Tel    string `json:"tel"`
}

type EmployeeRequres struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Tel    string `json:"tel"`
	Status int    `json:"status"`
}

type EmployeeService interface {
	CreatedEmployee() error
	InsertEmployee(EmployeeRequres) (*EmployeeResponse, error)
	GetEmployee() ([]EmployeeResponse, error)
	GetEmployeeId(int) (*EmployeeResponse, error)
	DeleteEmployee() error
	DeleteEmployeeId(int) error
}
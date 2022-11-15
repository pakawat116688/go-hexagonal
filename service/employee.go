package service

type EmployeeResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Tel    string `json:"tel"`
}

type EmployeeService interface {
	CreatedEmployee() error
	InsertEmployee(string, int, string, int) error
	GetEmployee() ([]EmployeeResponse, error)
	GetEmployeeId(int) (*EmployeeResponse, error)
	DeleteEmployee() error
	DeleteEmployeeId(int) error
}
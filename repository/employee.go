package repository

type Employee struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Salary int    `db:"salary"`
	Tel    string `db:"tel"`
	Status int    `db:"status"`
}

type EmployeeInsert struct {
	Name   string `db:"name"`
	Salary int    `db:"salary"`
	Tel    string `db:"tel"`
	Status int    `db:"status"`
}

type EmployeeRepository interface {
	CreateTable() error
	InsertData(EmployeeInsert) (*Employee ,error)
	GetAll() ([]Employee, error)
	GetById(int) (*Employee, error)
	DeleteAll() error
	DeleteById(int) error
}

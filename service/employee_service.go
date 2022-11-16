package service

import (
	"database/sql"
	"github.com/pakawatkung/go-hexagonal/errs"
	"github.com/pakawatkung/go-hexagonal/logs"
	"github.com/pakawatkung/go-hexagonal/repository"
)

type employeeService struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeService(employeeRepo repository.EmployeeRepository) EmployeeService {
	return employeeService{employeeRepo: employeeRepo}
}

func (s employeeService) CreatedEmployee() error {
	err := s.employeeRepo.CreateTable()
	if err != nil {
		logs.Error(err)
		return errs.NewNotfoundError("cannot create table")
	}
	return nil
}

func (s employeeService) InsertEmployee(req EmployeeRequres) (*EmployeeResponse, error) {

	dataInsert := repository.EmployeeInsert{
		Name: req.Name,
		Salary: req.Salary,
		Tel: req.Tel,
		Status: req.Status,
	}
	data, err := s.employeeRepo.InsertData(dataInsert)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotfoundError("cannnot insert data")
	}

	response := EmployeeResponse{
		Id: data.Id,
		Name: data.Name,
		Tel: data.Tel,
	}

	return &response, nil
}

func (s employeeService) GetEmployee() ([]EmployeeResponse, error) {
	employees, err := s.employeeRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	employeeResponses := []EmployeeResponse{}
	for _, emp := range employees {
		empResponse := EmployeeResponse{
			Id:   emp.Id,
			Name: emp.Name,
			Tel:  emp.Tel,
		}
		employeeResponses = append(employeeResponses, empResponse)
	}
	return employeeResponses, nil
}

func (s employeeService) GetEmployeeId(id int) (*EmployeeResponse, error) {
	employee, err := s.employeeRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Error(err)
			return nil, errs.NewNotfoundError("service customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	employeeResponse := EmployeeResponse{
		Id:   employee.Id,
		Name: employee.Name,
		Tel:  employee.Tel,
	}
	return &employeeResponse, nil
}

func (s employeeService) DeleteEmployee() error {
	err := s.employeeRepo.DeleteAll()
	if err != nil {
		logs.Error(err)
		return errs.NewDeleteError()
	}
	return nil
}

func (s employeeService) DeleteEmployeeId(id int) error {
	err := s.employeeRepo.DeleteById(id)
	if err != nil {
		logs.Error(err)
		return errs.NewDeleteError()
	}
	return nil
}
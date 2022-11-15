package service

import (
	"database/sql"
	"errors"
	"log"

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
		log.Println(err)
		return err
	}
	return nil
}

func (s employeeService) InsertEmployee(name string, salary int, tel string, status int) error {
	err := s.employeeRepo.InsertData(name, salary, tel, status)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s employeeService) GetEmployee() ([]EmployeeResponse, error) {
	employees, err := s.employeeRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
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
			return nil, errors.New("employee not found")
		}
		log.Println(err)
		return nil, err
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
		log.Println(err)
		return err
	}
	return nil
}

func (s employeeService) DeleteEmployeeId(id int) error {
	err := s.employeeRepo.DeleteById(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
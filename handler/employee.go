package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pakawatkung/go-hexagonal/errs"
	"github.com/pakawatkung/go-hexagonal/service"
)

type employeeRest struct {
	empSrv service.EmployeeService
}

func NewEmployeeRest(empSrv service.EmployeeService) employeeRest {
	return employeeRest{empSrv: empSrv}
}

func (h employeeRest) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.empSrv.GetEmployee()
	if err != nil {
		handleEroor(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func (h employeeRest) GetEmployeeId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}
	employee, err := h.empSrv.GetEmployeeId(id)
	if err != nil {
		handleEroor(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

func (h employeeRest) InsertEmployee(w http.ResponseWriter, r *http.Request) {
	// parm := mux.Vars(r)
	// name := parm["name"]
	// salary, err := strconv.Atoi(parm["salary"])
	// if err != nil {
	// 	return
	// }
	// tel := parm["tel"]
	// status, err := strconv.Atoi(parm["status"])
	// if err != nil {
	// 	return
	// }

	if r.Header.Get("content-type") != "application/json" {
		handleEroor(w, errs.NewValidationError())
		return
	}

	req := service.EmployeeRequres{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleEroor(w, errs.NewValidationError())
		return
	}

	data, err := h.empSrv.InsertEmployee(req)
	if err != nil {
		handleEroor(w, err)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h employeeRest) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}
	err = h.empSrv.DeleteEmployeeId(id)
	if err != nil {
		handleEroor(w, err)
		return
	}
	msg := "Service Delete Success"
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

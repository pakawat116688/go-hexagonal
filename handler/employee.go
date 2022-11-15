package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pakawatkung/go-hexagonal/service"
)

type employeeRest struct {
	empSrv service.EmployeeService
}

func NewEmployeeRest(empSrv service.EmployeeService) employeeRest  {
	return employeeRest{empSrv : empSrv}
}

func (h employeeRest) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.empSrv.GetEmployee()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
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
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

func (h employeeRest) InsertEmployee(w http.ResponseWriter, r *http.Request) {
	parm := mux.Vars(r)
	name := parm["name"]
	salary, err := strconv.Atoi(parm["salary"])
	if err != nil {
		return
	}
	tel := parm["tel"]
	status, err := strconv.Atoi(parm["status"])
	if err != nil {
		return
	}
	err = h.empSrv.InsertEmployee(name,salary,tel,status)
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintln(w, err)
		return
	}
	msg := "Service Insert Success"
	json.NewEncoder(w).Encode(msg)
}

func (h employeeRest) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}
	err = h.empSrv.DeleteEmployeeId(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, err)
		return
	}
	msg := "Service Delete Success"
	json.NewEncoder(w).Encode(msg)
}
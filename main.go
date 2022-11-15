package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pakawatkung/go-hexagonal/handler"
	"github.com/pakawatkung/go-hexagonal/repository"
	"github.com/pakawatkung/go-hexagonal/service"
)

func main() {

	os.Remove("./mydata.db")
	db, err := sqlx.Open("sqlite3", "./mydata.db")
	if err != nil {
		println("Error Cannot Open the Database...")
		panic(err)
	}
	defer db.Close()

	employeeReposiotory := repository.NewEmployeeRepositoryDB(db)
	employeeService := service.NewEmployeeService(employeeReposiotory)
	employeeRests := handler.NewEmployeeRest(employeeService)

	err = employeeService.CreatedEmployee()
	if err != nil {
		panic(err)
	}

	// Default data
	err = employeeService.InsertEmployee("Doraemon", 900, "9xx-xxx-xxxx", 0)
	if err != nil {
		panic(err)
	}
	err = employeeService.InsertEmployee("Nobita", 999, "333-999-3333", 0)
	if err != nil {
		panic(err)
	}
	err = employeeService.InsertEmployee("Sisuga", 555, "3xx-xxx-xxxx", 0)
	if err != nil {
		panic(err)
	}
	err = employeeService.InsertEmployee("Giate", 600, "5xx-xxx-xxxx", 0)
	if err != nil {
		panic(err)
	}
	err = employeeService.InsertEmployee("Sunio", 1200, "999-999-9999", 0)
	if err != nil {
		panic(err)
	}


	router := mux.NewRouter()
	router.HandleFunc("/employee", employeeRests.GetEmployees).Methods(http.MethodGet)
	router.HandleFunc("/employee/{id:[0-9]+}", employeeRests.GetEmployeeId).Methods(http.MethodGet)
	router.HandleFunc("/employee/{name}/{salary}/{tel}/{status}", employeeRests.InsertEmployee)
	// employee/test/30000/06x-xxx-9999/0
	router.HandleFunc("/employee/delete/{id:[0-9]+}", employeeRests.DeleteEmployee)
	http.ListenAndServe(":8000", router)

}
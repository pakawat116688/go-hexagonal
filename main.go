package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pakawatkung/go-hexagonal/handler"
	"github.com/pakawatkung/go-hexagonal/logs"
	"github.com/pakawatkung/go-hexagonal/repository"
	"github.com/pakawatkung/go-hexagonal/service"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func inittimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func main() {

	inittimeZone()
	initConfig()

	os.Remove("./mydata.db")
	db, err := sqlx.Open(viper.GetString("db.driver"), viper.GetString("db.file"))
	if err != nil {
		println("Error Cannot Open the Database...")
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)

	employeeReposiotory := repository.NewEmployeeRepositoryDB(db)
	employeeService := service.NewEmployeeService(employeeReposiotory)
	employeeRests := handler.NewEmployeeRest(employeeService)

	err = employeeService.CreatedEmployee()
	if err != nil {
		panic(err)
	}

	// Default data
	data := []repository.EmployeeInsert{
		{Name:"Doraemon", Salary:900, Tel:"9xx-xxx-xxxx", Status:0},
		{Name:"Nobita", Salary:999, Tel:"333-999-3333", Status:0},
		{Name:"Sisuga", Salary:555, Tel:"3xx-xxx-xxxx", Status:0},
		{Name:"Giate", Salary:600, Tel:"5xx-xxx-xxxx", Status:0},
		{Name:"Sunio", Salary:1200, Tel:"999-999-9999", Status:1},
	}
	_, err = employeeService.InsertEmployee(service.EmployeeRequres(data[0]))
	if err != nil {
		panic(err)
	}
	_, err = employeeService.InsertEmployee(service.EmployeeRequres(data[1]))
	if err != nil {
		panic(err)
	}
	_, err = employeeService.InsertEmployee(service.EmployeeRequres(data[2]))
	if err != nil {
		panic(err)
	}
	_, err = employeeService.InsertEmployee(service.EmployeeRequres(data[3]))
	if err != nil {
		panic(err)
	}
	_, err = employeeService.InsertEmployee(service.EmployeeRequres(data[4]))
	if err != nil {
		panic(err)
	}


	router := mux.NewRouter()
	router.HandleFunc("/employee", employeeRests.GetEmployees).Methods(http.MethodGet)
	router.HandleFunc("/employee/{id:[0-9]+}", employeeRests.GetEmployeeId).Methods(http.MethodGet)
	router.HandleFunc("/employee/{name}/{salary:[0-9]+}/{tel}/{status:[0-1]}", employeeRests.InsertEmployee).Methods(http.MethodPost)
	// employee/test/30000/06x-xxx-9999/0
	router.HandleFunc("/employee/delete/{id:[0-9]+}", employeeRests.DeleteEmployee).Methods(http.MethodDelete)

	logs.Info("Development Employee Service Started at port "+viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetString("app.port")), router)

}
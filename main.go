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
	router.HandleFunc("/employee/{name}/{salary:[0-9]+}/{tel}/{status:[0-1]}", employeeRests.InsertEmployee)
	// employee/test/30000/06x-xxx-9999/0
	router.HandleFunc("/employee/delete/{id:[0-9]+}", employeeRests.DeleteEmployee)
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetString("app.port")), router)

}
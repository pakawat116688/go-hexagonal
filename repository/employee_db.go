package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type employeeRepositoryDB struct {
	db *sqlx.DB
}

func NewEmployeeRepositoryDB(db *sqlx.DB) EmployeeRepository {
	return employeeRepositoryDB{db: db}
}

func (r employeeRepositoryDB) CreateTable() error {

	create_table := `CREATE TABLE "Employee" (
		"id"	INTEGER,
		"name"	TEXT,
		"salary"	NUMERIC,
		"tel"	TEXT,
		"status"	NUMERIC,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	statement, err := r.db.Prepare(create_table)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}

	println("Created Table Success...")
	return nil
}

func (r employeeRepositoryDB) InsertData(name string, salary int, tel string, status int) error {

	query := `insert into employee(name, salary, tel, status)
		values (?, ?, ?, ?)`

	state, err := r.db.Exec(query, name, salary, tel, status)
	if err != nil {
		return err
	}
	affected, err := state.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("cannot insert")
	}

	println("Insert Data Success...")
	return nil
}

func (r employeeRepositoryDB) GetAll() ([]Employee, error) {
	employees := []Employee{}
	query := "select * from Employee"
	err := r.db.Select(&employees, query)
	if err != nil {
		return nil, err
	}
	println("GetAll Success...")
	return employees, nil
}

func (r employeeRepositoryDB) GetById(id int) (*Employee, error) {
	employee := Employee{}
	query := "select * from Employee where id=?"
	err := r.db.Get(&employee, query, id)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r employeeRepositoryDB) DeleteAll() error {
	
	query := "delete from Employee"
	state, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	affected, err := state.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("cannot delete")
	}

	println("delete all success...")

	return nil
}

func (r employeeRepositoryDB) DeleteById(id int) error {

	query := "delete from Employee where id=?"
	state, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	affected, err := state.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("cannot delete")
	}

	println("delete id ",id," success...")

	return nil
}
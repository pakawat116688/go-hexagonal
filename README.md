# go-hexagonal
Hexagonal Architecture Golang

- repository
    - employee.go    = port
    - employee_db.go = adapter

- service
    - employee.go          = port
    - employee_service.go  = adapter

- handler
    สร้าง RRSTful โดยใช้ mux

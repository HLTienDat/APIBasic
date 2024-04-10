package apimodel

import (
	"time"
)

type Account struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	EmployeeId string `json:"employeeId"`
	Role       int    `json:"role"`
}

type Employee struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Dob       time.Time `json:"dob"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CitizenId string    `json:"citizenId"`
	Address   string    `json:"address"`
}

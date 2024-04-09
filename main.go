package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	// "html"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "091372"
	dbname   = "Users"
)

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getSpecificUser)
	router.POST("/users", postUsers)
	router.DELETE("/users/:id", delUsers)
	router.PUT("/users/:id", updateUser)
	router.Run("localhost:8080")

}

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

func CreateEmployee(db *sql.DB, name string, dob time.Time, email string, phone string, citizenId string, address string) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO Employee(name, dob, email, phone, citizenId, address) VALUES($1, $2,$3,$4,$5,$6) RETURNING id", name, dob, email, phone, citizenId, address).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func CreateAccount(db *sql.DB, username string, password string, employeeId int, role int) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO Account(username, password,employeeId, role) VALUES($1, $2,$3,$4) RETURNING id", username, password, employeeId, role).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func ReadEmployee(db *sql.DB, id int) (Employee, error) {
	var e Employee
	row := db.QueryRow("SELECT id, name,dob, email,phone,citizenId,address FROM Employee WHERE id=$1", id)

	err := row.Scan(&e.Id, &e.Name, &e.Dob, &e.Email, &e.Phone, &e.CitizenId, &e.Address)
	if err != nil {
		return e, err
	}
	return e, nil
}
func ReadAllEmployee(db *sql.DB) ([]Employee, error) {
	rows, err := db.Query("SELECT id, name,dob, email,phone,citizenId,address FROM Employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var e Employee
		err := rows.Scan(&e.Id, &e.Name, &e.Dob, &e.Email, &e.Phone, &e.CitizenId, &e.Address)
		if err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, nil
}

func UpdateEmployee(db *sql.DB, id int, name string, dob time.Time, email string, phone string, citizenId string, address string) (int, error) {
	_, err := db.Exec("UPDATE Employee SET name=$1, dob=$2, email=$3, phone=$4, citizenId=$5, address=$6 WHERE id=$7", name, dob, email, phone, citizenId, address, id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func DeleteUser(db *sql.DB, id int) (int, error) {
	_, err1 := db.Exec("DELETE FROM ACCOUNT WHERE id=$1", id)
	if err1 != nil {
		return -1, err1
	}
	_, err := db.Exec("DELETE FROM EMPLOYEE WHERE id=$1", id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
func getUsers(c *gin.Context) {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
	employees, err := ReadAllEmployee(db)
	if err != nil {
		log.Fatal("Error reading database: ", err)
	}
	c.IndentedJSON(http.StatusOK, employees)
}
func getSpecificUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id does not valid"})
		return
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
	employee, err := ReadEmployee(db, id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "something wrong happen"})
		return
	}
	c.IndentedJSON(http.StatusOK, employee)
}
func postUsers(c *gin.Context) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
	var newEmployee Employee
	if err := c.BindJSON(&newEmployee); err != nil {
		return
	}
	i, err := CreateEmployee(db, newEmployee.Name, newEmployee.Dob, newEmployee.Email, newEmployee.Phone, newEmployee.CitizenId, newEmployee.Address)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Employee %v has been created\n", i)
	i1, err := CreateAccount(db, newEmployee.Email, newEmployee.CitizenId, i, 1)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Account %v has been created \n", i1)
	c.IndentedJSON(http.StatusCreated, newEmployee)
}

func updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id is not valid"})
		return
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
	var newEmployee Employee
	if err := c.BindJSON(&newEmployee); err != nil {
		return
	}
	i, err := UpdateEmployee(db, id, newEmployee.Name, newEmployee.Dob, newEmployee.Email, newEmployee.Phone, newEmployee.CitizenId, newEmployee.Address)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Employee %v has been updated\n", i)

	c.IndentedJSON(http.StatusCreated, newEmployee)
}

func delUsers(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id is not valid"})
		return
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
	i, err := DeleteUser(db, id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %v has been deleted", i)})
}

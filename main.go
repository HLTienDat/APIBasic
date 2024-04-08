package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// "html"
	// "net/http"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0125661"
	dbname   = "Users"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	// http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hi")
	// })

	// log.Fatal(http.ListenAndServe(":8081", nil))
	// Establish a connection to the PostgreSQL database
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
	ts := time.Now()

	idx, err := CreateEmployee(db, "Hoang Le Tien Dat", ts, "hltiendat2002@gmail.com", "0943211621", "274981446", "Nguyen Thi Minh Khai")
	if err != nil {
		log.Fatal("Error creating employee: ", err)
	}
	fmt.Println("Created employee with ID:", idx)
	idy, err := CreateAccount(db, "tiendathola", "matkhauDonGian", idx, 1)
	if err != nil {
		log.Fatal("Error creating account: ", err)
	}
	fmt.Println("Created account with ID:", idy)

	e, err := ReadEmployee(db, idx)
	if err != nil {
		log.Fatal("Error reading records: ", err)
	}
	fmt.Println("===Employee===")
	fmt.Println("ID:", e.id)
	fmt.Println("Name:", e.name)
	fmt.Println("Email:", e.email)
	fmt.Println("Address:", e.address)
	fmt.Println("Phone:", e.phone)

	err = UpdateEmployee(db, idx, "Hoang Le Tien Dinh", ts, "susieglitter@rhodes.com", "0943911616", "255", "Sparkle .JR Caster Rhodes")
	if err != nil {
		log.Fatal("Error updating employee: ", err)
	}
	fmt.Println("Updated employee with ID:", idx)
	e1, err := ReadEmployee(db, idx)
	if err != nil {
		log.Fatal("Error reading employee: ", err)
	}
	fmt.Println("Employee Ater Updating:")
	fmt.Println("===Employee===")
	fmt.Println("ID:", e1.id)
	fmt.Println("Name:", e1.name)
	fmt.Println("Email:", e1.email)
	fmt.Println("Address:", e1.address)
	fmt.Println("Phone:", e1.phone)

	err = DeleteRecord(db, 1)
	if err != nil {
		log.Fatal("Error deleting record: ", err)
	}
	fmt.Println("Deleted record with ID:", 1)
	records, err := ReadAllEmployee(db)
	if err != nil {
		log.Fatal("Error reading records: ", err)
	}
	fmt.Println("All records:")
	for _, e2 := range records {
		fmt.Println("===Employee===")
		fmt.Println("ID:", e2.id)
		fmt.Println("Name:", e2.name)
		fmt.Println("Email:", e2.email)
		fmt.Println("Address:", e2.address)
		fmt.Println("Phone:", e2.phone)
	}

}

type Account struct {
	id         int
	username   string
	password   string
	employeeId string
	role       int
}

type Employee struct {
	id        int
	name      string
	dob       time.Time
	email     string
	phone     string
	citizenId string
	address   string
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

	err := row.Scan(&e.id, &e.name, &e.dob, &e.email, &e.phone, &e.citizenId, &e.address)
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

	var records []Employee
	for rows.Next() {
		var e Employee
		err := rows.Scan(&e.id, &e.name, &e.dob, &e.email, &e.phone, &e.citizenId, &e.address)
		if err != nil {
			return nil, err
		}
		records = append(records, e)
	}
	return records, nil
}

func UpdateEmployee(db *sql.DB, id int, name string, dob time.Time, email string, phone string, citizenId string, address string) error {
	_, err := db.Exec("UPDATE Employee SET name=$1, dob=$2, email=$3, phone=$4, citizenId=$5, address=$6 WHERE id=$7", name, dob, email, phone, citizenId, address, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRecord(db *sql.DB, id int) error {
	_, err1 := db.Exec("DELETE FROM ACCOUNT WHERE id=$1", id)
	if err1 != nil {
		return err1
	}
	_, err := db.Exec("DELETE FROM EMPLOYEE WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	apimodel "test3/module/item/model"
	"time"

	_ "github.com/lib/pq"
)

func ReadEmployee(id int) (apimodel.Employee, error) {
	var e apimodel.Employee
	db, err := ConnectDB()
	if err != nil {
		return e, err
	}
	row := db.QueryRow("SELECT id, name,dob, email,phone,citizenId,address FROM Employee WHERE id=$1", id)
	err = row.Scan(&e.Id, &e.Name, &e.Dob, &e.Email, &e.Phone, &e.CitizenId, &e.Address)
	if err != nil {
		return e, err
	}
	return e, nil
}
func ReadAllEmployee() ([]apimodel.Employee, error) {
	var employees []apimodel.Employee
	db, err := ConnectDB()
	if err != nil {
		return employees, err
	}
	rows, err := db.Query("SELECT id, name,dob, email,phone,citizenId,address FROM Employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e apimodel.Employee
		err := rows.Scan(&e.Id, &e.Name, &e.Dob, &e.Email, &e.Phone, &e.CitizenId, &e.Address)
		if err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, nil
}

func CreateEmployee(name string, dob time.Time, email string, phone string, citizenId string, address string) (int, error) {
	var id int
	db, err := ConnectDB()
	if err != nil {
		return 0, err
	}
	err = db.QueryRow("INSERT INTO Employee(name, dob, email, phone, citizenId, address) VALUES($1, $2,$3,$4,$5,$6) RETURNING id", name, dob, email, phone, citizenId, address).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateEmployee(id int, name string, dob time.Time, email string, phone string, citizenId string, address string) (int, error) {
	db, err := ConnectDB()
	if err != nil {
		return 0, err
	}
	_, err = db.Exec("UPDATE Employee SET name=$1, dob=$2, email=$3, phone=$4, citizenId=$5, address=$6 WHERE id=$7", name, dob, email, phone, citizenId, address, id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func DeleteEmployee(id int) (int, error) {
	db, err := ConnectDB()
	if err != nil {
		return 0, err
	}
	_, err = db.Exec("DELETE FROM EMPLOYEE WHERE id=$1", id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

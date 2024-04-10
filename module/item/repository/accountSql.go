package repository

import (
	_ "github.com/lib/pq"
)

func FindUsername(username string) (int, error) {
	var id int
	db, err := ConnectDB()
	if err != nil {
		return -1, err
	}
	err = db.QueryRow("SELECT id  FROM Account WHERE username=$1", username).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func CreateAccount(username string, password string, employeeId int, role int) (int, error) {
	var id int
	db, err := ConnectDB()
	if err != nil {
		return 0, err
	}
	err = db.QueryRow("INSERT INTO Account(username, password,employeeId, role) VALUES($1, $2,$3,$4) RETURNING id", username, password, employeeId, role).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdatePassword(id int, password string) (bool, error) {
	db, err := ConnectDB()
	if err != nil {
		return false, err
	}
	_, err = db.Exec("UPDATE Account SET password=$2 WHERE id=$1", id, password)

	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteAccount(employeeId int) (int, error) {
	db, err := ConnectDB()
	if err != nil {
		return 0, err
	}
	_, err = db.Exec("DELETE FROM ACCOUNT WHERE employeeId=$1", employeeId)
	if err != nil {
		return -1, err
	}
	return employeeId, nil
}

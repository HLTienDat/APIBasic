package business

import (
	"database/sql"
	"errors"
	"fmt"
	repository "test3/module/item/repository"
	"time"
)

func WriteNewAccount(username, password string, role int) (int, error) {
	id, err := repository.FindUsername(username)
	if err != nil {
		fmt.Println(err)
		if err != sql.ErrNoRows {
			return -1, err
		}
	}
	if id != -1 {
		return -1, errors.New("username existed")
	}
	var time time.Time
	employeeId, err := WriteNewEmployee("", "", "", "", "", time)
	if err != nil {
		return -1, err
	}
	id, err = repository.CreateAccount(username, password, employeeId, role)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func WritePassword(username, password string) (int, error) {
	id, err := repository.FindUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, errors.New("not found")
		}
		return -1, err
	}
	_, err = repository.UpdatePassword(id, password)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func RemoveAccount(id int) (int, error) {
	_, err := repository.DeleteAccount(id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

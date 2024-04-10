package business

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	model "test3/module/item/model"
	repository "test3/module/item/repository"
)

func FetchUsers() ([]model.Employee, error) {
	employees, err := repository.ReadAllEmployee()
	if err != nil {
		return []model.Employee{}, err
	}
	return employees, nil
}

func FetchSpecificUser(c *gin.Context) (model.Employee, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.Employee{}, err
	}
	employee, err := repository.ReadEmployee(id)
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

func WriteNewEmployee(name, email, phone, citizenId, address string, dob time.Time) (int, error) {
	id, err := repository.CreateEmployee(name, dob, email, phone, citizenId, address)
	if err != nil {
		return -1, err
	}
	return id, nil
}
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
		return -1, err
	}
	_, err = repository.UpdatePassword(id, password)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func WriteEmployee(id int, name string, dob time.Time, email string, phone string, citizenId string, address string) (int, error) {
	_, err := repository.UpdateEmployee(id, name, dob, email, phone, citizenId, address)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func RemoveEmployee(id int) (int, error) {
	_, err := RemoveAccount(id)
	if err != nil {
		return -1, err
	}
	_, err = repository.DeleteEmployee(id)
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

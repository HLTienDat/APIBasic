package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	business "test3/module/item/business"
	model "test3/module/item/model"
)

func Transfer(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetUsers(c *gin.Context) {
	employees, err := business.FetchUsers()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot get Employees: %s", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, employees)
}
func GetSpecificUser(c *gin.Context) {
	employee, err := business.FetchSpecificUser(c)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Employee not found"})
			return
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, employee)
}
func PostUsers(c *gin.Context) {
	var newAccount model.Account
	if err := c.BindJSON(&newAccount); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot search for username: %s", err)})
		return
	}
	i, err := business.WriteNewAccount(newAccount.Username, newAccount.Password, newAccount.Role)
	if err != nil {
		if err.Error() == "username existed" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Username already registered"})
			return
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot create new account: %s", err)})
		return
	}
	fmt.Printf("Employee %v has been created\n", i)
	c.IndentedJSON(http.StatusCreated, newAccount)
}

func UpdateUser(c *gin.Context) {
	id, err := Transfer(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot read id: %s", err)})
		return
	}
	var newEmployee model.Employee
	if err := c.BindJSON(&newEmployee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot read changes: %s", err)})
		return
	}
	i, err := business.WriteEmployee(id, newEmployee.Name, newEmployee.Dob, newEmployee.Email, newEmployee.Phone, newEmployee.CitizenId, newEmployee.Address)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot update new Employee: %s", err)})
		return
	}
	fmt.Printf("Employee %v has been updated\n", i)

	c.IndentedJSON(http.StatusCreated, newEmployee)
}

func DelUsers(c *gin.Context) {
	id, err := Transfer(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("ID is invalid: %s", err)})
		return
	}
	i, err := business.RemoveEmployee(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot delete Employee: %s", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %v has been deleted", i)})
}

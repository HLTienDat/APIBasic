package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	business "test3/module/item/business"
	model "test3/module/item/model"
)

func PostPassword(c *gin.Context) {
	var newAccount model.Account
	if err := c.BindJSON(&newAccount); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot search for username: %s", err)})
		return
	}
	i, err := business.WritePassword(newAccount.Username, newAccount.Password)
	if err != nil {
		if err.Error() == "not found" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Cannot find username"})
			return
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot change password: %s", err)})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Change password of account: %v", i)})
}

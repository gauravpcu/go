package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"supplier-portal/models"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "All Good"})
}

func HomeError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Nothing Good"})
}

func GetGoUsers(c *gin.Context) {

	var usersFromDb []models.GoUsers
	limit, nil := strconv.Atoi(c.Request.URL.Query().Get("limit"))
	if err := models.DB.Limit(limit).Find(&usersFromDb).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, usersFromDb)
}

func GoUsersAndManipulateThem(c *gin.Context) {

	var usersFromDb []models.GoUsers
	limit, nil := strconv.Atoi(c.Request.URL.Query().Get("limit"))
	if err := models.DB.Limit(limit).Find(&usersFromDb).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	for i := 0; i < len(usersFromDb); i++ {
		usersFromDb[i].Age += 1
		usersFromDb[i].FirstName = strings.ToUpper(usersFromDb[i].FirstName)
		usersFromDb[i].LastName = "," + strings.ToUpper(usersFromDb[i].LastName)
	}
	c.JSON(http.StatusOK, usersFromDb)

}

func CreateGoUsers(c *gin.Context) {

	var input []models.GoUsers
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &input)

	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	models.DB.Create(input)
	c.IndentedJSON(http.StatusCreated, input)
}

func UpdateGoUsers(c *gin.Context) {

	var input []models.GoUsers
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &input)

	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	for _, updateInput := range input {
		models.DB.Model(&updateInput).Updates(updateInput)
	}

	c.JSON(http.StatusOK, input)
}

/*
curl http://localhost:8080/api/user --include --header "Content-Type: application/json" --request "PUT" --data '{, "first_name": "Gaurav","last_name": "Singh","email": "gaurav.kumar@procurementpartners.com","password": "Password!", "Age": 41}'

curl http://localhost:8080/api/user --include --header "Content-Type: application/json" --request "POST" --data '[{
    	"first_name": "Test",
        "last_name": "one",
        "email": "test1@example.com",
        "password": "123456",
        "age": 5
    },
    {
        "first_name": "Test",
        "last_name": "two",
        "email": "test2@example.com",
        "password": "123456",
        "age": 6
    }]'

	curl http://localhost:8080/api/user --include --header "Content-Type: application/json" --request "PUT" --data '[{
    	"id": 599,
		"first_name": "Test",
        "last_name": "one",
        "email": "test1@example.com",
        "password": "123456",
        "age": 5
    },
    {
		"id": 600,
        "first_name": "Test",
        "last_name": "two",
        "email": "test2@example.com",
        "password": "123456",
        "age": 6
    }]'
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"supplier-portal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetURL(path string) string {
	return fmt.Sprintf("http://localhost:8080%s", path)
}
func TestGetProductHandler(t *testing.T) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", GetURL("/products"), nil)
	res, _ := client.Do(req)
	var products []models.Product
	assert.Equal(t, http.StatusOK, res.StatusCode)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &products)
	assert.NotEmpty(t, products)
}


/*
func TestGetProductHandler(t *testing.T) {
	r := SetUpRouter()

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var products []models.Product
	json.Unmarshal(w.Body.Bytes(), &products)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, products)
}



func TestNewCompanyHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/company", NewCompanyHandler)
	companyId := xid.New().String()
	company := Company{
		ID:      companyId,
		Name:    "Demo Company",
		CEO:     "Demo CEO",
		Revenue: "35 million",
	}
	jsonValue, _ := json.Marshal(company)
	req, _ := http.NewRequest("POST", "/company", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}*/

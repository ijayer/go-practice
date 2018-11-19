package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"go-web/api"
)

var app *api.App

// Test main
func TestMain(m *testing.M) {
	app = api.NewApp()
	app.InitMgo(api.DBName, api.DBAddr)
	code := m.Run()
	app.Clean(api.DBName)
	os.Exit(code)
}

// Test clean up database
func TestApp_Clean(t *testing.T) {
	app.Clean(api.DBName)

	// http request
	req, _ := http.NewRequest(http.MethodGet, "/v0/products", nil)
	resp := executeRequest(req)

	// check response code
	checkResponseCode(t, http.StatusOK, resp.Code)

	// check respond data
	var results []api.Product
	json.Unmarshal(resp.Body.Bytes(), results)
	if len(results) != 0 {
		t.Errorf("excepetd an empty array, got '%v'", results)
	}
}

// Test find a not exist product
func TestGotNonExistentProduct(t *testing.T) {
	app.Clean(api.DBName)

	req, _ := http.NewRequest(http.MethodGet, "/v0/products/201711190252", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, resp.Code)

	var m map[string]string
	json.Unmarshal(resp.Body.Bytes(), &m)
	if m["error"] != "not found" {
		t.Errorf("excepted the `error` key of the response "+
			"to be set to `product not found`, got '%v'", m["error"])
	}
}

// Test add a product to the database
func TestApp_Create(t *testing.T) {
	app.Clean(api.DBName)

	payload := []byte(`{"name": "product", "price":11.22}`)
	req, _ := http.NewRequest(http.MethodPost, "/v0/products", bytes.NewReader(payload))
	resp := executeRequest(req)

	if resp.Code == http.StatusBadRequest {
		t.Errorf("excepted respond code '%v', got %v\n", http.StatusCreated, http.StatusBadRequest)
	}
	if resp.Code == http.StatusInternalServerError {
		t.Errorf("excepted respond code '%v', got %v\n", http.StatusCreated, http.StatusInternalServerError)
	}
	checkResponseCode(t, http.StatusCreated, resp.Code)

	var m map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &m)
	data := m["data"].(map[string]interface{})
	if data["name"] != "product" {
		t.Errorf("excepted product name is 'product', got '%v'", m["name"])
	}
	if data["price"] != 11.22 {
		t.Errorf("excepted product price is '11.22', got '%v'", m["price"])
	}
	if data["pid"] != 0.0 { // here, the id of respond is 0.0, not 0
		t.Errorf("excepted product id is '0', got '%v'", m["pid"])
	}
}

// Test create a product
func TestApp_Create2(t *testing.T) {
	// app.Clean(DBName)
	payload := []byte(`"name": "product", "price":11.22`)
	req, _ := http.NewRequest(http.MethodPost, "/v0/products", bytes.NewReader(payload))
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, resp.Code)

	var m map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &m)
	if m["error"] != "invalid request payload" {
		t.Errorf("except error: 'invalid request payload', got %v\n", resp.Body.String())
	}
}

// Test find a product
func TestApp_FindOne(t *testing.T) {
	// app.Clean(DBName)
	// addProducts(1)

	req, _ := http.NewRequest("GET", "/v0/products/0", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

// Test update a product
func TestApp_Update(t *testing.T) {
	// app.Clean(DBName)
	// addProducts(1)

	// obtain original product data
	req, _ := http.NewRequest("GET", "/v0/products/0", nil)
	resp := executeRequest(req)
	var originalProduct map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &originalProduct)
	originalProduct = originalProduct["data"].(map[string]interface{})

	// update the product data
	payload := []byte(`{"name":"product-u0","price":22.22}`)
	req, _ = http.NewRequest("PUT", "/v0/products/0", bytes.NewBuffer(payload))
	resp = executeRequest(req)

	checkResponseCode(t, http.StatusOK, resp.Code)

	var m map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &m)
	m = m["data"].(map[string]interface{})
	if m["pid"] != originalProduct["pid"] {
		t.Errorf("expected the id to remain the same '%v'. "+
			"Got %v", originalProduct["pid"], m["pid"])
	}
	if m["name"] == originalProduct["name"] {
		t.Errorf("expected the name to change from '%v' to '%v'. "+
			"Got '%v'", originalProduct["name"], m["name"], m["name"])
	}
	if m["price"] == originalProduct["price"] {
		t.Errorf("expected the price to change from '%v' to '%v'. "+
			"Got '%v'", originalProduct["price"], m["price"], m["price"])
	}
}

// Test delete a product
func TestApp_Delete(t *testing.T) {
	// app.Clean(DBName)
	// addProducts(1)

	req, _ := http.NewRequest("GET", "/v0/products/0", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/v0/products/0", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/v0/products/0", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

// Execute a request
func executeRequest(request *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, request)
	return rr
}

// Check response code
func checkResponseCode(t *testing.T, want, usual int) {
	if want != usual {
		t.Errorf("excepted response code '%v', got '%v'", want, usual)
	}
}

// Add products to the database
func addProducts(num int) {
	if num < 1 {
		num = 1
	}

	session := app.Session.Copy()
	defer session.Close()
	co := session.DB(api.DBName).C("products")
	for i := 0; i < num; i++ {
		co.Insert(&api.Product{PID: i, Name: "product", Price: (float64(i) + 1.0) * 10})
	}
}

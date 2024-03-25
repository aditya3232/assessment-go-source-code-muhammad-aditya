package test

import (
	"assessment-go-source-code-muhammad-aditya/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	ClearAll()
	requestBody := model.CreateCustomerRequest{
		NationalId:    18071999,
		Name:          "Muhammad Aditya",
		DetailAddress: "jl.kebenaran jawa timur, malang",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/customers/", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request) // tes menggunakan fiber
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CustomerResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.NationalId, responseBody.Data.NationalId)
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
	assert.Equal(t, requestBody.DetailAddress, responseBody.Data.DetailAddress)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
	assert.NotNil(t, responseBody.Data.ID)
}

func TestCreateCustomerFailed(t *testing.T) {
	ClearAll()

	requestBody := model.CreateCustomerRequest{
		NationalId: 18071999,
		// Name:          "Muhammad Aditya",
		DetailAddress: "jl.kebenaran jawa timur, malang",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/customers/", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CustomerResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestListCustomer(t *testing.T) {
	ClearAll()
	CreateCustomers(t, 5)

	request := httptest.NewRequest(http.MethodGet, "/api/customers/", nil)
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.CustomerResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 5, len(responseBody.Data))
}

func TestListCustomerFailed(t *testing.T) {
	ClearAll()

	request := httptest.NewRequest(http.MethodGet, "/api/customer/", nil)
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.CustomerResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestGetCustomer(t *testing.T) {
	ClearAll()

	CreateCustomers(t, 5)
	customer := GetFirstCustomer(t)

	request := httptest.NewRequest(http.MethodGet, "/api/customers/"+customer.ID, nil)
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CustomerResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, customer.ID, responseBody.Data.ID)
	assert.Equal(t, customer.NationalId, responseBody.Data.NationalId)
	assert.Equal(t, customer.Name, responseBody.Data.Name)
	assert.Equal(t, customer.DetailAddress, responseBody.Data.DetailAddress)
	assert.Equal(t, customer.CreatedAt, responseBody.Data.CreatedAt)
	assert.Equal(t, customer.UpdatedAt, responseBody.Data.UpdatedAt)
}

func TestGetCustomerFailed(t *testing.T) {
	ClearAll()

	CreateCustomers(t, 5)
	uuid := "e8cc8ec6-5a9c-4cde-9e52-b435d15bb936"

	request := httptest.NewRequest(http.MethodGet, "/api/customers/"+uuid, nil)
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CustomerResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestUpdateCustomer(t *testing.T) {
	ClearAll()

	CreateCustomers(t, 5)
	customer := GetFirstCustomer(t)

	requestBody := model.UpdateCustomerRequest{
		ID:            customer.ID,
		Name:          "Ichsan Ashiddiqi",
		DetailAddress: "jl.lurus jawa timur, surabaya",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/customers/"+customer.ID, strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CustomerResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
	assert.Equal(t, requestBody.DetailAddress, responseBody.Data.DetailAddress)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
	assert.NotNil(t, responseBody.Data.ID)
}

func TestUpdateCustomerFailed(t *testing.T) {
	ClearAll()

	CreateCustomers(t, 5)
	uuid := "e8cc8ec6-5a9c-4cde-9e52-b435d15bb936"

	requestBody := model.UpdateCustomerRequest{
		ID:            uuid,
		Name:          "Ichsan Ashiddiqi",
		DetailAddress: "jl.lurus jawa timur, surabaya",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/customers/"+uuid, strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestDeleteCustomer(t *testing.T) {
	ClearAll()

	CreateCustomers(t, 5)
	customer := GetFirstCustomer(t)

	request := httptest.NewRequest(http.MethodDelete, "/api/customers/"+customer.ID, nil)
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestDeleteCustomerFailed(t *testing.T) {
	ClearAll()

	CreateCustomers(t, 5)
	uuid := "e8cc8ec6-5a9c-4cde-9e52-b435d15bb936"

	request := httptest.NewRequest(http.MethodDelete, "/api/customers/"+uuid, nil)
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestCreateCustomerBigData(t *testing.T) {
	ClearAll()

	start := time.Now()

	for i := 0; i < 30; i++ {
		requestBody := model.CreateCustomerRequest{
			NationalId:    int64(18071999 + i),
			Name:          fmt.Sprintf("Muhammad Aditya %d", i),
			DetailAddress: fmt.Sprintf("jl.kebenaran jawa timur, malang %d", i),
		}

		bodyJson, err := json.Marshal(requestBody)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/api/customers/", strings.NewReader(string(bodyJson)))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		defer response.Body.Close()

		assert.Equal(t, http.StatusOK, response.StatusCode)

		bytes, err := io.ReadAll(response.Body)
		assert.Nil(t, err)

		responseBody := new(model.WebResponse[model.CustomerResponse])
		err = json.Unmarshal(bytes, responseBody)
		assert.Nil(t, err)

	}

	// Assert execution time is within a reasonable limit
	elapsed := time.Since(start)
	assert.True(t, elapsed < 10*time.Second, "execution time should be less than 5 seconds for 30 records")
}

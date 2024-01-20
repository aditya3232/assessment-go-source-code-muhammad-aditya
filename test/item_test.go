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

func TestCreateItem(t *testing.T) {
	ClearItem()
	requestBody := model.CreateItemRequest{
		ItemCode: 121314,
		ItemName: "printer",
		Type:     "hardware",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/items/", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ItemResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.ItemCode, responseBody.Data.ItemCode)
	assert.Equal(t, requestBody.ItemName, responseBody.Data.ItemName)
	assert.Equal(t, requestBody.Type, responseBody.Data.Type)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
	assert.NotNil(t, responseBody.Data.ID)
}

func TestCreateItemFailed(t *testing.T) {
	ClearItem()
	requestBody := model.CreateItemRequest{
		ItemCode: 121314,
		// ItemName: "printer",
		Type: "hardware",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/items/", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ItemResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestListItem(t *testing.T) {
	ClearItem()
	CreateItems(t, 5)

	request := httptest.NewRequest(http.MethodGet, "/api/items/", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.ItemResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 5, len(responseBody.Data))
}

func TestListItemFailed(t *testing.T) {
	ClearItem()

	request := httptest.NewRequest(http.MethodGet, "/api/item/", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.ItemResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestGetItem(t *testing.T) {
	ClearItem()
	CreateItems(t, 5)
	item := GetFirstItem(t)

	request := httptest.NewRequest(http.MethodGet, "/api/items/"+item.ID, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ItemResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, item.ItemCode, responseBody.Data.ItemCode)
	assert.Equal(t, item.ItemName, responseBody.Data.ItemName)
	assert.Equal(t, item.Type, responseBody.Data.Type)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
	assert.NotNil(t, responseBody.Data.ID)
}

func TestGetItemFailed(t *testing.T) {
	ClearItem()
	CreateItems(t, 5)
	uuid := "e8cc8ec6-5a9c-4cde-9e52-b435d15bb936"

	request := httptest.NewRequest(http.MethodGet, "/api/items/"+uuid, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ItemResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestUpdateItem(t *testing.T) {
	ClearItem()
	CreateItems(t, 5)
	item := GetFirstItem(t)

	requestBody := model.UpdateItemRequest{
		ID:       item.ID,
		ItemName: "printer",
		Type:     "hardware",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/items/"+item.ID, strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ItemResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.ItemName, responseBody.Data.ItemName)
	assert.Equal(t, requestBody.Type, responseBody.Data.Type)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
	assert.NotNil(t, responseBody.Data.ID)
}

func TestUpdateItemFailed(t *testing.T) {
	ClearItem()
	CreateItems(t, 5)
	item := GetFirstItem(t)

	requestBody := model.UpdateItemRequest{
		ID: item.ID,
		// ItemName: "printer",
		Type: "hardware",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/items/"+item.ID, strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ItemResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestDeleteItem(t *testing.T) {
	ClearItem()
	CreateItems(t, 5)
	item := GetFirstItem(t)

	request := httptest.NewRequest(http.MethodDelete, "/api/items/"+item.ID, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestDeleteItemFailed(t *testing.T) {
	ClearItem()
	CreateItems(t, 5)
	uuid := "e8cc8ec6-5a9c-4cde-9e52-b435d15bb936"

	request := httptest.NewRequest(http.MethodDelete, "/api/items/"+uuid, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestCreateItemBigData(t *testing.T) {
	ClearItem()

	start := time.Now()

	for i := 0; i < 30; i++ {
		requestBody := model.CreateItemRequest{
			ItemCode: int64(18071999 + i),
			ItemName: fmt.Sprintf("printer %d", i),
			Type:     fmt.Sprintf("hardware %d", i),
		}

		bodyJson, err := json.Marshal(requestBody)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/api/items/", strings.NewReader(string(bodyJson)))
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

package test

import (
	"assessment-go-source-code-muhammad-aditya/internal/model"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInvoice(t *testing.T) {
	// ClearInvoiceItem()
	// ClearInvoice()
	// ClearCustomer()
	// ClearItem()
	ClearAll()

	CreateCustomers(t, 1)
	customer := GetFirstCustomer(t)

	CreateItems(t, 5)
	item := GetAllItem(t)

	requestBody := model.CreateInvoiceRequest{
		InvoiceNumber: "0002",
		CustomerId:    customer.ID,
		Subject:       "winter marketing campaign",
		IssuedDate:    1705561215583,
		DueDate:       1705560536496,
		TotalItem:     3,
		SubTotal:      1400,
		GrandTotal:    1540,
		Status:        "padi",
		InvoiceItems: []model.CreateInvoiceItemRequest{
			{
				ItemId:       item[0].ID,
				ItemPrice:    250,
				ItemQuantity: 2,
				Amount:       500,
			},
			{
				ItemId:       item[1].ID,
				ItemPrice:    150,
				ItemQuantity: 2,
				Amount:       300,
			},
			{
				ItemId:       item[2].ID,
				ItemPrice:    300,
				ItemQuantity: 2,
				Amount:       600,
			},
		},
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/invoices/", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.InvoiceResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.InvoiceNumber, responseBody.Data.InvoiceNumber)
	assert.Equal(t, requestBody.CustomerId, responseBody.Data.CustomerId)
	assert.Equal(t, requestBody.Subject, responseBody.Data.Subject)
	assert.Equal(t, requestBody.IssuedDate, responseBody.Data.IssuedDate)
	assert.Equal(t, requestBody.DueDate, responseBody.Data.DueDate)
	assert.Equal(t, requestBody.TotalItem, responseBody.Data.TotalItem)
	assert.Equal(t, requestBody.SubTotal, responseBody.Data.SubTotal)
	assert.Equal(t, requestBody.GrandTotal, responseBody.Data.GrandTotal)
	assert.Equal(t, requestBody.Status, responseBody.Data.Status)
	for i, v := range requestBody.InvoiceItems {
		assert.Equal(t, v.ItemId, responseBody.Data.InvoiceItems[i].ItemId)
		assert.Equal(t, v.ItemPrice, responseBody.Data.InvoiceItems[i].ItemPrice)
		assert.Equal(t, v.ItemQuantity, responseBody.Data.InvoiceItems[i].ItemQuantity)
		assert.Equal(t, v.Amount, responseBody.Data.InvoiceItems[i].Amount)
	}
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
	assert.NotNil(t, responseBody.Data.ID)
}

func TestCreateInvoiceFailed(t *testing.T) {
	ClearAll()

	CreateCustomers(t, 1)
	customer := GetFirstCustomer(t)

	CreateItems(t, 5)
	item := GetAllItem(t)

	requestBody := model.CreateInvoiceRequest{
		// InvoiceNumber: "0002",
		CustomerId: customer.ID,
		Subject:    "winter marketing campaign",
		IssuedDate: 1705561215583,
		DueDate:    1705560536496,
		TotalItem:  3,
		SubTotal:   1400,
		GrandTotal: 1540,
		Status:     "padi",
		InvoiceItems: []model.CreateInvoiceItemRequest{
			{
				ItemId:       item[0].ID,
				ItemPrice:    250,
				ItemQuantity: 2,
				Amount:       500,
			},
			{
				ItemId:       item[1].ID,
				ItemPrice:    150,
				ItemQuantity: 2,
				Amount:       300,
			},
			{
				ItemId:       item[2].ID,
				ItemPrice:    300,
				ItemQuantity: 2,
				Amount:       600,
			},
		},
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/invoices/", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.InvoiceResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestListInvoice(t *testing.T) {
	ClearAll()

	CreateCustomers(t, 5)
	customer := GetFirstCustomer(t)

	CreateInvoices(t, customer, 5)

	request := httptest.NewRequest(http.MethodGet, "/api/invoices/", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.InvoiceListResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 5, len(responseBody.Data))
}

func TestListInvoiceFailed(t *testing.T) {
	ClearAll()

	CreateCustomers(t, 5)
	customer := GetFirstCustomer(t)

	CreateInvoices(t, customer, 5)

	request := httptest.NewRequest(http.MethodGet, "/api/invoice/", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.InvoiceListResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestGetInvoice(t *testing.T) {
	ClearAll()

	CreateCustomers(t, 5)
	customer := GetFirstCustomer(t)

	CreateInvoices(t, customer, 5)
	invoice := GetFirstInvoice(t)

	request := httptest.NewRequest(http.MethodGet, "/api/invoices/"+invoice.ID, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.InvoiceResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, invoice.ID, responseBody.Data.ID)
	assert.Equal(t, invoice.InvoiceNumber, responseBody.Data.InvoiceNumber)
	assert.Equal(t, invoice.CustomerId, responseBody.Data.CustomerId)
	assert.Equal(t, invoice.Subject, responseBody.Data.Subject)
	assert.Equal(t, invoice.IssuedDate, responseBody.Data.IssuedDate)
	assert.Equal(t, invoice.DueDate, responseBody.Data.DueDate)
	assert.Equal(t, invoice.TotalItem, responseBody.Data.TotalItem)
	assert.Equal(t, invoice.SubTotal, responseBody.Data.SubTotal)
	assert.Equal(t, invoice.GrandTotal, responseBody.Data.GrandTotal)
	assert.Equal(t, invoice.Status, responseBody.Data.Status)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
	assert.NotNil(t, responseBody.Data.ID)
}

func TestGetInvoiceFailed(t *testing.T) {
	ClearAll()

	CreateCustomers(t, 5)
	customer := GetFirstCustomer(t)

	CreateInvoices(t, customer, 5)
	uuid := "e8cc8ec6-5a9c-4cde-9e52-b435d15bb936"

	request := httptest.NewRequest(http.MethodGet, "/api/invoices/"+uuid, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.InvoiceResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestUpdateInvoice(t *testing.T) {
	ClearAll()

	CreateCustomers(t, 5)
	customer := GetFirstCustomer(t)

	CreateInvoices(t, customer, 5)
	invoice := GetFirstInvoice(t)

	requestBody := model.UpdateInvoiceRequest{
		ID:           invoice.ID,
		CustomerId:   customer.ID,
		Subject:      "winter marketing campaign",
		IssuedDate:   1705561215583,
		DueDate:      1705560536496,
		TotalItem:    3,
		SubTotal:     1400,
		GrandTotal:   1540,
		Status:       "paid",
		InvoiceItems: []model.CreateInvoiceItemRequest{},
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/invoices/"+invoice.ID, strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.InvoiceResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.CustomerId, responseBody.Data.CustomerId)
	assert.Equal(t, requestBody.Subject, responseBody.Data.Subject)
	assert.Equal(t, requestBody.IssuedDate, responseBody.Data.IssuedDate)
	assert.Equal(t, requestBody.DueDate, responseBody.Data.DueDate)
	assert.Equal(t, requestBody.TotalItem, responseBody.Data.TotalItem)
	assert.Equal(t, requestBody.SubTotal, responseBody.Data.SubTotal)
	assert.Equal(t, requestBody.GrandTotal, responseBody.Data.GrandTotal)
	assert.Equal(t, requestBody.Status, responseBody.Data.Status)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
	assert.NotNil(t, responseBody.Data.ID)
}

func TestUpdateInvoiceFailed(t *testing.T) {
	ClearAll()

	CreateCustomers(t, 5)
	customer := GetFirstCustomer(t)

	CreateInvoices(t, customer, 5)
	invoice := GetFirstInvoice(t)

	requestBody := model.UpdateInvoiceRequest{
		ID:         invoice.ID,
		CustomerId: customer.ID,
		// Subject:      "winter marketing campaign",
		IssuedDate:   1705561215583,
		DueDate:      1705560536496,
		TotalItem:    3,
		SubTotal:     1400,
		GrandTotal:   1540,
		Status:       "paid",
		InvoiceItems: []model.CreateInvoiceItemRequest{},
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/invoices/"+invoice.ID, strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.InvoiceResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

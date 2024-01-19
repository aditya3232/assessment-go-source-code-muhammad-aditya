package usecase

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"
	"assessment-go-source-code-muhammad-aditya/internal/model"
	"assessment-go-source-code-muhammad-aditya/internal/model/converter"
	"assessment-go-source-code-muhammad-aditya/internal/repository"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InvoiceUseCase struct {
	DB                    *gorm.DB
	Log                   *logrus.Logger
	Validate              *validator.Validate
	InvoiceRepository     *repository.InvoiceRepository
	CustomerRepository    *repository.CustomerRepository
	InvoiceItemRepository *repository.InvoiceItemRepository
}

func NewInvoiceUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	invoiceRepository *repository.InvoiceRepository, customerRepository *repository.CustomerRepository, invoiceItemRepository *repository.InvoiceItemRepository) *InvoiceUseCase {
	return &InvoiceUseCase{
		DB:                    db,
		Log:                   logger,
		Validate:              validate,
		InvoiceRepository:     invoiceRepository,
		CustomerRepository:    customerRepository,
		InvoiceItemRepository: invoiceItemRepository,
	}
}

func (c *InvoiceUseCase) Create(ctx context.Context, request *model.CreateInvoiceRequest) (*model.InvoiceResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, customer, request.CustomerId); err != nil {
		c.Log.WithError(err).Error("failed to find customer")
		return nil, fiber.ErrNotFound
	}

	invoiceId := uuid.New().String()
	invoiceItems := make([]entity.InvoiceItem, len(request.InvoiceItems))

	for i, itemRequest := range request.InvoiceItems {
		itemId := uuid.New().String()
		invoiceItems[i] = entity.InvoiceItem{
			ID:           itemId,
			InvoiceId:    invoiceId,
			ItemId:       itemRequest.ItemId,
			ItemPrice:    itemRequest.ItemPrice,
			ItemQuantity: itemRequest.ItemQuantity,
			Amount:       itemRequest.Amount,
		}
	}

	invoice := &entity.Invoice{
		ID:            invoiceId,
		InvoiceNumber: request.InvoiceNumber,
		CustomerId:    customer.ID,
		Subject:       request.Subject,
		IssuedDate:    request.IssuedDate,
		DueDate:       request.DueDate,
		TotalItem:     request.TotalItem,
		SubTotal:      request.SubTotal,
		GrandTotal:    request.GrandTotal,
		Status:        request.Status,
		Customer:      *customer,
		InvoiceItems:  invoiceItems,
	}

	totalInvoiceNumber, err := c.InvoiceRepository.CountByInvoiceNumber(tx, invoice)
	if err != nil {
		c.Log.Warnf("Failed count invoice from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if totalInvoiceNumber > 0 {
		c.Log.Warnf("Invoice already exists : %+v", err)
		return nil, fiber.ErrConflict
	}

	if err := c.InvoiceRepository.Create(tx, invoice); err != nil {
		c.Log.WithError(err).Error("error creating invoice")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating invoice")
		return nil, fiber.ErrInternalServerError
	}

	return converter.InvoiceToResponse(invoice), nil
}

func (c *InvoiceUseCase) Get(ctx context.Context, request *model.GetInvoiceRequest) (*model.InvoiceResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	invoice := new(entity.Invoice)
	if err := c.InvoiceRepository.FindById(tx, invoice, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting invoice")
		return nil, fiber.ErrNotFound
	}

	customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, customer, invoice.CustomerId); err != nil {
		c.Log.WithError(err).Error("failed to find customer")
		return nil, fiber.ErrNotFound
	}

	invoiceItems := make([]entity.InvoiceItem, 0)
	if err := c.InvoiceItemRepository.FindByInvoiceId(tx, &invoiceItems, invoice.ID); err != nil {
		c.Log.WithError(err).Error("failed to find invoice items")
		return nil, fiber.ErrNotFound
	}

	invoice.Customer = *customer
	invoice.InvoiceItems = invoiceItems

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting invoice")
		return nil, fiber.ErrInternalServerError
	}

	return converter.InvoiceToResponse(invoice), nil
}

func (c *InvoiceUseCase) Search(ctx context.Context, request *model.SearchInvoiceRequest) ([]model.InvoiceListResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, 0, fiber.ErrBadRequest
	}

	invoices, total, err := c.InvoiceRepository.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Error("error getting invoices")
		return nil, 0, fiber.ErrInternalServerError
	}

	customerIds := make([]string, len(invoices))
	for i, invoice := range invoices {
		customerIds[i] = invoice.CustomerId
	}

	responses := make([]model.InvoiceListResponse, len(invoices))
	for i, invoice := range invoices {
		customer := new(entity.Customer)

		if err := c.CustomerRepository.FindById(tx, customer, customerIds[i]); err != nil {
			c.Log.WithError(err).Error("error getting customer name")
			return []model.InvoiceListResponse{}, 0, fiber.ErrNotFound
		}

		response := converter.InvoiceListToResponse(&invoice)
		response.CustomerName = customer.Name
		responses[i] = *response
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting invoices")
		return nil, 0, fiber.ErrInternalServerError
	}

	return responses, total, nil
}

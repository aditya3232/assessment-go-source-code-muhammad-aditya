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

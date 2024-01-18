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

type CustomerUseCase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	CustomerRepository *repository.CustomerRepository
}

func NewCustomerUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	customerRepository *repository.CustomerRepository) *CustomerUseCase {
	return &CustomerUseCase{
		DB:                 db,
		Log:                logger,
		Validate:           validate,
		CustomerRepository: customerRepository,
	}
}

func (c *CustomerUseCase) Create(ctx context.Context, request *model.CreateCustomerRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	customer := &entity.Customer{
		ID:            uuid.New().String(),
		Name:          request.Name,
		DetailAddress: request.DetailAddress,
	}

	if err := c.CustomerRepository.Create(tx, customer); err != nil {
		c.Log.WithError(err).Error("error creating customer")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating customer")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CustomerToResponse(customer), nil
}

func (c *CustomerUseCase) Update(ctx context.Context, request *model.UpdateCustomerRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, customer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting customer")
		return nil, fiber.ErrNotFound
	}

	customer.Name = request.Name
	customer.DetailAddress = request.DetailAddress

	if err := c.CustomerRepository.Update(tx, customer); err != nil {
		c.Log.WithError(err).Error("error updating customer")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating customer")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CustomerToResponse(customer), nil
}

func (c *CustomerUseCase) Get(ctx context.Context, request *model.GetCustomerRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, customer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting customer")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting customer")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CustomerToResponse(customer), nil
}

func (c *CustomerUseCase) Delete(ctx context.Context, request *model.DeleteCustomerRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return fiber.ErrBadRequest
	}

	customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, customer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting customer")
		return fiber.ErrNotFound
	}

	if err := c.CustomerRepository.Delete(tx, customer); err != nil {
		c.Log.WithError(err).Error("error deleting customer")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting customer")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *CustomerUseCase) Search(ctx context.Context, request *model.SearchCustomerRequest) ([]model.CustomerResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, 0, fiber.ErrBadRequest
	}

	customers, total, err := c.CustomerRepository.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Error("error getting customers")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting customers")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.CustomerResponse, len(customers))
	for i, customer := range customers {
		responses[i] = *converter.CustomerToResponse(&customer)
	}

	return responses, total, nil
}

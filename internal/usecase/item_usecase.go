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

type ItemUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	ItemRepository *repository.ItemRepository
}

func NewItemUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	itemRepository *repository.ItemRepository) *ItemUseCase {
	return &ItemUseCase{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		ItemRepository: itemRepository,
	}
}

func (c *ItemUseCase) Create(ctx context.Context, request *model.CreateItemRequest) (*model.ItemResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	item := &entity.Item{
		ID:       uuid.New().String(),
		ItemCode: request.ItemCode,
		ItemName: request.ItemName,
		Type:     request.Type,
	}

	totalItemCode, err := c.ItemRepository.CountByItemCode(tx, item)
	if err != nil {
		c.Log.Warnf("Failed count user from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if totalItemCode > 0 {
		c.Log.Warnf("item already exists : %+v", err)
		return nil, fiber.ErrConflict
	}

	if err := c.ItemRepository.Create(tx, item); err != nil {
		c.Log.WithError(err).Error("error creating item")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating item")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ItemToResponse(item), nil
}

func (c *ItemUseCase) Update(ctx context.Context, request *model.UpdateItemRequest) (*model.ItemResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	item := new(entity.Item)
	if err := c.ItemRepository.FindById(tx, item, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting item")
		return nil, fiber.ErrNotFound
	}

	item.ItemName = request.ItemName
	item.Type = request.Type

	if err := c.ItemRepository.Update(tx, item); err != nil {
		c.Log.WithError(err).Error("error updating item")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating item")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ItemToResponse(item), nil
}

func (c *ItemUseCase) Get(ctx context.Context, request *model.GetItemRequest) (*model.ItemResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	item := new(entity.Item)
	if err := c.ItemRepository.FindById(tx, item, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting item")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting item")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ItemToResponse(item), nil
}

func (c *ItemUseCase) Delete(ctx context.Context, request *model.DeleteItemRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return fiber.ErrBadRequest
	}

	item := new(entity.Item)
	if err := c.ItemRepository.FindById(tx, item, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting item")
		return fiber.ErrNotFound
	}

	if err := c.ItemRepository.Delete(tx, item); err != nil {
		c.Log.WithError(err).Error("error deleting item")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting item")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *ItemUseCase) Search(ctx context.Context, request *model.SearchItemRequest) ([]model.ItemResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, 0, fiber.ErrBadRequest
	}

	items, total, err := c.ItemRepository.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Error("error getting items")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting items")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.ItemResponse, len(items))
	for i, item := range items {
		responses[i] = *converter.ItemToResponse(&item)
	}

	return responses, total, nil
}

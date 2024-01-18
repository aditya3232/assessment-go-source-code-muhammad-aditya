package http

import (
	"assessment-go-source-code-muhammad-aditya/internal/model"
	"assessment-go-source-code-muhammad-aditya/internal/usecase"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ItemController struct {
	UseCase *usecase.ItemUseCase
	Log     *logrus.Logger
}

func NewItemController(useCase *usecase.ItemUseCase, log *logrus.Logger) *ItemController {
	return &ItemController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *ItemController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateItemRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating item")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ItemResponse]{Data: response})
}

func (c *ItemController) List(ctx *fiber.Ctx) error {

	request := &model.SearchItemRequest{
		ItemName:  ctx.Query("item_name", ""),
		Type:      ctx.Query("type", ""),
		ItemPrice: ctx.QueryInt("item_price", 0),
		Page:      ctx.QueryInt("page", 1),
		Size:      ctx.QueryInt("size", 10),
	}

	responses, total, err := c.UseCase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching item")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.ItemResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func (c *ItemController) Get(ctx *fiber.Ctx) error {
	request := &model.GetItemRequest{
		ID: ctx.Params("itemId"),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error getting item")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ItemResponse]{Data: response})
}

func (c *ItemController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateItemRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("itemId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating item")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ItemResponse]{Data: response})
}

func (c *ItemController) Delete(ctx *fiber.Ctx) error {
	itemId := ctx.Params("itemId")

	request := &model.DeleteItemRequest{
		ID: itemId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("error deleting item")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}

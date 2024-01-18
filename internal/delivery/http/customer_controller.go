package http

import (
	"assessment-go-source-code-muhammad-aditya/internal/model"
	"assessment-go-source-code-muhammad-aditya/internal/usecase"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CustomerController struct {
	UseCase *usecase.CustomerUseCase
	Log     *logrus.Logger
}

func NewCustomerController(useCase *usecase.CustomerUseCase, log *logrus.Logger) *CustomerController {
	return &CustomerController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *CustomerController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateCustomerRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating customer")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CustomerResponse]{Data: response})
}

func (c *CustomerController) List(ctx *fiber.Ctx) error {

	request := &model.SearchCustomerRequest{
		Name: ctx.Query("name", ""),
		Page: ctx.QueryInt("page", 1),
		Size: ctx.QueryInt("size", 10),
	}

	responses, total, err := c.UseCase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching customer")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.CustomerResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func (c *CustomerController) Get(ctx *fiber.Ctx) error {
	request := &model.GetCustomerRequest{
		ID: ctx.Params("customerId"),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error getting customer")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CustomerResponse]{Data: response})
}

func (c *CustomerController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateCustomerRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("customerId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating customer")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CustomerResponse]{Data: response})
}

func (c *CustomerController) Delete(ctx *fiber.Ctx) error {
	customerId := ctx.Params("customerId")

	request := &model.DeleteCustomerRequest{
		ID: customerId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("error deleting customer")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}

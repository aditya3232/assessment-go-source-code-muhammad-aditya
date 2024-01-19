package http

import (
	"assessment-go-source-code-muhammad-aditya/internal/model"
	"assessment-go-source-code-muhammad-aditya/internal/usecase"
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type InvoiceController struct {
	UseCase *usecase.InvoiceUseCase
	Log     *logrus.Logger
}

func NewInvoiceController(useCase *usecase.InvoiceUseCase, log *logrus.Logger) *InvoiceController {
	return &InvoiceController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *InvoiceController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateInvoiceRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating invoice")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.InvoiceResponse]{Data: response})
}

func (c *InvoiceController) Get(ctx *fiber.Ctx) error {
	request := &model.GetInvoiceRequest{
		ID: ctx.Params("invoiceId"),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error getting invoice")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.InvoiceResponse]{Data: response})
}

func (c *InvoiceController) List(ctx *fiber.Ctx) error {
	issuedDateString := ctx.Query("issued_date", "")
	dueDateString := ctx.Query("due_date", "")

	var issuedDate, dueDate time.Time
	var err error

	if issuedDateString != "" {
		issuedDate, err = time.Parse("01/02/2006", issuedDateString)
		if err != nil {
			c.Log.WithError(err).Error("error issued date parsing")
			return err
		}
	}

	if dueDateString != "" {
		dueDate, err = time.Parse("01/02/2006", dueDateString)
		if err != nil {
			c.Log.WithError(err).Error("error due date parsing")
			return err
		}
	}

	request := &model.SearchInvoiceRequest{
		InvoiceNumber: ctx.Query("invoice_number", ""),
		Subject:       ctx.Query("subject", ""),
		TotalItem:     ctx.QueryInt("total_item", 0),
		CustomerId:    ctx.Query("customer_id", ""),
		Status:        ctx.Query("status", ""),
		Page:          ctx.QueryInt("page", 1),
		Size:          ctx.QueryInt("size", 10),
	}

	if !issuedDate.IsZero() {
		request.IssuedDate = issuedDate.UnixNano() / int64(time.Millisecond)
	}

	if !dueDate.IsZero() {
		request.DueDate = dueDate.UnixNano() / int64(time.Millisecond)
	}

	responses, total, err := c.UseCase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching invoice")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.InvoiceListResponse]{
		Data:   responses,
		Paging: paging,
	})
}

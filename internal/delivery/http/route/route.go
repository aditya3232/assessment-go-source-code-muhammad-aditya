package route

import (
	"assessment-go-source-code-muhammad-aditya/internal/delivery/http"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                *fiber.App
	CustomerController *http.CustomerController
	ItemController     *http.ItemController
	InvoiceController  *http.InvoiceController
}

func (c *RouteConfig) Setup() {
	c.App.Use(recoverPanic)
	c.SetupGuestRoute()
}

func recoverPanic(ctx *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("Panic occurred: %v", r)
			log.Println(err)
			ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
				"error":   err.Error(),
			})
		}
	}()

	return ctx.Next()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/api/customers", c.CustomerController.List)
	c.App.Post("/api/customers", c.CustomerController.Create)
	c.App.Put("/api/customers/:customerId", c.CustomerController.Update)
	c.App.Get("/api/customers/:customerId", c.CustomerController.Get)
	c.App.Delete("/api/customers/:customerId", c.CustomerController.Delete)

	c.App.Get("/api/items", c.ItemController.List)
	c.App.Post("/api/items", c.ItemController.Create)
	c.App.Put("/api/items/:itemId", c.ItemController.Update)
	c.App.Get("/api/items/:itemId", c.ItemController.Get)
	c.App.Delete("/api/items/:itemId", c.ItemController.Delete)

	c.App.Post("/api/invoices", c.InvoiceController.Create)
	c.App.Get("/api/invoices/:invoiceId", c.InvoiceController.Get)
	c.App.Get("/api/invoices", c.InvoiceController.List)
	c.App.Put("/api/invoices/:invoiceId", c.InvoiceController.Update)
}

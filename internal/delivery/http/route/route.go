package route

import (
	"assessment-go-source-code-muhammad-aditya/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                *fiber.App
	CustomerController *http.CustomerController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/api/customers", c.CustomerController.List)
	c.App.Post("/api/customers", c.CustomerController.Create)
	c.App.Put("/api/customers/:customerId", c.CustomerController.Update)
	c.App.Get("/api/customers/:customerId", c.CustomerController.Get)
	c.App.Delete("/api/customers/:customerId", c.CustomerController.Delete)
}

package config

import (
	"assessment-go-source-code-muhammad-aditya/internal/delivery/http"
	"assessment-go-source-code-muhammad-aditya/internal/delivery/http/route"
	"assessment-go-source-code-muhammad-aditya/internal/repository"
	"assessment-go-source-code-muhammad-aditya/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	customerRepository := repository.NewCustomerRepository(config.Log)

	// setup use cases
	customerUseCase := usecase.NewCustomerUseCase(config.DB, config.Log, config.Validate, customerRepository)

	// setup controller
	customerController := http.NewCustomerController(customerUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:                config.App,
		CustomerController: customerController,
	}
	routeConfig.Setup()
}

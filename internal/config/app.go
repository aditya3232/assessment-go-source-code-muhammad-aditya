package config

import (
	"assessment-go-source-code-muhammad-aditya/internal/delivery/http"
	"assessment-go-source-code-muhammad-aditya/internal/delivery/http/route"
	"assessment-go-source-code-muhammad-aditya/internal/gateway/messaging"
	"assessment-go-source-code-muhammad-aditya/internal/repository"
	"assessment-go-source-code-muhammad-aditya/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
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
	Writer   *kafka.Writer
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	customerRepository := repository.NewCustomerRepository(config.Log)
	itemRepository := repository.NewItemRepository(config.Log)
	invoiceRepository := repository.NewInvoiceRepository(config.Log)
	invoiceItemRepository := repository.NewInvoiceItemRepository(config.Log)

	// setup writer
	customerWriter := messaging.NewCustomerWriter(config.Writer, config.Log)

	// setup use cases
	customerUseCase := usecase.NewCustomerUseCase(config.DB, config.Log, config.Validate, customerRepository, customerWriter)
	itemUseCase := usecase.NewItemUseCase(config.DB, config.Log, config.Validate, itemRepository)
	invoiceUseCase := usecase.NewInvoiceUseCase(config.DB, config.Log, config.Validate, invoiceRepository, customerRepository, invoiceItemRepository)

	// setup controller
	customerController := http.NewCustomerController(customerUseCase, config.Log)
	itemController := http.NewItemController(itemUseCase, config.Log)
	invoiceController := http.NewInvoiceController(invoiceUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:                config.App,
		CustomerController: customerController,
		ItemController:     itemController,
		InvoiceController:  invoiceController,
	}
	routeConfig.Setup()
}

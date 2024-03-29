package controller

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
)

type BmiController struct {
	service.BmiService
	configuration.Config
}

func NewBmiController(bmiService *service.BmiService, config configuration.Config) *BmiController {
	return &BmiController{BmiService: *bmiService, Config: config}
}

func (controller BmiController) Route(app *fiber.App) {
	app.Post("/v1/api/bmi/calculator", controller.Calculator)
	app.Get("/v1/api/bmi/history", controller.History)
}

func (controller BmiController) Calculator(c *fiber.Ctx) error {
	var request model.BmiCreateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.BmiService.Calculator(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (controller BmiController) History(c *fiber.Ctx) error {
	response := controller.BmiService.FindAll(c.Context())
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

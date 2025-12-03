package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kodra-pay/fx-service/internal/handlers"
	"github.com/kodra-pay/fx-service/internal/services"
)

func Register(app *fiber.App, service string) {
	health := handlers.NewHealthHandler(service)
	health.Register(app)

	svc := services.NewFXService()
	h := handlers.NewFXHandler(svc)
	api := app.Group("/fx")
	api.Get("/rates", h.Rates)
	api.Post("/convert", h.Convert)
}

package dev

import (
	"time"

	"blurble.com/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type Ping struct {
	Alive bool `json:"alive"`
}

func health(c *fiber.Ctx) error {
	middleware.Cache(c, time.Minute)
	return c.JSON(Ping{Alive: true})
}

func Router(r fiber.Router) {
	handler := r.Group("/dev")
	// routes
	handler.Get("/health", health)
	handler.Get("/metrics", monitor.New(monitor.Config{Title: "Blurble Monitor"}))
}

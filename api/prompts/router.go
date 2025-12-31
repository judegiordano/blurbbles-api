package prompts

import (
	"time"

	"blurbbles/middleware"
	"blurbbles/pkg"

	"github.com/gofiber/fiber/v2"
)

func daily(c *fiber.Ctx) error {
	middleware.Cache(c, time.Minute)
	return c.JSON(pkg.GetDailyPrompt())
}

func Router(r fiber.Router) {
	handler := r.Group("/prompts")
	// routes
	handler.Get("/daily", daily)
}

package prompts

import (
	"time"

	"blurble.com/middleware"
	"blurble.com/pkg"
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

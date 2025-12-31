package prompts

import (
	"errors"
	"strings"
	"time"

	"blurbbles/middleware"
	"blurbbles/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/judegiordano/gogetem/pkg/fibererrors"
)

func daily(c *fiber.Ctx) error {
	middleware.Cache(c, time.Minute)
	return c.JSON(pkg.GetDailyPrompt())
}

func getById(c *fiber.Ctx) error {
	middleware.Cache(c, time.Hour)
	id := strings.TrimSpace(c.Params("id"))
	prompt := pkg.GetPromptById(id)
	if prompt.Id != id {
		return fibererrors.NotFound(c, errors.New("prompt not found"))
	}
	return c.JSON(prompt)
}

func Router(r fiber.Router) {
	handler := r.Group("/prompts")
	// routes
	handler.Get("/daily", daily)
	handler.Get("/:id", getById)
}

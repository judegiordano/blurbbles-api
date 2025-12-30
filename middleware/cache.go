package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Cache(c *fiber.Ctx, exp time.Duration) {
	ms := exp.Milliseconds()
	mss := strconv.Itoa(int(ms))
	c.Response().Header.Add(CACHE_KEY, mss)
	c.Response().Header.Add("Cache-Control", fmt.Sprintf("max-age=%v, public", exp.Seconds()))
}

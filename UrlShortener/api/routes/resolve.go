package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"

	"awesomeProject5/UrlShortener/api/database"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")
	r := database.CreateClient(0)
	defer r.Close()
	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short not found in the "})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error connecting in db "})
	}
	rinr := database.CreateClient(1)
	defer rinr.Close()
	_ = rinr.Incr(database.Ctx, "counter")
	return c.Redirect(value, 301)
}

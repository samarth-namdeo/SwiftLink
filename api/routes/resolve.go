package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/samarth-namdeo/SwiftLink/database"
)

func ResolveURL(c *fiber.Ctx) error {
	short := c.Params("short")
	r := database.CreateClient(0)
	defer r.Close()
	url, err := r.Get(database.Ctx, short).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Short not found",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot connect to db",
		})
	}

	iInr := database.CreateClient(1)
	defer iInr.Close()

	_ = iInr.Incr(database.Ctx, "counter")

	return c.Redirect(url, 301)
}

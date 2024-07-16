package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samarth-namdeo/SwiftLink/helpers"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {

	req := new(request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// implement rate limiting

	// check if the URL is valid
	if !goValidator.IsURL(req.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	// check for domain error
	if !helpers.RemoveDomainError(req.URL) {
		return c.Status(fiber.StaturServiceUnavailable).JSON(fiber.Map{
			"error": "URL is not allowed",
		})
	}

	// enforce https, ssl

	req.URL = helpers.EnforceHTTPS(req.URL)

}

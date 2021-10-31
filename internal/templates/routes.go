package templates

import "github.com/gofiber/fiber/v2"

// AddRoutesTo method configures mentions routes for the api
func AddRoutesTo(app fiber.Router, service *TemplateService) {
	app.Get("/templates", get(service))
}

func get(service *TemplateService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if e := service.Get(); e != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   e,
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"error":   nil,
		})
	}
}

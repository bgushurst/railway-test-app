package handler

import (
	"github.com/a-h/templ"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

func render(c *fiber.Ctx, component templ.Component) error {
	return adaptor.HTTPHandler(templ.Handler(component))(c)
}
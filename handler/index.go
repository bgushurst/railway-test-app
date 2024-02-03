package handler

import (
	"github.com/bgushurst/railway-test-app/view"
	"github.com/gofiber/fiber/v2"
)

type IndexHandler struct {}

func (h IndexHandler) HandleIndexShow(c *fiber.Ctx) error {
	return render(c, view.Hello("Brad"))
}
package url

import (
	"github.com/Bukulapak/KontakAPI/controller"
	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	page.Get("/", controller.Home)
	page.Get("/kontak", controller.GetAll)
}

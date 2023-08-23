package url

import (
	"github.com/Bukulapak/KontakAPI/controller"
	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	page.Get("/", controller.Home)
	page.Get("/kontak", controller.GetAll)
	page.Get("/kontak/:id", controller.GetKontakID)
	page.Post("/insert", controller.InsertData)
	page.Put("/update/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeleteKontak)
}

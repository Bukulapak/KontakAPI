package controller

import (
	"github.com/Bukulapak/KontakAPI/config"
	"github.com/gofiber/fiber/v2"
	inimodul "github.com/indrariksa/be_presensi/module"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/Bukulapak/KontakAPI",
		"message":     "You are at the root endpoint 😉",
		"success":     true,
	})
}

func GetAll(c *fiber.Ctx) error {
	ps := inimodul.GetAllContact(config.Ulbimongoconn, "kontak")
	return c.JSON(ps)
}

package controller

import (
	"github.com/Bukulapak/KontakAPI/config"
	"github.com/gofiber/fiber/v2"
	inimodel "github.com/indrariksa/be_presensi/model"
	inimodul "github.com/indrariksa/be_presensi/module"
	"net/http"
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

func InsertData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var kontak inimodel.Kontak
	if err := c.BodyParser(&kontak); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodul.InsertKontak(db, "kontak",
		kontak.NamaKontak,
		kontak.NomorHp,
		kontak.Alamat,
		kontak.Keterangan,
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

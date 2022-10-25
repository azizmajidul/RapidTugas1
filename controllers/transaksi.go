package controllers

import (
	"aziz/restshopcart/database"
	"aziz/restshopcart/models"
	//"strconv"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type TransaksiController struct {
	// Declare variables
	Db    *gorm.DB
	
}

func InitTransaksiController() *TransaksiController {
	db := database.DatabaseInit()
	// gorm sync
	db.AutoMigrate(&models.Transaksi{})

	return &TransaksiController{Db: db}
}
func (controller *TransaksiController) InsertCart(c *fiber.Ctx) error {
	var trans models.Transaksi
	var cart models.Cart
	 idcart :=  trans.CartId
	 if err := c.BodyParser(&trans); err != nil {
		return c.SendStatus(400)
	}

	err := models.DeleteCart(controller.Db, &cart, idcart)
	if err!=nil {
		return c.SendStatus(500)
	}

	errs := models.InsertTransaksi(controller.Db, &trans)
	if errs!=nil {
		return c.SendStatus(500)
	}
	return c.JSON(fiber.Map{
		"message" : "Berhasil",
		"Transaksi" : trans,
	})
}




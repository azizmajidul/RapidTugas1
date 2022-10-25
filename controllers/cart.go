package controllers



import (
	"aziz/restshopcart/database"
	"aziz/restshopcart/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type CartController struct {
	// Declare variables
	Db    *gorm.DB
	//store *session.Store
}


func InitCartController() *CartController {
	db := database.DatabaseInit()
	// gorm sync
	db.AutoMigrate(&models.Cart{})

	return &CartController{Db: db}
}

func (controller *CartController) InsertToCart(c *fiber.Ctx) error {
	//  params := c.AllParams() // "{"id": "1"}"

	// intCartId, _ := strconv.Atoi(params["cartid"])
	// intProductId, _ := strconv.Atoi(params["productid"])

	 var cart models.Cart
	 // var product models.Product
	if err := c.BodyParser(&cart); err != nil {
		return c.SendStatus(400)
	}

	err := models.CreateCart(controller.Db, &cart)
	if err != nil {
		return c.SendStatus(500)
	}
	// if succeed
	return c.JSON(fiber.Map{
		"Message" : "Berhasil",
		"Status" : 200,
		"Data" : cart,
	})


	// // Find the product first,
	// err := models.ReadProdukById(controller.Db, &product, intProductId)
	// if err != nil {
	// 	return c.SendStatus(500) // http 500 internal server error
	// }

	// // // Then find the cart
	// errs := models.ReadCartById(controller.Db, &cart, intCartId)
	// if errs != nil {
	// 	return c.SendStatus(500) // http 500 internal server error
	// }

	// // Finally, insert the product to cart
	// errss := models.InsertProductToCart(controller.Db, &cart)
	// if errss != nil {
	// 	return c.SendStatus(500) // http 500 internal server error
	// }

	// return c.JSON(fiber.Map{
	// 	"Message":"Success",
	// 	"Status" :200,
	// 	"Cart" :cart,
	// 	"Product" : product,
	// })
}

func (controller *CartController) GetShoppingCart(c *fiber.Ctx) error {
	
	
	var cart []models.Cart
	
	err := models.ReadAllProductsInCart(controller.Db, &cart)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	
	

	return c.JSON( fiber.Map{
		"Cart":cart,
	
		
		
	})
}

func (controller *CartController) ReadCartById(c *fiber.Ctx) error {
	id := c.Params("user_id")
	idn,_ := strconv.Atoi(id)


	var cart []models.Cart
	err := models.ReadCartById(controller.Db, &cart, idn)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(fiber.Map{
		"Message" : "Berhasil",
		"Status" : 200,
		"Data" : cart,
	})
}
package controllers

import (
	"fmt"
	"aziz/restshopcart/database"
	"aziz/restshopcart/models"
	"strconv"
	//"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


type ProductController struct {
	// declare variables
	Db *gorm.DB
	//store *session.Store
}

func InitProductController() *ProductController {
	db := database.DatabaseInit()
	// gorm
	db.AutoMigrate(&models.Product{})

	return &ProductController{Db: db}
}

//CREATE NEW PRODUCT
func (controller *ProductController) AddPostedProduct(c *fiber.Ctx) error {
	//myform := new(models.Product)
	var product models.Product
	

	if err := c.BodyParser(&product); err != nil {
		return c.SendStatus(400)
	}

	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		
	
		// Get all files from "documents" key:
		gambar := form.File["gambar"]
		// => []*multipart.FileHeader
	
		// Loop through files:
		for _, file := range gambar {
		  fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
		  // => "tutorial.pdf" 360641 "application/pdf"
	
		  // Save the files to disk:
		  if err := c.SaveFile(file, fmt.Sprintf("./upload/%s", file.Filename)); err != nil {
			return err
		  }

		  product.Gambar = file.Filename
		  
		}
		
	  }
	  
	// save product
	err := models.CreateProduct(controller.Db, &product)
	if err != nil {
		return c.SendStatus(500)
	}
	// if succeed
	return c.JSON(fiber.Map{
		"Message" : "Berhasil",
		"Status" : 200,
		"Data" : product,
	})
}


//Get All Product
func (controller *ProductController) GetAllProducts(c *fiber.Ctx) error {
	// load all products
	var products []models.Product
	err := models.ReadProducts(controller.Db, &products)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(fiber.Map{
		"Message" : "Berhasil",
		"Status" : 200,
		"Data" : products,
	})
}

//GET DETAIL PRODUCT

//GET DETAIL USER BY ID
func (controller *ProductController) GetDetailProduk(c *fiber.Ctx) error {
	id := c.Params("id")
	idn,_ := strconv.Atoi(id)


	var product models.Product
	err := models.ReadProdukById(controller.Db, &product, idn)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(fiber.Map{
		"Message" : "Berhasil",
		"Status" : 200,
		"Data" : product,
	})
}



func (controller *ProductController) EditProduk(c *fiber.Ctx) error {
	id := c.Params("id")
	idn,_ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProdukById(controller.Db, &product, idn)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	var updateProduct models.Product

	if err := c.BodyParser(&updateProduct); err != nil {
		return c.SendStatus(400)
	}


	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		
	
		// Get all files from "documents" key:
		gambar := form.File["gambar"]
		// => []*multipart.FileHeader
	
		// Loop through files:
		for _, file := range gambar {
		  fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
		  // => "tutorial.pdf" 360641 "application/pdf"
	
		  // Save the files to disk:
		  if err := c.SaveFile(file, fmt.Sprintf("./upload/%s", file.Filename)); err != nil {
			return err
		  }

		  product.Gambar = file.Filename
		  
		}
		
	  }

	
	product.Nama = updateProduct.Nama
	product.Kategori = updateProduct.Kategori
	product.Quantity = updateProduct.Quantity
	product.Harga = updateProduct.Harga
	product.Stok = updateProduct.Stok
	

	// save suer
	models.UpdateProduct(controller.Db, &product)
	
	return c.JSON(fiber.Map{
		"Message" : "Berhasil Update Data",
		"Status" : 200,
		"Data" : product,
	})
}

func (controller *ProductController) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn,_ := strconv.Atoi(id)

	var product models.Product
	models.DeleteProductById(controller.Db, &product, idn)

	//return c.JSON(user)	
	return c.JSON(fiber.Map{
		"message": "data was deleted",
	})
}










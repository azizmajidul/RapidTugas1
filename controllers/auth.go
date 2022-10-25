package controllers

import (
	//"fmt"
	"aziz/restshopcart/database"
	"aziz/restshopcart/models"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
)
type UserController struct {
	//declare variables
	Db    *gorm.DB
	//store *session.Store
}

//Models to SignIn

type SignIn struct {
	Email    string `form:"email" json:"email"  validate:"required"`
	Password string `form:"password"  json:"password"  validate:"required"`
}


func InitAuthController() *UserController {
	db := database.DatabaseInit()
	// gorm
	db.AutoMigrate(&models.User{})

	return &UserController{Db: db}
}

//INSERT USER OR REGISTTRASI

func (controller *UserController) RegisterUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(400) // bad request http
	}
	// save User
	err := models.Registrasi(controller.Db, &user)
	if err!=nil {
		return c.SendStatus(500)
	}
	// if succeed
	return c.JSON(fiber.Map{
		"Message" : "Berhasil",
		"Status" : 200,
		"Data" : user,
	})
}

//GET DATA ALL USER

func (controller *UserController) GetAllUser(c *fiber.Ctx) error {
	// load all User
	var user []models.User
	err := models.ReadUser(controller.Db, &user)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(fiber.Map{
		"Message" : "Berhasil",
		"Status" : 200,
		"Data" : user,
	})
}

//GET DETAIL USER BY ID
func (controller *UserController) GetDetailUser(c *fiber.Ctx) error {
	id := c.Params("id")
	idn,_ := strconv.Atoi(id)


	var user models.User
	err := models.ReadUserById(controller.Db, &user, idn)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(fiber.Map{
		"Message" : "Berhasil",
		"Status" : 200,
		"Data" : user,
	})
}

//UPDATE USER
func (controller *UserController) EditUser(c *fiber.Ctx) error {
	id := c.Params("id")
	idn,_ := strconv.Atoi(id)

	var user models.User
	err := models.ReadUserById(controller.Db, &user, idn)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	var updateUser models.User

	if err := c.BodyParser(&updateUser); err != nil {
		return c.SendStatus(400)
	}
	user.Nama = updateUser.Nama
	user.Email = updateUser.Email
	user.Telp = updateUser.Telp
	user.Password = updateUser.Password
	user.Alamat = updateUser.Alamat

	// save suer
	models.UpdateUser(controller.Db, &user)
	
	return c.JSON(fiber.Map{
		"Message" : "Berhasil Update Data",
		"Status" : 200,
		"Data" : user,
	})
}

func (controller *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	idn,_ := strconv.Atoi(id)

	var user models.User
	models.DeleteUserById(controller.Db, &user, idn)

	//return c.JSON(user)	
	return c.JSON(fiber.Map{
		"message": "data was deleted",
	})
}

//Login
func (controller *UserController) LoginUser(c *fiber.Ctx) error {
	var myform SignIn
	var user models.User
	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/SignIn")
	}
	err := models.GetUserByEmail(controller.Db, &user, myform.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with the given email is not found",
		}) // http 500 internal server error
	}
	if user.Email == myform.Email {
		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(myform.Password)); err != nil {

			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "Password is empty",
			})

		}
		return c.JSON(fiber.Map{
			"message": "Login Success",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Login Failed",
	})

}





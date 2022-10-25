package routes

import (
	"github.com/gofiber/fiber/v2"
	"aziz/restshopcart/controllers"
)

func RoutInit(route *fiber.App ){


	AuthControllers := controllers.InitAuthController()
	user := route.Group("/Auth") 
	user.Get("/", AuthControllers.GetAllUser)
	user.Post("/Register", AuthControllers.RegisterUser)
	user.Get("/user/detail/:id", AuthControllers.GetDetailUser)
	user.Post("/edituser/:id", AuthControllers.EditUser)
	user.Delete("/deleteuser/:id", AuthControllers.DeleteUser)
	user.Post("/SignIn",AuthControllers.LoginUser )

	//Module Product
	ProductControllers := controllers.InitProductController()
	product := route.Group("/Products") 
	product.Get("/", ProductControllers.GetAllProducts)
	product.Post("/tambah", ProductControllers.AddPostedProduct)
	product.Get("/produk/detail/:id", ProductControllers.GetDetailProduk)
	product.Put("/ediproduct/:id", ProductControllers.EditProduk)
	product.Delete("/delete/:id", ProductControllers.DeleteProduct)

	//Module Cart
	CartController := controllers.InitCartController()

	cart := route.Group("/Cart")
	cart.Post("/tambahcart",CartController.InsertToCart)
	cart.Get("/", CartController.GetShoppingCart)
	cart.Get("/byuserid/:user_id", CartController.ReadCartById)


	//Module Transaksi
	 TransaksiController := controllers.InitTransaksiController()
	trans := route.Group("/Transaki")
	trans.Post("/tambah", TransaksiController.InsertCart)
	


	

	

	
	

}
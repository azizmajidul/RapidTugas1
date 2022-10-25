package database

import(
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

var Db *gorm.DB
func DatabaseInit() *gorm.DB  {

	Db = connectDB()
	return Db
	
}

func connectDB() (*gorm.DB) {
	dsn := "host=localhost user=postgres password=Yuize411 dbname=shoppingCart port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err !=nil {

	fmt.Println("Error...")
	panic("Canot Connect To database")

	//return nil

}
fmt.Println("Coonect...")

return db
}
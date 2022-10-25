package models



import (
	"gorm.io/gorm"
)

// Product model info
// @Description Product information
type Product struct {
	gorm.Model
	Id 				int64 		`gorm:"primaryKey" json : "id"`
	Nama    		string  	`gorm:"type:varchar(128)" json :"nama"`
	Kategori     	string  	`gorm:"type:varchar(128)" json :"kategori"`
	Quantity 		int64    	` json:"quantity" validate:"required"`
	Harga    		float32 	` json:"harga" validate:"required"`
	Stok 			int64     	` json:"stok" validate:"required"`
	Gambar			string 		`gorm:"type:varchar(256)" form:"image" json:"gambar"`
	
	

	
}

// CRUD
func CreateProduct(db *gorm.DB, newProduct *Product) (err error) {
	err = db.Create(newProduct).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadProducts(db *gorm.DB, products *[]Product)(err error) {
	err = db.Find(products).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadProdukById(db *gorm.DB, product *Product, id int)(err error) {
	err = db.Where("id=?", id).First(product).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateProduct(db *gorm.DB, product *Product)(err error) {
	db.Save(product)
	
	return nil
}
func DeleteProductById(db *gorm.DB, product *Product, id int)(err error) {
	db.Where("id=?", id).Delete(product)
	
	return nil
}
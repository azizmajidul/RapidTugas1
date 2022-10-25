package models

import(
	"gorm.io/gorm"
)

type Cart struct{
	gorm.Model
	Id int64 `gorm:"primaryKey" json : "id"`
	UserId int `json:"user_id" form:"user_id"`
	ProductID  int `json:"product_id" form:"product_id"`
	Product     Product   `gorm:"foreignkey:ProductID"` 
	User     User   `gorm:"foreignkey:UserId"`
	Transaksi 	 []Transaksi `json:"cart" gorm:"many2many:transaktion"` 
	

}

func CreateCart(db *gorm.DB, newCart *Cart) (err error) {
	// newCart.UserId = userId
	// newCart.ProductId = product_id
	err = db.Create(newCart).Error
	if err != nil {
		return err
	}
	return nil
}

func InsertProductToCart(db *gorm.DB, insertedCart *Cart) (err error) {
	
	err = db.Save(insertedCart).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadAllProductsInCart(db *gorm.DB, cart *[]Cart) (err error) {
	 err = db.Preload("Product").Preload("User").Find(cart).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadCartById(db *gorm.DB, cart *[]Cart, id int) (err error) {
	err = db.Where("user_id=?", id).Preload("Product").Preload("User").First(cart).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCart(db *gorm.DB, products []*Product, newCart *Cart, userId uint) (err error) {
	db.Model(&newCart).Association("Product").Delete(products)

	return nil
}

func DeleteCart(db *gorm.DB, cart *Cart, id int)(err error) {
	db.Where("id=?", id).Delete(cart)
	
	return nil
}

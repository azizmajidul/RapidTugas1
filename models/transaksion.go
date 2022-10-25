package models

import (
	"gorm.io/gorm"
)

type Transaksi struct {
	gorm.Model
	Id       int `form:"id" json: "id" gorm:"primaryKey"`
	Cart 	 []Cart `json:"cart" gorm:"many2many:transaktion"`
	CartId	 int	`json:"cart_id" form:"cart_id" gorm:"-"`
}	
func CreateTransaksi(db *gorm.DB, newTrans *Transaksi) (err error) {
	// newCart.UserId = userId
	// newCart.ProductId = product_id
	err = db.Create(newTrans).Error
	if err != nil {
		return err
	}
	return nil
}

func InsertTransaksi(db *gorm.DB, insertedCart *Transaksi) (err error) {
	err = db.Save(insertedCart).Error
	if err != nil {
		return err
	}
	return nil
}


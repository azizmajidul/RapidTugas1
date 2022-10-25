package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct{
	Id int64 `gorm:"primaryKey" json : "id"`
	Nama string `gorm:"type:varchar(128)" json :"nama"`
	Email string `gorm:"type:varchar(128)" json :"email"`
	Telp string `gorm:"type:varchar(128)" json :"telp"`
	Password string `gorm:"type:varchar(128)" json :"password"`
	Alamat string `gorm:"type:text" json :"alamat"`
	
}

func Registrasi(db *gorm.DB, newUser *User) (err error) {
	plainPassword := newUser.Password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(plainPassword),10)
	sHash := string(bytes)
	fmt.Println("Hash password: ", sHash)
	newUser.Password = sHash
	err = db.Create(newUser).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(db *gorm.DB, users *User, email string) (err error) {
	err = db.Where("email=?", email).First(users).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadUser(db *gorm.DB, user *[]User)(err error) {
	err = db.Find(user).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadUserById(db *gorm.DB, user *User, id int)(err error) {
	err = db.Where("id=?", id).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, user *User)(err error) {
	plainPassword := user.Password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(plainPassword),10)
	sHash := string(bytes)
	fmt.Println("Hash password: ", sHash)
	user.Password = sHash
	db.Save(user)
	
	return nil
}
func DeleteUserById(db *gorm.DB, user *User, id int)(err error) {
	db.Where("id=?", id).Delete(user)
	
	return nil
}
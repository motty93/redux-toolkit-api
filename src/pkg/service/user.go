package service

import (
	"app/pkg/db/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewServiceUser(db *gorm.DB) *User {
	return &User{db: db}
}

// Session login process
func (u *User) Session(email string) (*model.User, error) {
	user := new(model.User)
	if err := u.db.Find(&user, "email=?", email).Error; err != nil {
		return nil, err
	}
	// if err := u.db.Find(&user, "email=? and password=?", email, password).Error; err != nil {
	// 	return nil, err
	// }

	return user, nil
}

// Users find all
func (u *User) Users() (*[]model.User, error) {
	users := new([]model.User)
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// User find by id user
func (u *User) User(id int) (*model.User, error) {
	user := new(model.User)
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Create user
func (u *User) Create(userReq *model.UserReq) error {
	// passwordの暗号化, 第二引数はコスト
	bs, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
	if err != nil {
		return err
	}

	user := model.User{
		Email:    userReq.Email,
		Password: bs,
	}
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// Update user
func (u *User) Update(nt *model.User, id int) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Updates(nt).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Delete user
func (u *User) Delete(id int) error {
	user := new(model.User)
	if err := u.db.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}

package dao

import (
	"errors"
	"gorm.io/gorm"
	"summer/server/cmd/user/model"
)

var (
	ErrNoSuchUser = errors.New("no such user")
	ErrUserExist  = errors.New("user already exist")
)

type UserManger struct {
	db *gorm.DB
}

// NewUserManger create a user dao.
func NewUserManger(db *gorm.DB) *UserManger {
	m := db.Migrator()
	if !m.HasTable(&model.User{}) {
		err := m.CreateTable(&model.User{})
		if err != nil {
			panic(err)
		}
	}
	return &UserManger{
		db: db,
	}
}

// GetUserByUsername get user by username
func (u *UserManger) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := u.db.Model(&model.User{}).
		Where(&model.User{Username: username}).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrNoSuchUser
	}
	return &user, err
}

// CreateUser creates a user.
func (u *UserManger) CreateUser(user *model.User) error {
	err := u.db.Model(&model.User{}).
		Where(&model.User{Username: user.Username}).First(&model.User{}).Error
	if err == nil {
		return ErrUserExist
	} else if err != gorm.ErrRecordNotFound {
		return err
	}
	return u.db.Model(&model.User{}).Create(user).Error
}

// DeleteUserById delete a user by id.
func (u *UserManger) DeleteUserById(userId int64) error {
	return u.db.Model(&model.User{}).Delete(&model.User{ID: userId}).Error
}

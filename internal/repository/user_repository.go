package repository

import (
	"delivery/internal/db"
	"delivery/internal/model"
)

// User repository interface and implementation

func GetUserByMobile(mobile string) (*model.Users, error) {

	var user model.Users

	err := db.DB.Where("mobile = ?", mobile).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

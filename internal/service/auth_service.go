package service

import (
	"delivery/internal/auth"
	"delivery/internal/repository"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Authentication service logic

type LoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserID uint   `json:"user_id"`
}

func Login(req LoginRequest) (*LoginResponse, error) {

	user, err := repository.GetUserByMobile(req.Mobile)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	pass, err := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(string(pass))

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, errors.New("password not matched " + user.Password + " " + req.Password)
	}

	token, err := auth.GenetateToken(user.ID)
	if err != nil {
		return nil, err
	}

	// Log login event asynchronously
	go logLoginAsync(user.Mobile, user.ID)

	return &LoginResponse{Token: token, UserID: user.ID}, nil
}

// Helper function to log login events asynchronously
func logLoginAsync(mobile string, userID uint) {
	log.Printf("📝 [Async] User %s (ID: %d) logged in successfully", mobile, userID)
}

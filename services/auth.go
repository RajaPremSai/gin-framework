package services

import (
	"fmt"
	model "gin-framework/internal/model"
	utils "gin-framework/utils"

	"gorm.io/gorm"
)

type AuthService struct{
	db *gorm.DB
}


func (a *AuthService) InitAuthService(database *gorm.DB){
	a.db=database;
	a.db.AutoMigrate((&model.User{}))
}

func (a *AuthService) LoginUser(email, password string) (*model.User, error) {
    if email == "" || password == "" {
        return nil, gorm.ErrInvalidData
    }

    var user model.User
    // Find user by email
    if err := a.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }

    // Compare the provided password with the stored hash
    if !utils.CheckPasswordHash(password, user.Password) {
        return nil, fmt.Errorf("invalid credentials")
    }

    return &user, nil
}

func (a *AuthService) RegisterUser (email,password string) (*model.User, error) {
	if email == "" || password == "" {
		return nil, gorm.ErrInvalidData
	}
	//Hash the password using a utility function
	hashedPassword,err := utils.HashPassword(password)
	if err != nil {
		return nil, gorm.ErrInvalidData // Return an error if hashing fails
	}
	// Create a new user instance
	user := &model.User{
		Email:    email,
		Password: hashedPassword,
	}

	// Check if user already exists with the given email, if exists return user cannot be created, if not create the user
	var existingUser model.User
	if err := a.db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, gorm.ErrRecordNotFound // User already exists
	} else if err != gorm.ErrRecordNotFound {
		return nil, err // Some other error occurred
	}

	// If user does not exist, proceed to create the new user
	// Create the new user
	if err := a.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}


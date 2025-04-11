package services

import (
	"errors"
	"go-jwt/initializers"
	"go-jwt/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	db *gorm.DB
}

func NewUserService(database *gorm.DB) IUserService {
	return &userService{db: database}
}

func (service *userService) CreateUser(
	email string,
	password string,
) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email must be present")
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	var existingUser models.User
	initializers.DB.First(&existingUser, "email = ?", email)
	if existingUser.ID != 0 {
		return nil, errors.New("email already in use")
	}

	// Create user
	user := &models.User{Email: email, Password: string(hash)}
	result := service.db.Create(user)
	if result.Error != nil {
		return nil, errors.New("failed to create user")
	}

	return user, nil
}

func (service *userService) GenerateJWT(
	email string,
	password string,
) (string, error) {
	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", email)

	if user.ID == 0 {
		return "", errors.New("invalid email or password")
	}

	// Compare passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", errors.New("failed to create token")
	}

	return tokenString, nil
}

func (service *userService) UpdateUserPassword(
	email string,
	password string,
	newPassword string,
	newPassword2 string,
) (*models.User, error) {
	if newPassword != newPassword2 {
		return nil, errors.New("new passwords must match")
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", email)
	if user.ID == 0 {
		return nil, errors.New("invalid email or password")
	}
	// Compare passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), 10)
	if err != nil {
		return nil, errors.New("error hashing password")
	}
	user.Password = string(hash)
	result := initializers.DB.Save(&user)
	if result.Error != nil {
		return nil, errors.New("failed to update password")
	}

	return &user, nil
}

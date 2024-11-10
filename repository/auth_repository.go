package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"koriebruh/management/domain"
)

type AuthRepository interface {
	Register(ctx *fiber.Ctx, db *gorm.DB, admin domain.Admin) error
	Login(ctx *fiber.Ctx, db *gorm.DB, admin domain.Admin) (uint, error) //catch id for save in cooke with jwt
	FindAllAdmin(ctx context.Context, db *gorm.DB) ([]domain.Admin, error)
}

type AuthRepositoryImpl struct {
}

func NewAuthRepository() *AuthRepositoryImpl {
	return &AuthRepositoryImpl{}
}

func (repo AuthRepositoryImpl) Register(ctx *fiber.Ctx, db *gorm.DB, admin domain.Admin) error {
	var existingAdmin domain.Admin

	if err := db.Where("email = ?", admin.Email).First(&existingAdmin).Error; err == nil {
		return errors.New("email already registered")
	}

	if err := db.Where("username = ?", admin.Username).First(&existingAdmin).Error; err == nil {
		return errors.New("username already taken")
	}

	if err := db.Create(&admin).Error; err != nil {
		return err
	}

	return nil
}

func (repo AuthRepositoryImpl) Login(ctx *fiber.Ctx, db *gorm.DB, admin domain.Admin) (uint, error) {
	var userLogin domain.Admin

	// Query user berdasarkan username
	result := db.Select("id", "username", "password").
		Where("username = ?", admin.Username).
		First(&userLogin)

	// Handle jika user tidak ditemukan
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("login failed: incorrect user and password")
		}
		return 0, fmt.Errorf("login failed: %v", result.Error)
	}

	// Validasi password
	errPass := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(admin.Password))
	if errPass != nil {
		return 0, errors.New("login failed: incorrect password")
	}

	// Return user id untuk JWT
	return userLogin.ID, nil
}

func (repo AuthRepositoryImpl) FindAllAdmin(ctx context.Context, db *gorm.DB) ([]domain.Admin, error) {
	var AllAdmin []domain.Admin

	tx := db.WithContext(ctx).Find(&AllAdmin)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return AllAdmin, nil
}

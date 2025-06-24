package controllers

import (
	"log"
	"net/http"
	"time"

	"api_bengkel/config"
	"api_bengkel/models"
	"api_bengkel/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=mekanik admin keuangan gudang customer"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Register binding error:", err.Error())
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Cek email sudah ada atau belum
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		utils.ResponseError(c, http.StatusBadRequest, "Email sudah terdaftar")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		log.Println("DB create error:", err.Error())
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal membuat user: "+err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, "Register berhasil", gin.H{
		"user_id": user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"role":    user.Role,
	})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Login binding error:", err.Error())
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		log.Println("Login failed, user not found:", input.Email)
		utils.ResponseError(c, http.StatusUnauthorized, "Email atau password salah")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.Println("Login failed, wrong password for email:", input.Email)
		utils.ResponseError(c, http.StatusUnauthorized, "Email atau password salah")
		return
	}

	secretKey := []byte("your-secret-key") // simpan di config/env lebih aman

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal generate token")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Login berhasil", gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

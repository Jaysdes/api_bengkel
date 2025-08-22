package controllers

import (
	"log"
	"net/http"
	"strings"
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
	Identifier string `json:"identifier" binding:"required"` // bisa email atau username
	Password   string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Login binding error:", err.Error())
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User

	// cek apakah identifier berupa email atau username
	if strings.Contains(input.Identifier, "@") {
		// login pakai email
		if err := config.DB.Where("email = ?", input.Identifier).First(&user).Error; err != nil {
			log.Println("Login failed, email not found:", input.Identifier)
			utils.ResponseError(c, http.StatusUnauthorized, "Email atau password salah")
			return
		}
	} else {
		// login pakai username
		if err := config.DB.Where("name = ?", input.Identifier).First(&user).Error; err != nil {
			log.Println("Login failed, username not found:", input.Identifier)
			utils.ResponseError(c, http.StatusUnauthorized, "Username atau password salah")
			return
		}
	}

	// cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.Println("Login failed, wrong password for:", input.Identifier)
		utils.ResponseError(c, http.StatusUnauthorized, "Email/Username atau password salah")
		return
	}

	// generate token JWT
	secretKey := []byte("your-secret-key") // simpan di env
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

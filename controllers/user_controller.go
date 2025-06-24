package controllers

import (
	"api_bengkel/config"
	"api_bengkel/models"
	"api_bengkel/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GET /api/users
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Select("id", "name", "email", "role", "created_at").
		Order("created_at desc").Find(&users).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal mengambil data pengguna")
		return
	}

	for i := range users {
		users[i].Status = "aktif"
	}

	utils.ResponseSuccess(c, http.StatusOK, "Berhasil mengambil data", users)
}
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user models.User
	if err := config.DB.Where("id = ?", idInt).First(&user).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "User not found")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Success", user)
}

// POST /api/users
func CreateUser(c *gin.Context) {
	var input models.User

	// Bind input langsung ke model User
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal mengenkripsi password")
		return
	}
	input.Password = string(hashedPassword)

	// Simpan ke database
	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal menyimpan pengguna")
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, "Pengguna berhasil ditambahkan", input)
}

// PUT /api/users/:id
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "ID pengguna tidak valid")
		return
	}

	var user models.User
	if err := config.DB.First(&user, idInt).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Pengguna tidak ditemukan")
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Input tidak valid")
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Role = input.Role

	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, "Gagal mengenkripsi password")
			return
		}
		user.Password = string(hashedPassword)
	}

	if err := config.DB.Save(&user).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal mengupdate pengguna")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Pengguna berhasil diupdate", user)
}

// DELETE /api/users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "ID pengguna tidak valid")
		return
	}

	var user models.User
	if err := config.DB.First(&user, idInt).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Pengguna tidak ditemukan")
		return
	}

	if err := config.DB.Unscoped().Delete(&user).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal menghapus pengguna")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Pengguna berhasil dihapus", nil)
}

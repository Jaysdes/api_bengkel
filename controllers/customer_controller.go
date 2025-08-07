package controllers

import (
	"net/http"
	"strconv"
	"time"

	"api_bengkel/config"
	"api_bengkel/models"
	"api_bengkel/utils"

	"github.com/gin-gonic/gin"
)

// ====================== JENIS KENDARAAN ==========================

func GetJenisKendaraans(c *gin.Context) {
	var jenis []models.JenisKendaraan
	if err := config.DB.Find(&jenis).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to fetch jenis kendaraan")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", jenis)
}

func GetJenisKendaraanByID(c *gin.Context) {
	id := c.Param("id")
	var jenis models.JenisKendaraan
	if err := config.DB.Where("id_jenis = ?", id).First(&jenis).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Jenis kendaraan not found")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", jenis)
}

func CreateJenisKendaraan(c *gin.Context) {
	var input models.JenisKendaraan
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to create jenis kendaraan")
		return
	}
	utils.ResponseSuccess(c, http.StatusCreated, "Jenis kendaraan created", input)
}

func UpdateJenisKendaraan(c *gin.Context) {
	id := c.Param("id")
	var jenis models.JenisKendaraan
	if err := config.DB.Where("id_jenis = ?", id).First(&jenis).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Jenis kendaraan not found")
		return
	}

	var input models.JenisKendaraan
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	jenis.JenisKendaraan = input.JenisKendaraan

	if err := config.DB.Save(&jenis).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to update jenis kendaraan")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Jenis kendaraan updated", jenis)
}

func DeleteJenisKendaraan(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid id")
		return
	}

	if err := config.DB.Delete(&models.JenisKendaraan{}, idInt).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to delete jenis kendaraan")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Jenis kendaraan deleted", nil)
}

// ========================= CUSTOMER ============================

func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	if err := config.DB.Find(&customers).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to fetch customers")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", customers)
}

func GetCustomerByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid customer ID")
		return
	}

	var customer models.Customer
	if err := config.DB.Where("id_customer = ?", idInt).First(&customer).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Customer not found")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", customer)
}
func CreateCustomer(c *gin.Context) {
	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.TanggalMasuk.IsZero() {
		input.TanggalMasuk = time.Now()
	}

	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to create customer")
		return
	}
	utils.ResponseSuccess(c, http.StatusCreated, "Data Customer Baru Berhasil di Buat", input)
}
func UpdateCustomer(c *gin.Context) {

	id := c.Param("id")

	var customer models.Customer
	if err := config.DB.Where("id_customer = ?", id).First(&customer).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Data Customer not found")
		return
	}

	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Update data
	customer.NamaCustomer = input.NamaCustomer
	customer.IDJenis = input.IDJenis
	customer.NoKendaraan = input.NoKendaraan
	customer.Alamat = input.Alamat
	customer.Telepon = input.Telepon
	customer.TanggalMasuk = input.TanggalMasuk

	// Simpan ke database
	if err := config.DB.Save(&customer).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to update customer")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Data Customer Berhasil Di Update", customer)
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid customer ID")
		return
	}

	if err := config.DB.Delete(&models.Customer{}, idInt).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to delete customer")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Customer deleted", nil)
}

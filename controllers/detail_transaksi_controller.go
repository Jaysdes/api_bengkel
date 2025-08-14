package controllers

import (
	"net/http"
	"strconv"

	"api_bengkel/config"
	"api_bengkel/models"
	"api_bengkel/utils"

	"github.com/gin-gonic/gin"
)

func GetDetailTransaksi(c *gin.Context) {
	var details []models.DetailTransaksi
	if err := config.DB.Order("id_detail desc").Find(&details).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to fetch data")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", details)
}

func GetDetailTransaksiByID(c *gin.Context) {
	id := c.Param("id")
	var detail models.DetailTransaksi
	if err := config.DB.Where("id_detail = ?", id).First(&detail).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Data not found")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", detail)
}
func CreateDetailTransaksi(c *gin.Context) {
	var input models.DetailTransaksi
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Ambil total dari tabel transaksi
	var transaksi models.Transaksi
	if err := config.DB.First(&transaksi, input.IDTransaksi).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Transaksi tidak ditemukan")
		return
	}

	// Set total dari transaksi
	input.Total = int64(transaksi.Total)

	// Hitung kembalian
	if input.Bayar > 0 {
		input.Kembalian = input.Bayar - input.Total
		if input.Kembalian < 0 {
			input.Kembalian = 0 // jika bayar kurang dari total
		}
	}

	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal membuat data")
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, "Created", input)
}
func UpdateDetailTransaksi(c *gin.Context) {
	id := c.Param("id")
	var detail models.DetailTransaksi
	if err := config.DB.First(&detail, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Data tidak ditemukan")
		return
	}

	var input models.DetailTransaksi
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Ambil total dari tabel transaksi
	var transaksi models.Transaksi
	if err := config.DB.First(&transaksi, input.IDTransaksi).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Transaksi tidak ditemukan")
		return
	}

	// Update field
	detail.IDTransaksi = input.IDTransaksi
	detail.NoSPK = input.NoSPK
	detail.IDCustomer = input.IDCustomer
	detail.NoKendaraan = input.NoKendaraan
	detail.IDSparepart = input.IDSparepart
	detail.IDService = input.IDService
	detail.IDJasa = input.IDJasa
	detail.Total = int64(transaksi.Total)

	detail.Bayar = input.Bayar

	// Hitung kembalian
	if detail.Bayar > 0 {
		detail.Kembalian = detail.Bayar - detail.Total
		if detail.Kembalian < 0 {
			detail.Kembalian = 0
		}
	}

	if err := config.DB.Save(&detail).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal update data")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Updated", detail)
}

func DeleteDetailTransaksi(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := config.DB.Delete(&models.DetailTransaksi{}, idInt).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to delete")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Deleted", nil)
}

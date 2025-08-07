package controllers

import (
	"net/http"
	"strconv"

	"api_bengkel/config"
	"api_bengkel/models"
	"api_bengkel/utils"

	"github.com/gin-gonic/gin"
)

func GetTransaksis(c *gin.Context) {
	var transaksis []models.Transaksi
	if err := config.DB.Find(&transaksis).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to fetch transaksi")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", transaksis)
}

func GetTransaksiByID(c *gin.Context) {
	id := c.Param("id")
	var transaksi models.Transaksi
	if err := config.DB.Where("id_transaksi = ?", id).First(&transaksi).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Transaksi not found")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", transaksi)
}

func CreateTransaksi(c *gin.Context) {
	var input models.Transaksi
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Hitung total dari harga jasa + harga sparepart
	input.Total = input.HargaJasa + input.HargaSparepart

	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to create transaksi")
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, "Transaksi created", input)
}

func UpdateTransaksi(c *gin.Context) {
	id := c.Param("id")
	var transaksi models.Transaksi
	if err := config.DB.Where("id_transaksi = ?", id).First(&transaksi).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Transaksi not found")
		return
	}

	var input models.Transaksi
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	transaksi.IDSPK = input.IDSPK
	transaksi.IDCustomer = input.IDCustomer
	transaksi.IDJenis = input.IDJenis
	transaksi.NoKendaraan = input.NoKendaraan
	transaksi.Telepon = input.Telepon
	transaksi.IDMekanik = input.IDMekanik
	transaksi.HargaJasa = input.HargaJasa
	transaksi.HargaSparepart = input.HargaSparepart
	transaksi.Total = input.HargaJasa + input.HargaSparepart

	if err := config.DB.Save(&transaksi).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to update transaksi")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Transaksi updated", transaksi)
}

func DeleteTransaksi(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid id")
		return
	}

	if err := config.DB.Delete(&models.Transaksi{}, idInt).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to delete transaksi")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Transaksi deleted", nil)
}

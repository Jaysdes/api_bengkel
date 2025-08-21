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

func GetSPK(c *gin.Context) {
	var spks []models.SPK
	if err := config.DB.Find(&spks).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to fetch SPK")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", spks)
}

func GetSPKByID(c *gin.Context) {
	id := c.Param("id")
	var spk models.SPK
	if err := config.DB.Where("id_spk = ?", id).First(&spk).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "SPK not found")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", spk)
}

// CreateSPK
func CreateSPK(c *gin.Context) {
	var input models.SPK
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.TanggalSPK.IsZero() {
		input.TanggalSPK = time.Now()
	}

	if input.Status == "" {
		input.Status = "di proses mekanik"
	}

	// catatan default kosong
	input.Catatan = ""

	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to create SPK")
		return
	}
	utils.ResponseSuccess(c, http.StatusCreated, "SPK created", input)
}

// UpdateSPK
func UpdateSPK(c *gin.Context) {
	id := c.Param("id")
	var spk models.SPK
	if err := config.DB.Where("id_spk = ?", id).First(&spk).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "SPK not found")
		return
	}

	var input models.SPK
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Update data umum
	spk.TanggalSPK = input.TanggalSPK
	spk.IDService = input.IDService
	spk.IDJasa = input.IDJasa
	spk.IDCustomer = input.IDCustomer
	spk.IDJenis = input.IDJenis
	spk.NoKendaraan = input.NoKendaraan
	spk.Keluhan = input.Keluhan

	// Catatan hanya bisa diisi lewat teknisi
	if input.Catatan != "" {
		spk.Catatan = input.Catatan
	}

	// Status otomatis
	if spk.Status == "di proses mekanik" {
		spk.Status = "sedang di proses"
	}
	if input.Status == "selesai" {
		spk.Status = "selesai di proses"
	}

	// Simpan perubahan SPK
	if err := config.DB.Save(&spk).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to update SPK")
		return
	}

	// Update tabel proses -> isi otomatis "sudah di proses mekanik"
	if err := config.DB.Model(&models.Proses{}).
		Where("id_spk = ?", spk.IDSpk).
		Update("keterangan", "sudah di proses mekanik").Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to update Proses")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "SPK updated", spk)
}

func DeleteSPK(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid SPK ID")
		return
	}

	if err := config.DB.Delete(&models.SPK{}, idInt).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to delete SPK")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "SPK deleted", nil)
}

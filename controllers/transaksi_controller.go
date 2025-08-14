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

	// Simpan transaksi
	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to create transaksi")
		return
	}

	// Tambahkan otomatis ke tabel proses
	proses := models.Proses{
		IDTransaksi: input.IDTransaksi,
		IDMekanik:   input.IDMekanik,
		Status:      "transaksi di proses",
		Keterangan:  "menunggu konfirmasi",
		WaktuMulai:  time.Now(),
	}
	if err := config.DB.Create(&proses).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to create proses")
		return
	}

	// Tambahkan otomatis ke tabel detail_transaksi tanpa bayar & kembalian
	detail := models.DetailTransaksi{
		IDTransaksi: input.IDTransaksi,
		NoSPK:       input.IDSPK,
		IDCustomer:  input.IDCustomer,
		NoKendaraan: input.NoKendaraan,
		IDSparepart: 0,
		IDService:   0,
		IDJasa:      0,
		Total:       int64(input.Total),
		Bayar:       0,
		Kembalian:   0,
	}
	if err := config.DB.Create(&detail).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to create detail transaksi")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Transaksi & Proses & Detail berhasil dibuat", input)
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
func BayarTransaksi(c *gin.Context) {
	id := c.Param("id")
	var transaksi models.Transaksi

	if err := config.DB.Where("id_transaksi = ?", id).First(&transaksi).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Transaksi tidak ditemukan")
		return
	}

	if transaksi.StatusPembayaran == "lunas" {
		utils.ResponseError(c, http.StatusBadRequest, "Pembayaran sudah lunas")
		return
	}

	// Update status transaksi
	transaksi.StatusPembayaran = "lunas"
	if err := config.DB.Save(&transaksi).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal memvalidasi pembayaran")
		return
	}

	// Update waktu_selesai pada proses
	config.DB.Model(&models.Proses{}).
		Where("id_transaksi = ?", transaksi.IDTransaksi).
		Update("waktu_selesai", time.Now()).
		Update("status", "transaksi selesai")

	utils.ResponseSuccess(c, http.StatusOK, "Pembayaran berhasil divalidasi", transaksi)
}

func CetakNota(c *gin.Context) {
	id := c.Param("id")
	var transaksi models.Transaksi

	if err := config.DB.Preload("Customer").Preload("Mekanik").
		Where("id_transaksi = ?", id).First(&transaksi).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Transaksi tidak ditemukan")
		return
	}

	// Kirim ke view Laravel (atau HTML render di Go)
	c.HTML(http.StatusOK, "nota.html", gin.H{
		"Transaksi": transaksi,
	})
}

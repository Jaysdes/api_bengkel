package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"api_bengkel/config"
	"api_bengkel/models"
	"api_bengkel/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GET /api/transaksi
func GetTransaksis(c *gin.Context) {
	var transaksis []models.Transaksi
	if err := config.DB.Find(&transaksis).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to fetch transaksi")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", transaksis)
}

// GET /api/transaksi/:id
func GetTransaksiByID(c *gin.Context) {
	id := c.Param("id")
	var transaksi models.Transaksi
	if err := config.DB.Where("id_transaksi = ?", id).First(&transaksi).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Transaksi not found")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", transaksi)
}

func toInt(v interface{}) (int, bool) {
	switch t := v.(type) {
	case nil:
		return 0, false
	case int:
		return t, true
	case int32:
		return int(t), true
	case int64:
		return int(t), true
	case float64:
		return int(t), true
	case string:
		if t == "" {
			return 0, false
		}
		n, err := strconv.Atoi(t)
		if err != nil {
			return 0, false
		}
		return n, true
	default:
		return 0, false
	}
}
func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch t := v.(type) {
	case string:
		return t
	default:
		return fmt.Sprintf("%v", v)
	}
}

// POST /api/transaksi
func CreateTransaksi(c *gin.Context) {
	var raw map[string]interface{}
	if err := c.ShouldBindJSON(&raw); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Payload JSON tidak valid")
		return
	}

	idSPK, okSPK := toInt(raw["id_spk"])
	if !okSPK || idSPK <= 0 {
		utils.ResponseError(c, http.StatusBadRequest, "id_spk wajib diisi dan harus valid")
		return
	}

	// Ambil SPK untuk fallback
	var spk struct {
		IDSPK       int    `gorm:"column:id_spk"`
		IDService   int    `gorm:"column:id_service"`
		IDJasa      int    `gorm:"column:id_jasa"`
		IDCustomer  int    `gorm:"column:id_customer"`
		IDJenis     int    `gorm:"column:id_jenis"`
		NoKendaraan string `gorm:"column:no_kendaraan"`
	}
	_ = config.DB.Table("spk").
		Select("id_spk,id_service,id_jasa,id_customer,id_jenis,no_kendaraan").
		Where("id_spk = ?", idSPK).
		Scan(&spk).Error

	idCustomer, okCust := toInt(raw["id_customer"])
	if (!okCust || idCustomer <= 0) && spk.IDCustomer > 0 {
		idCustomer = spk.IDCustomer
		okCust = true
	}
	if !okCust || idCustomer <= 0 {
		utils.ResponseError(c, http.StatusBadRequest, "id_customer wajib diisi (atau SPK.id_customer tidak ditemukan)")
		return
	}

	idMekanik, okMek := toInt(raw["id_mekanik"])
	if !okMek || idMekanik <= 0 {
		utils.ResponseError(c, http.StatusBadRequest, "id_mekanik wajib diisi")
		return
	}

	noKendaraan := toString(raw["no_kendaraan"])
	if noKendaraan == "" && spk.NoKendaraan != "" {
		noKendaraan = spk.NoKendaraan
	}

	idJenis, _ := toInt(raw["id_jenis"])
	if idJenis <= 0 && spk.IDJenis > 0 {
		idJenis = spk.IDJenis
	}
	setIDJenis := idJenis > 0

	hargaJasa, _ := toInt(raw["harga_jasa"])
	hargaSpare, _ := toInt(raw["harga_sparepart"])
	jenisService, _ := toInt(raw["jenis_service"])
	telepon := toString(raw["telepon"])

	trx := models.Transaksi{
		IDSPK:          idSPK,
		IDCustomer:     idCustomer,
		NoKendaraan:    noKendaraan,
		Telepon:        telepon,
		IDMekanik:      idMekanik,
		HargaJasa:      hargaJasa,
		HargaSparepart: hargaSpare,
		Total:          hargaJasa + hargaSpare,
	}
	if setIDJenis {
		trx.IDJenis = idJenis
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// 1) Transaksi
		if err := tx.Create(&trx).Error; err != nil {
			return fmt.Errorf("gagal membuat transaksi: %w", err)
		}

		// 2) Proses
		proses := models.Proses{
			IDTransaksi: trx.IDTransaksi,
			IDMekanik:   trx.IDMekanik,
			Status:      "transaksi di proses",
			Keterangan:  "menunggu konfirmasi",
			WaktuMulai:  time.Now(),
		}
		if err := tx.Create(&proses).Error; err != nil {
			return fmt.Errorf("gagal membuat proses: %w", err)
		}

		// 3) Detail Transaksi (id_sparepart dibiarkan NULL)
		resolvedService := 0
		if jenisService > 0 {
			resolvedService = jenisService
		} else if spk.IDService > 0 {
			resolvedService = spk.IDService
		}

		detailMap := map[string]interface{}{
			"id_transaksi": trx.IDTransaksi,
			"no_spk":       trx.IDSPK,
			"id_customer":  trx.IDCustomer,
			"no_kendaraan": trx.NoKendaraan,
			"total":        int64(trx.Total),
			"bayar":        0,
			"kembalian":    0,
			"status":       "Belum Lunas",
		}
		if resolvedService > 0 {
			detailMap["id_service"] = resolvedService
		}
		if spk.IDJasa > 0 {
			detailMap["id_jasa"] = spk.IDJasa
		}
		if err := tx.Table("detail_transaksi").Create(detailMap).Error; err != nil {
			return fmt.Errorf("gagal membuat detail_transaksi: %w", err)
		}

		// 4) LAPORAN sesuai tabel di DB (BUKAN id_proses)
		lap := models.Laporan{
			IDTransaksi:    trx.IDTransaksi,
			IDCustomer:     trx.IDCustomer,
			TotalBiaya:     int64(trx.Total),
			TanggalLaporan: time.Now(),
			Catatan:        "Transaksi dibuat & diproses",
		}
		if err := tx.Create(&lap).Error; err != nil {
			return fmt.Errorf("gagal membuat laporan: %w", err)
		}

		return nil
	})

	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Transaksi, Proses, Detail & Laporan berhasil dibuat", trx)
}

func BayarTransaksi(c *gin.Context) {
	id := c.Param("id")

	// 1) Ambil nilai bayar dari JSON, FORM, atau query string (toleran untuk hindari 400)
	var bayar int64

	// a) JSON { "bayar": number }
	var inJSON struct {
		Bayar int64 `json:"bayar"`
	}
	if err := c.ShouldBindJSON(&inJSON); err == nil && inJSON.Bayar > 0 {
		bayar = inJSON.Bayar
	} else {
		// b) FORM "bayar=<num>"
		var inForm struct {
			Bayar string `form:"bayar"`
		}
		_ = c.ShouldBind(&inForm) // tidak fatal
		if inForm.Bayar != "" {
			if v, err := strconv.ParseInt(inForm.Bayar, 10, 64); err == nil && v > 0 {
				bayar = v
			}
		}
		// c) Query ?bayar=<num>
		if bayar <= 0 {
			if q := c.Query("bayar"); q != "" {
				if v, err := strconv.ParseInt(q, 10, 64); err == nil && v > 0 {
					bayar = v
				}
			}
		}
	}

	if bayar <= 0 {
		utils.ResponseError(c, http.StatusBadRequest, "Jumlah bayar tidak valid")
		return
	}

	// 2) Ambil transaksi
	var trx models.Transaksi
	if err := config.DB.Where("id_transaksi = ?", id).First(&trx).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Transaksi tidak ditemukan")
		return
	}
	if trx.StatusPembayaran == "lunas" {
		utils.ResponseError(c, http.StatusBadRequest, "Transaksi sudah lunas")
		return
	}

	total := int64(trx.Total)
	if bayar < total {
		utils.ResponseError(c, http.StatusBadRequest, fmt.Sprintf("Jumlah bayar kurang dari total (%d < %d)", bayar, total))
		return
	}
	kembalian := bayar - total

	// 3) Transactional update
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Transaksi{}).
			Where("id_transaksi = ?", trx.IDTransaksi).
			Update("status_pembayaran", "lunas").Error; err != nil {
			return err
		}
		result := tx.Table("detail_transaksi").
			Where("id_transaksi = ?", trx.IDTransaksi).
			Updates(map[string]interface{}{
				"bayar":     bayar,
				"kembalian": kembalian,
				"status":    "lunas",
			})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			// buat minimal row agar tidak kosong (fallback)
			createMap := map[string]interface{}{
				"id_transaksi": trx.IDTransaksi,
				"no_spk":       trx.IDSPK,
				"id_customer":  trx.IDCustomer,
				"no_kendaraan": trx.NoKendaraan,
				"total":        total,
				"bayar":        bayar,
				"kembalian":    kembalian,
			}
			if err := tx.Table("detail_transaksi").Create(createMap).Error; err != nil {
				return err
			}
		}

		// c) Update proses -> selesai + waktu_selesai
		if err := tx.Model(&models.Proses{}).
			Where("id_transaksi = ?", trx.IDTransaksi).
			Updates(map[string]interface{}{
				"status":        "selesai",
				"keterangan":    "transaksi selesai",
				"waktu_selesai": time.Now(),
			}).Error; err != nil {
			return err
		}

		// d) Tambah catatan ke laporan (mengacu ke skema tabel laporan yang benar)
		lap := models.Laporan{
			IDTransaksi:    trx.IDTransaksi,
			IDCustomer:     trx.IDCustomer,
			TotalBiaya:     int64(trx.Total),
			TanggalLaporan: time.Now(),
			Catatan:        "Pembayaran lunas",
		}
		if err := tx.Create(&lap).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal memproses pembayaran: "+err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Pembayaran berhasil divalidasi", gin.H{
		"id_transaksi": trx.IDTransaksi,
		"total":        total,
		"bayar":        bayar,
		"kembalian":    kembalian,
		"status":       "lunas",
	})
}

// PUT /api/transaksi/:id
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

// DELETE /api/transaksi/:id
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

func CetakNota(c *gin.Context) {
	id := c.Param("id")

	var detail models.DetailTransaksi
	if err := config.DB.Where("id_detail = ?", id).First(&detail).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Detail transaksi tidak ditemukan")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Success", detail)
}

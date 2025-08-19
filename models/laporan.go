package models

import "time"

type Laporan struct {
	IDLaporan      int       `gorm:"primaryKey;column:id_laporan" json:"id_laporan"`
	IDTransaksi    int       `gorm:"column:id_transaksi" json:"id_transaksi"`
	IDCustomer     int       `gorm:"column:id_customer" json:"id_customer"`
	TotalBiaya     int64     `gorm:"column:total_biaya" json:"total_biaya"`
	TanggalLaporan time.Time `gorm:"column:tanggal_laporan" json:"tanggal_laporan"`
	Catatan        string    `gorm:"column:catatan" json:"catatan"`
}

func (Laporan) TableName() string {
	return "laporan"
}

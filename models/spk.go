package models

import "time"

type SPK struct {
	IDSpk       int       `gorm:"primaryKey;column:id_spk" json:"id_spk"`
	TanggalSPK  time.Time `gorm:"column:tanggal_spk" json:"tanggal_spk"`
	IDService   int       `gorm:"column:id_service" json:"id_service"`
	IDJasa      int       `gorm:"column:id_jasa" json:"id_jasa"`
	IDCustomer  int       `gorm:"column:id_customer" json:"id_customer"`
	IDJenis     int       `gorm:"column:id_jenis" json:"id_jenis"`
	NoKendaraan string    `gorm:"column:no_kendaraan" json:"no_kendaraan"`
	Keluhan     string    `gorm:"column:keluhan" json:"keluhan"`
}

func (SPK) TableName() string {
	return "spk"
}

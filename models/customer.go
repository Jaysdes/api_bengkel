package models

import "time"

type Customer struct {
	IDCustomer   uint      `gorm:"primaryKey;column:id_customer" json:"id_customer"`
	NamaCustomer string    `gorm:"column:nama_customer" json:"nama_customer"`
	IDJenis      int       `gorm:"column:id_jenis" json:"id_jenis"`
	NoKendaraan  string    `gorm:"column:no_kendaraan" json:"no_kendaraan"`
	Alamat       string    `gorm:"column:alamat" json:"alamat"`
	Telepon      string    `gorm:"column:telepon" json:"telepon"`
	TanggalMasuk time.Time `gorm:"column:tanggal_masuk" json:"tanggal_masuk" time_format:"2006-01-02" time_utc:"true"`
}

func (Customer) TableName() string {
	return "customer"
}

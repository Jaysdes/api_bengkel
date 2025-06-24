package models

type JenisKendaraan struct {
	IDJenis        int    `gorm:"primaryKey;column:id_jenis" json:"id_jenis"`
	JenisKendaraan string `gorm:"column:jenis_kendaraan" json:"jenis_kendaraan"`
}

// Jika nama tabel berbeda dengan nama struct plural, tambahkan method TableName()
func (JenisKendaraan) TableName() string {
	return "jenis_kendaraan"
}

package models

type Proses struct {
	IDProses     int    `gorm:"primaryKey;column:id_proses" json:"id_proses"`
	IDTransaksi  int    `gorm:"column:id_transaksi" json:"id_transaksi"`
	IDMekanik    int    `gorm:"column:id_mekanik" json:"id_mekanik"`
	Status       string `gorm:"column:status" json:"status"`
	Keterangan   string `gorm:"column:keterangan" json:"keterangan"`
	WaktuMulai   string `gorm:"column:waktu_mulai" json:"waktu_mulai"`
	WaktuSelesai string `gorm:"column:waktu_selesai" json:"waktu_selesai"`
}

// Nama tabel sesuai database
func (Proses) TableName() string {
	return "proses"
}

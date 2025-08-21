package models

import "time"

type Proses struct {
	IDProses     int        `json:"id_proses" gorm:"primaryKey;autoIncrement"`
	IDTransaksi  int        `json:"id_transaksi"`
	IDSPK        int        `json:"id_spk"`
	IDMekanik    int        `json:"id_mekanik"`
	Status       string     `json:"status" gorm:"type:varchar(50)"`
	Keterangan   string     `json:"keterangan" gorm:"type:varchar(255)"`
	WaktuMulai   time.Time  `json:"waktu_mulai" gorm:"type:datetime"`
	WaktuSelesai *time.Time `json:"waktu_selesai" gorm:"type:datetime"`
}

// Nama tabel sesuai database
func (Proses) TableName() string {
	return "proses"
}

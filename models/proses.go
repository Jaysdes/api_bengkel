package models

type Proses struct {
	IDProses    int    `gorm:"primaryKey;column:id_proses" json:"id_proses"`
	IDTransaksi int    `gorm:"column:id_transaksi" json:"id_transaksi"`
	Tanggal     string `gorm:"column:tanggal" json:"tanggal"`
	Keterangan  string `gorm:"column:keterangan" json:"keterangan"`
	Status      string `gorm:"column:status" json:"status"`
}

package models

type Laporan struct {
	IDLaporan  int    `gorm:"primaryKey;column:id_laporan" json:"id_laporan"`
	IDProses   int    `gorm:"column:id_proses" json:"id_proses"`
	Tanggal    string `gorm:"column:tanggal" json:"tanggal"`
	Keterangan string `gorm:"column:keterangan" json:"keterangan"`
}

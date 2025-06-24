package models

type Transaksi struct {
	IDTransaksi int    `gorm:"primaryKey;column:id_transaksi" json:"id_transaksi"`
	IDSPK       int    `gorm:"column:id_spk" json:"id_spk"`
	IDCustomer  int    `gorm:"column:id_customer" json:"id_customer"`
	IDJenis     int    `gorm:"column:id_jenis" json:"id_jenis"`
	NoKendaraan string `gorm:"column:no_kendaraan" json:"no_kendaraan"`
	Pemilik     string `gorm:"column:pemilik" json:"pemilik"`
	Telepon     string `gorm:"column:telepon" json:"telepon"`
	IDMekanik   int    `gorm:"column:id_mekanik" json:"id_mekanik"`
}

// Nama tabel sesuai database
func (Transaksi) TableName() string {
	return "transaksi"
}

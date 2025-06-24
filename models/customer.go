package models

type Customer struct {
	IDCustomer   uint   `gorm:"primaryKey;column:id_customer" json:"id_customer"`
	NamaCustomer string `gorm:"column:nama_customer" json:"nama_customer"`
	IDJenis      int    `gorm:"column:id_jenis" json:"id_jenis"`
	NoKendaraan  string `gorm:"column:no_kendaraan" json:"no_kendaraan"`
	Alamat       string `gorm:"column:alamat" json:"alamat"`
	Telepon      string `gorm:"column:telepon" json:"telepon"`
}

func (Customer) TableName() string {
	return "customer"
}

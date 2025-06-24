package models

type Sparepart struct {
	IDSparepart   int    `gorm:"primaryKey;column:id_sparepart" json:"id_sparepart"`
	NamaSparepart string `gorm:"column:nama_sparepart" json:"nama_sparepart"`
	HargaBeli     int    `gorm:"column:harga_beli" json:"harga_beli"`
	HargaJual     int    `gorm:"column:harga_jual" json:"harga_jual"`
	Stok          int    `gorm:"column:stok" json:"stok"`
}

func (Sparepart) TableName() string {
	return "sparepart"
}

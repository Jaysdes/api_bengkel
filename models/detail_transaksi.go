package models

type DetailTransaksi struct {
	IDDetail    int    `gorm:"primaryKey;column:id_detail" json:"id_detail"`
	NoSPK       int    `gorm:"column:no_spk" json:"no_spk"`
	IDCustomer  int    `gorm:"column:id_customer" json:"id_customer"`
	NoKendaraan string `gorm:"column:no_kendaraan" json:"no_kendaraan"`
	IDSparepart int    `gorm:"column:id_sparepart" json:"id_sparepart"`
	IDService   int    `gorm:"column:id_service" json:"id_service"`
	IDJasa      int    `gorm:"column:id_jasa" json:"id_jasa"`
	Qty         int    `gorm:"column:qty" json:"qty"`
	Total       int64  `gorm:"column:total" json:"total"`
}

func (DetailTransaksi) TableName() string {
	return "detail_transaksi"
}

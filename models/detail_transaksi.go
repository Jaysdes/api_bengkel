package models

type DetailTransaksi struct {
	IDDetail    int    `gorm:"primaryKey;column:id_detail" json:"id_detail"`
	IDTransaksi int    `gorm:"column:id_transaksi" json:"id_transaksi"`
	NoSPK       int    `gorm:"column:no_spk" json:"no_spk"`
	IDCustomer  int    `gorm:"column:id_customer" json:"id_customer"`
	NoKendaraan string `gorm:"column:no_kendaraan" json:"no_kendaraan"`
	IDSparepart *int   `gorm:"column:id_sparepart" json:"id_sparepart,omitempty"` // nullable
	IDService   *int   `gorm:"column:id_service" json:"id_service,omitempty"`     // nullable
	IDJasa      *int   `gorm:"column:id_jasa" json:"id_jasa,omitempty"`           // nullable
	Total       int64  `gorm:"column:total" json:"total"`
	Bayar       int64  `gorm:"column:bayar" json:"bayar"`
	Kembalian   int64  `gorm:"column:kembalian" json:"kembalian"`
}

func (DetailTransaksi) TableName() string {
	return "detail_transaksi"
}


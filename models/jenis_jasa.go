package models

type JenisJasa struct {
	IDJasa    uint   `gorm:"primaryKey;column:id_jasa" json:"id_jasa"`
	NamaJasa  string `gorm:"column:nama_jasa;type:varchar(100)" json:"nama_jasa"`
	HargaJasa int64  `gorm:"column:harga_jasa;type:bigint" json:"harga_jasa"`
}

func (JenisJasa) TableName() string {
	return "jenis_jasa"
}

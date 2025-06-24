package models

type Mekanik struct {
	IDMekanik    int    `gorm:"primaryKey;column:id_mekanik" json:"id_mekanik"`
	NamaMekanik  string `gorm:"column:nama_mekanik" json:"nama_mekanik"`
	JenisKelamin string `gorm:"type:enum('L','P');column:jenis_kelamin" json:"jenis_kelamin"`
	Alamat       string `gorm:"column:alamat" json:"alamat"`
	Telepon      string `gorm:"column:telepon" json:"telepon"`
}

func (Mekanik) TableName() string {
	return "mekanik"
}

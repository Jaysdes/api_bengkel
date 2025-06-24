package models

type JenisService struct {
	IDService    int    `gorm:"primaryKey;column:id_service" json:"id_service"`
	JenisService string `gorm:"column:jenis_service" json:"jenis_service"`
}

func (JenisService) TableName() string {
	return "jenis_service"
}

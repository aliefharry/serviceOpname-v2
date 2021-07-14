package entity

// import(
// 	"time"
// )

type Opname struct {
	ID 				uint64 `gorm:"primary_key:auto_increment" json:"id"`
	No 				string `gorm:"type:varchar(255)" json:"no"`
	Id_aktiva 		string `gorm:"type:varchar(255)" json:"id_aktiva"`
	Kategori 		string `gorm:"type:varchar(255)" json:"kategori"`
	Nama 			string `gorm:"type:varchar(255)" json:"nama"`
	Lokasi 			string `gorm:"type:varchar(255)" json:"lokasi"`
	Foto 			string `gorm:"type:varchar(255)" json:"foto"`
	Tanda_opname 	string `gorm:"type:varchar(255)" json:"tanda_opname"`
	Kondisi 		string `gorm:"type:varchar(255)" json:"kondisi"`
	Keterangan 		string `gorm:"type:varchar(255)" json:"keterangan"`
	UserID 			uint64 `gorm:"type:varchar(255)" json:"user"`
	User 			User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	UpdateDate		string `json:"updatedate" form:"updatedate"`
}

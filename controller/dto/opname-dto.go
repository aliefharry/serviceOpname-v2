package dto

// import (
//     "time"
// )

type OpnameUpdDTO struct {
	ID 				uint64 `json:"id" form:"id" binding:"required"`
	Lokasi 			string `json:"lokasi" form:"lokasi" binding:"required"`
	Foto 			string `json:"foto" form:"foto" binding:"required"`
	Tanda_opname 	string `json:"tanda_opname" form:"tanda_opname" binding:"required"`
	Kondisi 		string `json:"kondisi" form:"kondisi" binding:"required"`
	Keterangan 		string `json:"keterangan" form:"keterangan" binding:"required"`
	UserID 			uint64 `json:"user" form:"user" binding:"required"`
	UpdateDate		string `json:"updatedate" form:"updatedate"`
}


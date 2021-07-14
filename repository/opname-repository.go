package repository

import(
	"serviceOpname-v2/config/entity"
	"serviceOpname-v2/config/entity/helper"
	"gorm.io/gorm"
	"math"
)

type OpnameRepository interface {
	AllOpname() []entity.Opname
	UpdateOpname(b entity.Opname) entity.Opname
	FindOpnameByID(opnameID uint64) entity.Opname
	Pagination(pagination helper.Pagination) (RepositoryResult, int)
	// GetAllOpname(opname *entity.Opname, pagination *helper.Pagination) (*[]entity.Opname, error)
}

type opnameConnection struct {
	connection *gorm.DB
}

func NewOpnameRepository(dbConn *gorm.DB) OpnameRepository {
	return &opnameConnection {
		connection: dbConn,
	}
}

func (db *opnameConnection) UpdateOpname(b entity.Opname) entity.Opname {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *opnameConnection) FindOpnameByID(opnameID uint64) entity.Opname {
	var opname entity.Opname
	db.connection.Preload("User").Find(&opname, opnameID)
	return opname
}

// func (db *opnameConnection) AllOpname() []entity.Opname {
// 	var opnames []entity.Opname
// 	db.connection.Preload("User").Find(&opnames)
// 	return opnames
// }

func (db *opnameConnection) AllOpname() []entity.Opname {
	var opnames []entity.Opname
	db.connection.Preload("User").Find(&opnames)
	return opnames
}

//jangan dibaca
// func (db *opnameConnection) GetAllOpname(opname *entity.Opname, pagination *helper.Pagination) (*[]entity.Opname, error){
// 	var opnames []entity.Opname

// 	offset := (pagination.Page - 1) * pagination.Limit
// 	queryBuider := db.connection.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
// 	result := queryBuider.Model(&entity.Opname{}).Where(opname).Find(&opnames)
// 	if result.Error != nil{
// 		msg := result.Error
// 		return nil, msg
// 	}

// 	return &opnames, nil
// }

func (db *opnameConnection) Pagination(pagination helper.Pagination) (RepositoryResult, int) {
	var contacts []entity.Opname
	//var pagination helper.Pagination

	totalRows, totalPages, fromRow, toRow := int64(0), 0, 0, 0
	pagination.Page = 1
	offset := pagination.Page * pagination.Limit
	// get data with limit, offset & order
	find := db.connection.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	find = find.Find(&contacts)
	// has error find data
	errFind := find.Error

	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	pagination.Rows = contacts
	// count all data
	errCount := db.connection.Model(&entity.Opname{}).Count(&totalRows).Error

	if errCount != nil {
		return RepositoryResult{Error: errCount}, totalPages
	}

	pagination.TotalRows = int(totalRows)
	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}
	if toRow > int(totalRows) {
		// set to row with total rows
		toRow = int(totalRows)
	}
	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return RepositoryResult{Result: pagination}, totalPages
}

package repository

import(
	"serviceOpname-v2/config/entity"
	// "serviceOpname-v2/config/entity/helper"
	"gorm.io/gorm"
	"math"
)

type OpnameRepository interface {
	AllOpname() []entity.Opname
	UpdateOpname(b entity.Opname) entity.Opname
	FindOpnameByID(opnameID uint64) entity.Opname
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

func (db *opnameConnection) Pagination() (RepositoryResult) {
	var contacts entity.Opname

	totalRows, totalPages, fromRow, toRow := int64(0), 0, 0, 0

	var Page int
	var Limit int
	var Sort string
	var Rows interface{}
	var TotalRows int
	var FromRow int
	var ToRow int

	offset := Page * Limit

	// get data with limit, offset & order
	find := db.connection.Limit(Limit).Offset(offset).Order(Sort)

	find = find.Find(&contacts)

	// has error find data
	errFind := find.Error

	if errFind != nil {
		return RepositoryResult{Error: errFind}
	}

	Rows = contacts

	// count all data
	errCount := db.connection.Model(&entity.Opname{}).Count(&totalRows).Error

	if errCount != nil {
		return RepositoryResult{Error: errCount}
	}

	TotalRows = int(totalRows)

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows)/float64(Limit))) - 1

	if Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = Limit
	} else {
		if Page <= totalPages {
			// calculate from & to row
			fromRow = Page*Limit + 1
			toRow = (Page + 1) * Limit
		}
	}

	if toRow > int(totalRows) {
		// set to row with total rows
		toRow = int(totalRows)
	}

	FromRow = fromRow
	ToRow = toRow

	return RepositoryResult{Result: pagination}
}

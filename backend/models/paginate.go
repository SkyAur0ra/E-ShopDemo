package models

import (
	"backend/settings"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, p int) (*gorm.DB, int, int, int, int) {
	// p is current page's number
	if p <= 0 {
		p = 1
	}
	var count int64
	db.Count(&count)
	pageCount := int(count) / settings.PageSize

	if int(count) % settings.PageSize > 0 {
		pageCount += 1
	}

	if p >= pageCount {
		p = pageCount
	}

	previous := 1
	if p >= 0 {
		previous = p - 1
	}
	next := p + 1
	if next > pageCount {
		next = pageCount
	}
	
	offset := (p - 1) * settings.PageSize
	res := db.Offset(offset).Limit(settings.PageSize)
	return res, previous, next, int(count), pageCount
}
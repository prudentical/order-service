package persistence

import (
	"fmt"

	"gorm.io/gorm"
)

func Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * size).Limit(size)
	}
}

func OptionalFilter(query *gorm.DB, fieldName string, value *string) *gorm.DB {
	if value != nil {
		return query.Where(fmt.Sprintf("%s = ?", fieldName), value)
	}
	return query
}

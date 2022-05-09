package dao

import (
	"gorm.io/gorm"
)

type baseDao struct {
	db *gorm.DB
}

func (d baseDao) getById(id int) (interface{}, error) {
	return nil, nil
}

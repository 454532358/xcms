package dao

import (
	"sync"
	"xcms/app/model"
)

var onece sync.Once

type CategoryDao struct {
	*baseDao
}

func (c CategoryDao) getList() ([]*model.Category, error) {
	var categories []*model.Category

	return categories, nil
}

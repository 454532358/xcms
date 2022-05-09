package service

import (
	"gorm.io/gorm"
	"xcms/app/model"
)

type AdminService struct {
	Db *gorm.DB
}

func (adminService AdminService) GetList(currPage int, pageSize int) ([]*model.Admin, error) {
	var admins []*model.Admin
	err := adminService.Db.Limit(pageSize).Offset((currPage - 1) * pageSize).Find(&admins).Error
	return admins, err
}
func (adminService AdminService) store(admin *model.Admin) error {
	return adminService.Db.Create(admin).Error
}

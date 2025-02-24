package vendor

import (
	"tdp-cloud/module/dborm"
)

// 创建厂商

type CreateParam struct {
	UserId      uint
	SecretId    string `binding:"required"`
	SecretKey   string `binding:"required"`
	Provider    string `binding:"required"`
	Description string `binding:"required"`
}

func Create(data *CreateParam) (uint, error) {

	item := &dborm.Vendor{
		UserId:      data.UserId,
		SecretId:    data.SecretId,
		SecretKey:   data.SecretKey,
		Provider:    data.Provider,
		Description: data.Description,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新厂商

type UpdateParam struct {
	Id          uint
	UserId      uint
	SecretId    string
	SecretKey   string
	Provider    string
	Description string
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Vendor{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Updates(dborm.Vendor{
			SecretId:    data.SecretId,
			SecretKey:   data.SecretKey,
			Provider:    data.Provider,
			Description: data.Description,
		})

	return result.Error

}

// 删除厂商

type DeleteParam struct {
	Id     uint
	UserId uint
}

func Delete(data *DeleteParam) error {

	var item *dborm.Vendor

	result := dborm.Db.
		Where(&dborm.Vendor{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&item)

	return result.Error

}

// 获取厂商

type FetchParam struct {
	Id     uint
	UserId uint
}

func Fetch(data *FetchParam) (*dborm.Vendor, error) {

	var item *dborm.Vendor

	result := dborm.Db.
		Where(&dborm.Vendor{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		First(&item)

	return item, result.Error

}

// 获取厂商列表

type FetchAllParam struct {
	UserId   uint
	Provider string
}

func FetchAll(data *FetchAllParam) ([]*dborm.Vendor, error) {

	var items []*dborm.Vendor

	result := dborm.Db.
		Where(&dborm.Vendor{
			UserId:   data.UserId,
			Provider: data.Provider,
		}).
		Find(&items)

	return items, result.Error

}

// 获取厂商总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&dborm.Vendor{}).
		Where(&dborm.Vendor{
			UserId:   data.UserId,
			Provider: data.Provider,
		}).
		Count(&count)

	return count, result.Error

}

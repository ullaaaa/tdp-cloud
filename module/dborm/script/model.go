package script

import (
	"tdp-cloud/module/dborm"
)

// 创建脚本

type CreateParam struct {
	UserId        uint
	Name          string `binding:"required"`
	CommandType   string `binding:"required"`
	Username      string `binding:"required"`
	WorkDirectory string `binding:"required"`
	Content       string `binding:"required"`
	Description   string
	Timeout       uint `binding:"required"`
}

func Create(data *CreateParam) (uint, error) {

	item := &dborm.Script{
		UserId:        data.UserId,
		Name:          data.Name,
		CommandType:   data.CommandType,
		Username:      data.Username,
		WorkDirectory: data.WorkDirectory,
		Content:       data.Content,
		Description:   data.Description,
		Timeout:       data.Timeout,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新脚本

type UpdateParam struct {
	Id            uint
	UserId        uint
	Name          string
	CommandType   string
	Username      string
	WorkDirectory string
	Content       string
	Description   string
	Timeout       uint
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Script{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Updates(dborm.Script{
			Name:          data.Name,
			CommandType:   data.CommandType,
			Username:      data.Username,
			WorkDirectory: data.WorkDirectory,
			Content:       data.Content,
			Description:   data.Description,
			Timeout:       data.Timeout,
		})

	return result.Error

}

// 删除脚本

type DeleteParam struct {
	Id     uint
	UserId uint
}

func Delete(data *DeleteParam) error {

	result := dborm.Db.
		Where(&dborm.Script{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&dborm.Script{})

	return result.Error

}

// 获取脚本

type FetchParam struct {
	Id     uint
	UserId uint
}

func Fetch(data *FetchParam) (*dborm.Script, error) {

	var item *dborm.Script

	result := dborm.Db.
		Where(&dborm.Script{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		First(&item)

	return item, result.Error

}

// 获取脚本列表

type FetchAllParam struct {
	UserId      uint
	CommandType string
}

func FetchAll(data *FetchAllParam) ([]*dborm.Script, error) {

	var items []*dborm.Script

	result := dborm.Db.
		Where(&dborm.Script{
			UserId:      data.UserId,
			CommandType: data.CommandType,
		}).
		Find(&items)

	return items, result.Error

}

// 获取脚本总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&dborm.Script{}).
		Where(&dborm.Script{
			UserId:      data.UserId,
			CommandType: data.CommandType,
		}).
		Count(&count)

	return count, result.Error

}

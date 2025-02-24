package machine

import (
	"tdp-cloud/module/dborm"
)

// 创建主机

type CreateParam struct {
	UserId      uint
	VendorId    uint
	HostName    string `binding:"required"`
	IpAddress   string `binding:"required"`
	OSType      string `binding:"required"`
	Region      string
	Model       string `binding:"required"`
	CloudId     string
	CloudMeta   any
	WorkerId    string
	WorkerMeta  any
	Description string
	Status      uint
}

func Create(data *CreateParam) (uint, error) {

	item := &dborm.Machine{
		UserId:      data.UserId,
		VendorId:    data.VendorId,
		HostName:    data.HostName,
		IpAddress:   data.IpAddress,
		OSType:      data.OSType,
		Region:      data.Region,
		Model:       data.Model,
		CloudId:     data.CloudId,
		CloudMeta:   data.CloudMeta,
		WorkerId:    data.WorkerId,
		WorkerMeta:  data.WorkerMeta,
		Description: data.Description,
		Status:      data.Status,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新主机

type UpdateParam struct {
	Id          uint
	UserId      uint
	VendorId    uint
	HostName    string
	IpAddress   string
	OSType      string
	Region      string
	Model       string
	CloudId     string
	CloudMeta   any
	WorkerId    string
	WorkerMeta  any
	Description string
	Status      uint
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Machine{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Updates(dborm.Machine{
			VendorId:    data.VendorId,
			HostName:    data.HostName,
			IpAddress:   data.IpAddress,
			OSType:      data.OSType,
			Region:      data.Region,
			Model:       data.Model,
			CloudId:     data.CloudId,
			CloudMeta:   data.CloudMeta,
			WorkerId:    data.WorkerId,
			WorkerMeta:  data.WorkerMeta,
			Description: data.Description,
			Status:      data.Status,
		})

	return result.Error

}

// 删除主机

type DeleteParam struct {
	Id     uint
	UserId uint
}

func Delete(data *DeleteParam) error {

	result := dborm.Db.
		Where(&dborm.Machine{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&dborm.Machine{})

	return result.Error

}

// 获取主机

type FetchParam struct {
	Id     uint
	UserId uint
}

func Fetch(data *FetchParam) (*dborm.Machine, error) {

	var item *dborm.Machine

	result := dborm.Db.
		Where(&dborm.Machine{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		First(&item)

	return item, result.Error

}

// 获取主机列表

type FetchAllParam struct {
	UserId   uint
	VendorId uint
}

func FetchAll(data *FetchAllParam) ([]*dborm.Machine, error) {

	var items []*dborm.Machine

	result := dborm.Db.
		Where(&dborm.Machine{
			UserId:   data.UserId,
			VendorId: data.VendorId,
		}).
		Find(&items)

	return items, result.Error

}

// 获取主机总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&dborm.Machine{}).
		Where(&dborm.Machine{
			UserId:   data.UserId,
			VendorId: data.VendorId,
		}).
		Count(&count)

	return count, result.Error

}

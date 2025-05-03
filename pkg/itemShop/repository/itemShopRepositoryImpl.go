package repository

import (
	"github.com/JARNBOY/jb-isekai-shop-tutorial/entities"
	_itemShopException "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/exception"
	_itemShopModel "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ItemShopRepositoryImpl struct{
	db *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	
	return &ItemShopRepositoryImpl{db, logger}
}

func (r *ItemShopRepositoryImpl) Listing(itemFiler *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	itemList := make ([]*entities.Item, 0) // itemList is a slice of pointers

	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false) // select * from items
	if itemFiler.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFiler.Name+"%")
	}
	if itemFiler.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFiler.Description+"%")
	}
	
	offset := int((itemFiler.Paginate.Page - 1) * itemFiler.Paginate.Size)
	limit := int(itemFiler.Paginate.Size)

	if err := query.Offset(offset).Limit(limit).Find(&itemList).Error; err != nil {
		r.logger.Errorf("Failed to list items: %s", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}

	return itemList, nil
}

func (r *ItemShopRepositoryImpl) Counting(itemFiler *_itemShopModel.ItemFilter) (int64, error) {
	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false) // select * from items
	if itemFiler.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFiler.Name+"%")
	}
	if itemFiler.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFiler.Description+"%")
	}

	var count int64 // แบบpointer ปกติ count := new(int64) แต่ประกาศแบบนี้ var count int64 ช่วยป้องกันเวลาที่ nil จะได้ไม่ panic แบบปกติที่จะ new ->  
	
	if err := query.Count(&count).Error; err != nil {
		r.logger.Errorf("Counting Items Failed: %s", err.Error())
		return -1, &_itemShopException.ItemCounting{}
	}

	return count, nil
}
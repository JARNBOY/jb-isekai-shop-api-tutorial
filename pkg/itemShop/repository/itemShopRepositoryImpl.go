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

	query := r.db.Model(&entities.Item{}) // select * from items
	if itemFiler.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFiler.Name+"%")
	}
	if itemFiler.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFiler.Description+"%")
	}

	if err := query.Find(&itemList).Error; err != nil {
		r.logger.Errorf("Failed to list items: %s", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}

	return itemList, nil
}
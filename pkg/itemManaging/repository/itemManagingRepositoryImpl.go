package repository

import (
	"github.com/JARNBOY/jb-isekai-shop-tutorial/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	_itemManagingException "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/exception"
	_itemManagingModel "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/model" 
)

type itemManagingRepositoryImpl struct {
	db			*gorm.DB
	logger		echo.Logger
}

func NewItemManagingRepositoryImpl(db *gorm.DB, logger echo.Logger) *itemManagingRepositoryImpl {
	return &itemManagingRepositoryImpl{
		db: db,
		logger: logger,
	}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Error("Creating Item Fail: %s", err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}

	return itemEntity, nil
}

func (r *itemManagingRepositoryImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {
	
	if err := r.db.Model(&entities.Item{}).Where(
		"id = ?", itemID,
	).Updates(
		itemEditingReq,
	).Error; err != nil {
		r.logger.Error("Editing Item Fail: %s", err.Error())
		return 0, &_itemManagingException.ItemEditing{ItemID: itemID}
	}
	
	return itemID, nil
}
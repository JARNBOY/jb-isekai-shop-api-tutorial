package repository

import (
	"github.com/JARNBOY/jb-isekai-shop-tutorial/databases"
	"github.com/JARNBOY/jb-isekai-shop-tutorial/entities"
	_itemManagingException "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/exception"
	_itemManagingModel "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/model"
	"github.com/labstack/echo/v4"
)

type itemManagingRepositoryImpl struct {
	db			databases.Database
	logger		echo.Logger
}

func NewItemManagingRepositoryImpl(db databases.Database, logger echo.Logger) *itemManagingRepositoryImpl {
	return &itemManagingRepositoryImpl{
		db: db,
		logger: logger,
	}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Connect().Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Error("Creating Item Fail: %s", err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}

	return itemEntity, nil
}

func (r *itemManagingRepositoryImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {
	
	if err := r.db.Connect().Model(&entities.Item{}).Where(
		"id = ?", itemID,
	).Updates(
		itemEditingReq,
	).Error; err != nil {
		r.logger.Error("Editing Item Fail: %s", err.Error())
		return 0, &_itemManagingException.ItemEditing{ItemID: itemID}
	}
	
	return itemID, nil
}

func (r *itemManagingRepositoryImpl) Archiving(itemID uint64) error {
	if err := r.db.Connect().Table(
		"items",
		).Where(
			"id = ?", itemID,
		).Update(
			"is_archive", true,
		).Error; err != nil {
		r.logger.Error("Archiving Item Fail: %s", err.Error())
		return &_itemManagingException.ItemArchiving{ItemID: itemID}
	}

	return nil
}
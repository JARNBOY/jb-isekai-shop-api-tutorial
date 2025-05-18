package repository

import "github.com/JARNBOY/jb-isekai-shop-tutorial/entities"

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
}
package server

import (
	_itemShopRepository "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/repository"
	_itemShopService "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/service"
	_itemShopController "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/controller"
)

func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopRepositoryImpl(itemShopRepository)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("", itemShopController.Listing)
}

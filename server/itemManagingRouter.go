package server

import (
	_itemManagingRepository "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/repository"
	_itemManagingService "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/service"
	_itemManagingController "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/controller"
	_itemShopRepository "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/repository"
)

func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemManagingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)

	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(
		itemManagingRepository, 
		itemShopRepository,
	)

	itemManagingController := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemManagingController.Creating)
	router.PATCH("/:itemID", itemManagingController.Editing)
	router.DELETE("/:itemID", itemManagingController.Archiving)
}
package server

import (
	_itemManagingRepository "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/repository"
	_itemManagingService "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/service"
	_itemManagingController "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/controller"
)

func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemManagingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemManagingRepository)
	itemManagingController := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	_ = itemManagingController
	_ = router
}
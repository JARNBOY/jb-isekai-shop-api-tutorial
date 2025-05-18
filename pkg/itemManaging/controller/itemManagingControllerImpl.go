package controller

import (
	"net/http"

	"github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/custom"
	_itemManagingModel "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/model"
	_itemManagingService "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/service"
	"github.com/labstack/echo/v4"
)

type itemManagingControllerImpl struct {
	itemManagingService _itemManagingService.ItemManagingService
}

func NewItemManagingControllerImpl(
	itemManagingService  _itemManagingService.ItemManagingService,
) itemManagingController {
	return &itemManagingControllerImpl{
		itemManagingService: itemManagingService,
	}
}

func (c *itemManagingControllerImpl) Creating(pctx echo.Context) error {
	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemCreatingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	item, err := c.itemManagingService.Creating(itemCreatingReq)
	if err != nil {
		return err
	}

	return pctx.JSON(http.StatusCreated, item)
}
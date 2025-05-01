package controller

import "github.com/labstack/echo/v4"

type itemShopController interface{
	Listing(pctx echo.Context) error
}
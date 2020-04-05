package app

import (
	"github.com/labstack/echo/v4"
)

type Application struct {
	handlers map[string]echo.HandlerFunc
}

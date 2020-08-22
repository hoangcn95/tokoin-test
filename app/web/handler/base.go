package handler

import (
	"app/common/config"
	ghandler "app/common/gstuff/handler"
)

var cfg = config.GetConfig()

func success(data interface{}) (int, interface{}) {
	return ghandler.Success(data)
}

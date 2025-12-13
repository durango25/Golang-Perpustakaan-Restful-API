package buku

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/ports"
)

type Controller struct {
	service ports.IBukuService
}

func NewBukuController(service ports.IBukuService) *Controller {
	return &Controller{
		service,
	}
}

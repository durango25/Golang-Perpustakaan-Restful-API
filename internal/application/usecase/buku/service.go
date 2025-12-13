package buku

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/ports"

type BukuService struct {
	Repository ports.IBukuRepository
}

func NewBukuService(Repository ports.IBukuRepository) *BukuService {
	return &BukuService{
		Repository,
	}
}

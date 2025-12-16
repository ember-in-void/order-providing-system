package usecase

import "order-providing-system/internal/storage"

type CustomService struct {
	NewRepo storage.RepoModule
}

func NewUsecase(repo storage.RepoModule) *CustomService {
	return &CustomService{NewRepo: repo}
}

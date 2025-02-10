package usecase

import "frappuccino/internal/storage"

type CustomService struct {
	NewRepo storage.RepoModule
}

func NewUsecase(repo storage.RepoModule) *CustomService {
	return &CustomService{NewRepo: repo}
}

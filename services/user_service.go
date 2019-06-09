package services

import "DuyStifler/GolangServer/manager"

type UserService struct {
	serverManager *manager.Manager
}

func NewUserService(serverManager *manager.Manager) *UserService {
	return &UserService{
		serverManager:serverManager,
	}
}
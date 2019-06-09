package manager

import (
	"DuyStifler/GolangServer/cache"
	"DuyStifler/GolangServer/cache/caos"
	"DuyStifler/GolangServer/database/daos"
	"DuyStifler/GolangServer/services"
)

type CacheManager struct {
	sessionCAO *caos.SessionCAO
	userCAO    *caos.UserCAO
}

type DaoManager struct {
	userDAO *daos.UserDAO
}

type ServiceManager struct {
	userService *services.UserService
}

type Manager struct {
	cacheManager *CacheManager
	daoManager *DaoManager
	serviceManager *ServiceManager

	cache *cache.Cache
}

func (m *Manager) ServiceManager() *ServiceManager {
	return m.serviceManager
}

func (m *Manager) DaoManager() *DaoManager {
	return m.daoManager
}

func (m *Manager) CacheManager() *CacheManager {
	return m.cacheManager
}


func (c CacheManager) UserCAO() *caos.UserCAO {
	return c.userCAO
}

func (c CacheManager) SessionCAO() *caos.SessionCAO {
	return c.sessionCAO
}


func (d DaoManager) UserDAO() *daos.UserDAO {
	return d.userDAO
}

func (s ServiceManager) UserService() *services.UserService {
	return s.userService
}
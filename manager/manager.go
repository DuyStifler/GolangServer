package manager

import (
	"DuyStifler/GolangServer/cache"
	"DuyStifler/GolangServer/cache/caos"
	"DuyStifler/GolangServer/database/daos"
	"DuyStifler/GolangServer/services"
)

var (
	cacheManager = &CacheManager{}
	daoManager = &DaoManager{}
	serviceManger = &ServiceManager{}
)

func NewManager(cache *cache.Cache) *Manager {
	return &Manager{
		cacheManager,
		daoManager,
		serviceManger,
		cache,
	}
}

func (m *Manager) Init() {
	m.InitCAO()
	m.InitDAO()
	m.InitService()
}

func (m *Manager) InitCAO() {
	m.CacheManager().userCAO = caos.NewUserCAO(m.cache)
	m.CacheManager().sessionCAO = caos.NewSessionCAO(m.cache)
}

func (m *Manager) InitDAO() {
	m.DaoManager().userDAO = daos.NewUserDAO()
}

func (m *Manager) InitService() {
	m.ServiceManager().userService = services.NewUserService(m)
}
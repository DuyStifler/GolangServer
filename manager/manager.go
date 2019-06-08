package manager

var (
	daoManager = &DaoManager{}
	serviceManger = &ServiceManager{}
)

func NewManager() *Manager {
	return &Manager{
		daoManager,
		serviceManger,
	}
}
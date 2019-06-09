package http_server

import (
	serverCache "DuyStifler/GolangServer/cache"
	serverDatabase "DuyStifler/GolangServer/database"
	"DuyStifler/GolangServer/models"
	"DuyStifler/GolangServer/utils"
	"fmt"

	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

const (
	SCOPE_FIELD = "user_connector"
)

type HttpServer struct {
	logger   *utils.Logger
	database *serverDatabase.Database
	cache    *serverCache.Cache
	config   *models.ServerConfig

	e *echo.Echo
}

func (a *HttpServer) Cache() *serverCache.Cache {
	return a.cache
}

func (a *HttpServer) Logger() *utils.Logger {
	return a.logger
}

func (a *HttpServer) Database() *serverDatabase.Database {
	return a.database
}

func (a *HttpServer) E() *echo.Echo {
	return a.e
}

func NewHttpServer(serverConfig *models.ServerConfig) *HttpServer {
	var err error
	httpServer := &HttpServer{}

	httpServer.config = serverConfig

	httpServer.logger, err = utils.NewLogger(httpServer.config.Log.LogErrorDir, httpServer.config.Log.LogInfoDir)
	if err != nil {
		return nil
	}

	httpServer.database, err = serverDatabase.NewDatabase(&httpServer.config.Database, httpServer.logger)
	if err != nil {
		return nil
	}

	httpServer.setupMiddleWare(httpServer.e, httpServer.database, httpServer.cache)

	return nil
}

func (a *HttpServer) setupMiddleWare(e *echo.Echo, database *serverDatabase.Database, cache *serverCache.Cache) {
	e.Use(TransactionHandler(database, cache))
	e.Use(echoMw.Gzip())
	e.Use(echoMw.Recover())
	e.Use(echoMw.Logger())
	e.Use(echoMw.CORS())

	e.Use(TransactionHandler(database, cache))
	//e.Use(MiddlewareAuthConfig(model.MiddlewareAuthConfig{
	//	Skipper: func(context echo.Context) bool {
	//		if context.Path() == "/api/login" {
	//			return true
	//		}
	//		return false
	//	},
	//}))
}

//func MiddlewareAuthConfig(config model.MiddlewareAuthConfig) echo.MiddlewareFunc {
//	if config.Skipper == nil {
//		config.Skipper = func(c echo.Context) bool {
//			if c.Request().URL.Path == "/api/login" {
//				return true
//			}
//			return false
//		}
//	}
//
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return echo.HandlerFunc(func(c echo.Context) error {
//			if c.Request().Header.Get(keys.VERSION_KEY) != config.Version {
//				return errors.New("version wrong")
//			}
//
//			//TODO
//			//de tam the nay, sau nay xem the nao se sua lai cho vao body sau
//			token := c.Request().Header.Get("Authorization")
//			if token == "" {
//				token = c.QueryParam("token")
//			}
//
//			appHandler := c.Get("app_handler").(handler.AppHandler)
//			userSession, err := appHandler.Cache.GetUserSession(token)
//			if err != nil {
//				return err
//			}
//			if userSession.UserID != "" {
//				appHandler.UserID = userSession.UserID
//				err = appHandler.Cache.UpdateSession(token)
//				if err != nil {
//					return err
//				}
//				return next(c)
//			} else {
//				//dont have a session for this token
//				if config.Skipper(c) {
//					//create token
//					err = appHandler.Cache.CreateSession(appHandler.UserID)
//					if err != nil {
//						return err
//					}
//					return next(c)
//				} else {
//					return echo.ErrUnauthorized
//				}
//			}
//		})
//	}
//}

func (a *HttpServer) GetScope(c echo.Context) *models.UserConnector {
	scope := c.Get(SCOPE_FIELD).(models.UserConnector)
	return &scope
}

func (a *HttpServer) Run() {
	a.e.Logger.Fatal(a.e.Start(fmt.Sprintf(":%d", a.config.Port)))
}

func TransactionHandler(database *serverDatabase.Database, cache *serverCache.Cache) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			userConnector := &models.UserConnector{}

			masterTx, err := database.DbConnector().MasterDB().Begin()
			if err != nil {
				return err
			}
			userConnector.SetMasterTx(masterTx)

			replicaTx, err := database.DbConnector().SlaveDB().Begin()
			if err != nil {
				return nil
			}
			userConnector.SetReplicaTx(replicaTx)

			c.Set(SCOPE_FIELD, userConnector)

			if err := next(c); err != nil {
				//get doesnt need rollback because it will not change data
				_ = masterTx.Rollback()
			}

			_ = masterTx.Commit()
			_ = replicaTx.Commit()

			return nil
		})
	}
}

func AuthWIthConfig(serverConfig models.ServerConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func (c echo.Context) error {
			if !utils.InStringArray(c.Path(), serverConfig.SkipPaths) {

			}

			return next(c)
		})
	}
}
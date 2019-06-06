package http_server

import (
	"os"

	serverCache "DuyStifler/GolangServer/cache"
	serverDatabase "DuyStifler/GolangServer/database"
	"DuyStifler/GolangServer/models"
	"DuyStifler/GolangServer/utils"

	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

type HttpServer struct {
	logger   *utils.Logger
	database *serverDatabase.Database
	cache    *serverCache.Cache
	config   *models.ServerConfig

	e *echo.Echo
}

func NewHttpServer() *HttpServer {
	httpServer := &HttpServer{}

	err := httpServer.generateConfig()
	if err != nil {
		return nil
	}

	httpServer.logger, err = utils.NewLogger(httpServer.config.Log.LogErrorDir, httpServer.config.Log.LogInfoDir)
	if err != nil {
		return nil
	}

	httpServer.database, err = serverDatabase.NewDatabase(&httpServer.config.Database, httpServer.logger)
	if err != nil {
		return nil
	}

	return nil
}

func (a *HttpServer) generateConfig() error {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer func() {
		jsonFile.Close()
	}()

	a.config, err = models.NewServerConfig(jsonFile)
	if err != nil {
		return err
	}

	return nil
}


func setupMiddleWare(e *echo.Echo, database *serverDatabase.Database, cache serverCache.Cache, rrCount *int64) {
	e.Use(TransactionHandler(database, cache))
	e.Use(echoMw.Gzip())
	e.Use(echoMw.Recover())
	e.Use(echoMw.Logger())
	e.Use(echoMw.CORS())
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

func TransactionHandler(database *serverDatabase.Database, cache serverCache.Cache) echo.MiddlewareFunc {
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

			c.Set("user_connector", userConnector)

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

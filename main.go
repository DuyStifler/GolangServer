package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	serverCache "DuyStifler/GolangServer/cache"
	"DuyStifler/GolangServer/handler"
	"DuyStifler/GolangServer/keys"
	"DuyStifler/GolangServer/model"
	"DuyStifler/GolangServer/utils"

	"github.com/gocraft/dbr"
	"github.com/google/logger"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

var rrCount int64 = 0

func init()  {
	//get all param when run server
	flag.Parse()
	//create seed for random
	rand.Seed(time.Now().UnixNano())
}

func main() {
	loggerFile, err := setupLogger()
	if err != nil {
		return
	}
	defer loggerFile.Close()
	defer logger.Init("LoggerServer", true, true, loggerFile).Close()

	cacheConfig, db, err := utils.GetConfig()
	masterSession := db.InitDatabaseMaster()
	replicasSession := db.GenerateSlave()

	cache, err := serverCache.NewRedisCache(cacheConfig)
	if err != nil {
		logger.Error("[ERROR - Generate Cache] generate new redis cache ", err)
		return
	}

	e := echo.New()
	setupMiddleWare(e, masterSession, replicasSession, cache, &rrCount)
}

func setupMiddleWare(e *echo.Echo, masterSession *dbr.Session, replicasSession []*dbr.Session, cache serverCache.Cache, rrCount *int64) {
	e.Use(TransactionHandler(masterSession, replicasSession, cache, rrCount))
	e.Use(echoMw.Gzip())
	e.Use(echoMw.Recover())
	e.Use(echoMw.Logger())
	e.Use(echoMw.CORS())
	e.Use(MiddlewareAuthConfig(model.MiddlewareAuthConfig{
		Skipper: func(context echo.Context) bool {
			if context.Path() == "/api/login" {
				return true
			}
			return false
		},
	}))
}

func MiddlewareAuthConfig(config model.MiddlewareAuthConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = func(c echo.Context) bool {
			//if strings.HasSuffix(c.Request().URL.String(), "login") {
			//	return true
			//}
			if c.Request().URL.Path == "/api/login" {
				return true
			}
			return false
		}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			//TODO
			//de tam the nay, sau nay xem the nao se sua lai cho vao body sau
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				token = c.QueryParam("token")
			}

			appHandler := c.Get("app_handler").(handler.AppHandler)
			userSession, err := appHandler.Cache.GetUserSession(token)
			if err != nil {
				return err
			}
			if userSession.UserID != "" {
				appHandler.UserID = userSession.UserID
				err = appHandler.Cache.UpdateSession(token)
				if err != nil {
					return err
				}
				return next(c)
			} else {
				//dont have a session for this token
				if config.Skipper(c) {
					//create token
					err = appHandler.Cache.CreateSession(appHandler.UserID)
					if err != nil {
						return err
					}
					return next(c)
				} else {
					return echo.ErrUnauthorized
				}
			}
		})
	}
}

func TransactionHandler(masterSession *dbr.Session, replicasSession []*dbr.Session, cache serverCache.Cache, rrCount *int64) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			appHandler := handler.AppHandler{
				Cache: cache,
			}

			tx, err := masterSession.Begin()
			if err != nil {
				logger.Error("[ERROR - transaction handler] set master tx ", err)
			}
			appHandler.SetMasterTx(tx)
			appHandler.SetReplicaTx(replicasSession, rrCount, len(replicasSession))

			c.Set("app_handler", appHandler)

			if err := next(c); err != nil {
				//get doesnt need rollback because it will not change data
				_ = appHandler.GetMasterTx().Rollback()
			}

			_ = appHandler.GetReplicaTx().Commit()
			_ = appHandler.GetMasterTx().Commit()

			return nil
		})
	}
}

func setupLogger() (*os.File, error) {
	//setup logger
	logger.SetFlags(log.Llongfile)
	loggerFile, err := os.OpenFile(keys.LOGGER_FILE_PATH, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

	if err != nil {
		logger.Error("[ERROR - main] create logger file")
		return nil, err
	}

	return loggerFile, nil
}

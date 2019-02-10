package handler

import (
	"DuyStifler/GolangServer/cache"
	"github.com/gocraft/dbr"
	"github.com/google/logger"
)

type AppHandler struct {
	Logger logger.Logger
	Cache  cache.Cache
	UserID string

	tx         *dbr.Tx // the currently active transaction
	sess       *dbr.Session
}

func(a *AppHandler) SetTx(tx *dbr.Tx) {
	a.tx = tx
}

func(a *AppHandler) GetTx() *dbr.Tx {
	return a.tx
}
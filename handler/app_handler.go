package handler

import (
	"sync/atomic"

	"DuyStifler/GolangServer/cache"
	"github.com/gocraft/dbr"
)

type AppHandler struct {
	Cache  cache.Cache
	UserID string

	masterTx   *dbr.Tx // the currently active transaction
	replicaTx *dbr.Tx
	sess       *dbr.Session
}

func (a *AppHandler) SetMasterTx(tx *dbr.Tx) {
	a.masterTx = tx
}

func (a *AppHandler) GetMasterTx() *dbr.Tx {
	return a.masterTx
}

func (a *AppHandler) SetReplicaTx(arr []*dbr.Session, rrCounter *int64, remainCount int) {
	remainCount--
	if remainCount == 0 {
		a.replicaTx = a.masterTx
		return
	}

	rrNum := atomic.AddInt64(rrCounter, 1) % int64(len(arr))
	tx, err := arr[rrNum].Begin()
	if err != nil {
		atomic.AddInt64(rrCounter, 1)
		a.SetReplicaTx(arr, rrCounter, remainCount)
	} else {
		a.replicaTx = tx
	}
}

func (a *AppHandler) GetReplicaTx() *dbr.Tx {
	return a.replicaTx
}
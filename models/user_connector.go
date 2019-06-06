package models

import "github.com/gocraft/dbr"

type UserConnector struct {
	userID string

	masterTx *dbr.Tx
	replicaTx *dbr.Tx
}

func (u *UserConnector) ReplicaTx() *dbr.Tx {
	return u.replicaTx
}

func (u *UserConnector) SetReplicaTx(replicaTx *dbr.Tx) {
	u.replicaTx = replicaTx
}

func (u *UserConnector) MasterTx() *dbr.Tx {
	return u.masterTx
}

func (u *UserConnector) SetMasterTx(masterTx *dbr.Tx) {
	u.masterTx = masterTx
}

func (u *UserConnector) SetUserID(userID string) {
	u.userID = userID
}

func (u *UserConnector) UserID() string {
	return u.userID
}
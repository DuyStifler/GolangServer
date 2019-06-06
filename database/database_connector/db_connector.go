package database_connector

import (
	"sync"

	"github.com/gocraft/dbr"
)

type DatabaseConnector struct {
	mySQLUser       string
	mySQLPassword   string
	mySQLPort       string
	mySQLMasterURL  string
	mySQLReplicaURL []string
	mySQLDatabase   string

	masterDb *dbr.Session
	slavesDb []*dbr.Session

	counterRequest int
	mux            *sync.Mutex
}

func NewDatabaseConnector(mySQLUser string, mySQLPassword string, mySQLPort string, mySQLMasterURL string, mySQLReplicaURL []string, mySQLDatabase string) (*DatabaseConnector, error) {
	db := &DatabaseConnector{mySQLUser: mySQLUser, mySQLPassword: mySQLPassword, mySQLPort: mySQLPort, mySQLMasterURL: mySQLMasterURL,
		mySQLReplicaURL: mySQLReplicaURL, mySQLDatabase: mySQLDatabase, mux: &sync.Mutex{}}

	var err error
	db.masterDb, err = db.initDatabaseMaster()
	if err != nil {
		return nil, err
	}

	db.slavesDb = db.generateSlaver()

	return db, nil
}

func (d *DatabaseConnector) MasterDB() *dbr.Session {
	return d.masterDb
}

func (d *DatabaseConnector) SlaveDB() *dbr.Session {
	if len(d.slavesDb) == 0 {
		return nil
	}

	if len(d.slavesDb) == 1 {
		return d.slavesDb[0]
	}

	d.mux.Lock()
	if d.counterRequest == len(d.slavesDb) {
		d.counterRequest = 1
	} else {
		d.counterRequest++
	}
	d.mux.Unlock()

	return d.slavesDb[d.counterRequest - 1]
}

func (d *DatabaseConnector) initDatabaseMaster() (*dbr.Session, error) {
	db, err := dbr.Open("mysql",
		d.mySQLUser+":"+d.mySQLPassword+"@tcp("+d.mySQLMasterURL+":"+d.mySQLPort+")/"+d.mySQLDatabase+"?parseTime=true&charset=utf8mb4,utf8",
		nil)

	if err != nil {
		return nil, err
	}

	return db.NewSession(nil), err
}

func (d *DatabaseConnector) generateSlaver() ([]*dbr.Session) {
	var sessions []*dbr.Session

	for _, url := range d.mySQLReplicaURL {
		db, err := dbr.Open("mysql",
			d.mySQLUser+":"+d.mySQLPassword+"@tcp("+url+":"+d.mySQLPort+")/"+d.mySQLDatabase+"?parseTime=true&charset=utf8mb4,utf8",
			nil)

		if err != nil {
			continue
		}

		sessions = append(sessions, db.NewSession(nil))
	}

	if len(sessions) == 0 {
		sessions = append(sessions, d.masterDb)
	}

	return sessions
}

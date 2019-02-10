package database

import (
	"github.com/gocraft/dbr"
	"log"
)

type Database struct {
	MySQLUser       string
	MySQLPassword   string
	MySQLPort       string
	MySQLMasterURL  string
	MySQLReplicaURL []string
	MySQLDatabase   string
}

func (d *Database) InitDatabaseMaster() *dbr.Session {
	db, err := dbr.Open("mysql",
		d.MySQLUser+":"+d.MySQLPassword+"@tcp("+d.MySQLMasterURL+":"+d.MySQLPort+")/"+d.MySQLDatabase+"?parseTime=true&charset=utf8mb4,utf8",
		nil)

	if err != nil {
		log.Println(err)
		return nil
	}
	return db.NewSession(nil)
}

func (d *Database) GenerateSlave() []*dbr.Session {
	var sessions []*dbr.Session

	for _, url := range d.MySQLReplicaURL {
		db, err := dbr.Open("mysql",
			d.MySQLUser+":"+d.MySQLPassword+"@tcp("+url+":"+d.MySQLPort+")/"+d.MySQLDatabase+"?parseTime=true&charset=utf8mb4,utf8",
			nil)

		if err != nil {
			log.Println(err)
			continue
		}

		sessions = append(sessions, db.NewSession(nil))
	}

	if len(sessions) == 0 {
		sessions = append(sessions, d.InitDatabaseMaster())
	}

	return sessions
}

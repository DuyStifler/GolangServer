package database

import (
	"DuyStifler/GolangServer/database/database_connector"
	"DuyStifler/GolangServer/models"
	"DuyStifler/GolangServer/utils"
)

type Database struct {
	dbConnector *database_connector.DatabaseConnector
	logger      *utils.Logger
}

func (d *Database) DbConnector() *database_connector.DatabaseConnector {
	return d.dbConnector
}

func NewDatabase(dbConfig *models.DatabaseConfig, serverLog *utils.Logger) (*Database, error) {
	dbConnection, err := database_connector.NewDatabaseConnector(dbConfig.UserName, dbConfig.Password, dbConfig.Port, dbConfig.UrlMaster,
		dbConfig.UrlReplicas, dbConfig.Schema)

	if err != nil {
		return nil, err
	}


	return &Database{dbConnection, serverLog}, nil
}

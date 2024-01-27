package db

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Driver de conexión con Postgres
)

// Conecta a base de dados
func EstablishDbConnection() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.GlobalConfig.DbData.Host, config.GlobalConfig.DbData.Port, config.GlobalConfig.DbData.RolName, config.GlobalConfig.DbData.RolPassword, config.GlobalConfig.DbData.DbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("Conectado com sucesso à base de dados")
	return db, nil
}

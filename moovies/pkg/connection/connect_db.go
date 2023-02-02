package connection

import (
	"database/sql"
	"fmt"
)

func ConnectToDB(postgres Postgres) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		postgres.Host,
		postgres.Port,
		postgres.User,
		postgres.Password,
		postgres.DBName,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}

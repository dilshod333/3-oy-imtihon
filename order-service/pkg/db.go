package pkg

import (
	"database/sql"
	"log"
)


func OpenSql(url string )*sql.DB{
	db, err := sql.Open("postgres", url) 

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err !=  nil {
		log.Fatal(err)
	}

	return db 

}

// import (
// 	"database/sql"
// 	"fmt"
// 	"order-service/config"

// 	_ "github.com/lib/pq"
// )

// func ConnectToDBForSuit(cfg config.Config) (*sql.DB, func()) {
// 	psqlString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		cfg.Postgres.Host,
// 		cfg.Postgres.Port,
// 		cfg.Postgres.User,
// 		cfg.Postgres.Password,
// 		cfg.Postgres.Database,
// 	)

// 	connDb, err := sql.Open("postgres", psqlString)
// 	if err != nil {
// 		return nil, func() {}
// 	}

// 	cleanUpfunc := func() {
// 		connDb.Close()
// 	}

// 	return connDb, cleanUpfunc
// }

// func ConnectToDB(cfg config.Config) (*sql.DB, error) {
// 	psqlString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		cfg.Postgres.Host,
// 		cfg.Postgres.Port,
// 		cfg.Postgres.User,
// 		cfg.Postgres.Password,
// 		cfg.Postgres.Database,
// 	)

// 	connDb, err := sql.Open("postgres", psqlString)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return connDb, nil
// }

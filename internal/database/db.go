package db

import (
	"basiccrud/internal/config"

	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPGDB(cnf config.Config) *gorm.DB {
	fmt.Println(cnf)
	postgresURL := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		cnf.PGHost, cnf.PGUserName, cnf.PGPassword, cnf.PGDBName, cnf.PgPort, cnf.PgSSLMode)

	db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to postgres db: ", err)
	}

	log.Println("connected to postgres db successfully!")

	return db
}

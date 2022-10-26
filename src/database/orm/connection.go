package orm

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func New() (*gorm.DB, error) {

	DBUSER := os.Getenv("DBUSER")
	DBPASSWORD := os.Getenv("DBPASSWORD")
	DBNAME := os.Getenv("DBNAME")
	DBHOST := os.Getenv("DBHOST")
	DBPORT := os.Getenv("DBPORT")
	SSLMODE := os.Getenv("SSLMODE")
	DBTIMEZONE := os.Getenv("DBTIMEZONE")

	conn := "host=" + DBHOST +
		" user=" + DBUSER +
		" password=" + DBPASSWORD +
		" dbname=" + DBNAME +
		" sslmode=" + SSLMODE +
		" port=" + DBPORT +
		" TimeZone=" + DBTIMEZONE

	var err error
	gormDb, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println("Connection Success")
	}

	db, err := gormDb.DB()
	if err != nil {
		log.Fatalln(err.Error())
	}
	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)

	return gormDb, nil
}

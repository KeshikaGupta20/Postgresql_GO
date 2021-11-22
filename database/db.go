package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/KeshikaGupta20/Postgresql_GO/models"

	"github.com/jinzhu/gorm"
)

// DB gorm connector
var DB *gorm.DB

// connection with database
func ConnectDB() {
	var err error
	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	DB, err = gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		port, os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	//for migrates the struct in the database
	DB.AutoMigrate(&models.Employee{})
	fmt.Println("Database Migrated")
}

/* import (
	"fmt"

	"github.com/KeshikaGupta20/Postgresql_GO/models"
	"github.com/jinzhu/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "keshika"
	dbname   = "godb"
)

var DB *gorm.DB

func Database() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Successfully connected!")

	db.AutoMigrate(&models.Employee{})

} */

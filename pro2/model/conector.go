package model

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Con *sql.DB = nil

func Conectar() bool {
	error := godotenv.Load()
	if error != nil {
		panic(error)
	}

	bd, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+
		""+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))

	if err != nil {
		log.Fatal(err)
		return false
	}

	Con = bd
	return true

}

func Close() {
	if Con != nil {
		Con.Close()
	}
}

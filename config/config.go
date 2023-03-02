package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	DataSourceName string
}

func (c *Config) initDb() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	// tampung nilai ENV dari terminal
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// data source name
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	c.DataSourceName = dsn
}

func NewConfig() Config {
	config := Config{}
	config.initDb() // untuk menjalankan method yang didalamnya membuka koneksi ke DB
	return config   // Mengirimkan Object Config yang didalamnya terdapat attribute bertipe data Koneksi, namun koneksi tidak dapat diakses secara
	// langsung karena attribute bertipe private
}

func GetConn() *sqlx.DB {
	var db *sqlx.DB
	return db
}

package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DataSourceName string
}

func (c *Config) initDb() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	} //

	// Tampung nilai ENV dari terminal
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	//dbDriver := os.Getenv("DB_DRIVER")
	dbName := os.Getenv("DB_NAME")

	// data source name
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	c.DataSourceName = dsn
}

func NewConfig() Config {
	config := Config{}
	config.initDb() // untuk menjalankan method yang didalmnya membuka koneksi ke DB
	return config   // Mengirimkan Object Config yang dalamnya terdapat attribute bertipe data koneksi, namun koneksi tidak dapat diakses secara
	// langsung karena attribute bertipe private
}

// Used for connecting to the database and returning a reusable database instance
package database

import (
	"fmt"
	"log"

	"github.com/Aadi022/Backend_Golang/internal/config" // Imports our application configuration.
	"gorm.io/driver/postgres"                           // PostgreSQL driver for GORM.
	"gorm.io/gorm"                                      // GORM ORM package.
)

// Connects to PostgreSQL and returns the DB instance.
func Connect(cfg config.Config) (*gorm.DB, error) { //*gorm.DB → The database object you'll use for all database operations
	dsn := fmt.Sprintf( //creates the PostgreSQL DSN (connection string)
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", //// DSN format expected by PostgreSQL
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	db, err := gorm.Open( //Opens a connection to PostgreSQL using GORM
		postgres.Open(dsn), //Creates the PostgreSQL driver using the DSN
		&gorm.Config{},     //GORM configuration (currently using default settings)
	)

	if err != nil { //checks if the connection has failed
		return nil, err //Returns no db object and the error
	}

	log.Println("Database connected successfully")

	return db, nil
}

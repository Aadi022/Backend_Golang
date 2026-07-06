//Migration File

package database

import (
	"github.com/Aadi022/Backend_Golang/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	//GORM interprets this as Ensure a table exists for the User model, if it exists then alter if required
	return db.AutoMigrate( //Calls GORM's AutoMigrate() method to automatically create or update database tables based on the provided Go structs.
		&model.User{}, //passes the changes to be done in User model (as a pointer)
	)
}

/*
What does AutoMigrate() do?

It compares your Go structs with the database and:

Creates missing tables.
Adds missing columns.
Creates indexes and constraints.

It does not delete columns or tables, making it safer than destructive migrations.
*/

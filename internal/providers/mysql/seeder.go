package mysql

import (
	"context"
	"log"
	"main/core/enums"
	"main/data/models"
	"sync"

	"gorm.io/gorm/clause"
)

// Seed is a function that seeds the database
// It creates the initial data if it does not exist
//
// Returns an error if any
func Seed() error {
	// Begin a transaction
	tx := Conn.Begin()

	// Defer the rollback
	defer func() {
		// Recover from panic
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Return an error if any
	if tx.Error != nil {
		log.Println("Failed to begin transaction")

		return tx.Error
	}

	// Create a context with a cancel function
	_, cancel := context.WithCancel(context.Background())

	// Defer the cancel function
	defer cancel()

	// Create a wait group and error
	var (
		wg  sync.WaitGroup
		err error
	)

	// Add a new wait group
	wg.Add(1)

	// Seed court types table
	go func() {
		defer wg.Done()

		courtTypes := make([]models.CourtType, 0)

		// Loop through the court types enum
		for i, v := range enums.CourtTypes() {
			courtTypes = append(courtTypes, models.CourtType{
				ID:   uint(i + 1),
				Type: v,
			})
		}

		// Create the court types
		e := tx.Clauses(clause.Insert{Modifier: "ignore"}).Create(courtTypes).Error

		if e != nil {
			err = e

			// Rollback the transaction
			tx.Rollback()

			cancel()
		}
	}()

	// Check if there is an error
	if err != nil {
		return err
	}

	// Return an error if any
	err = tx.Commit().Error

	return err
}

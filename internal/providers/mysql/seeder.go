package mysql

import (
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
		e := Conn.Clauses(clause.Insert{Modifier: "ignore"}).Create(courtTypes).Error

		if e != nil {
			log.Println("Failed to seed court types table: " + e.Error())

			err = e
		}
	}()

	wg.Wait()

	return err
}

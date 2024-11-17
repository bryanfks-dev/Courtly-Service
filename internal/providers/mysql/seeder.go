package mysql

import (
	"main/core/enums"
	"main/data/models"

	"gorm.io/gorm/clause"
)

// Seed is a function that seeds the database
// It creates the initial data if it does not exist
//
// Returns an error if any
func Seed() error {
	// Create a channel for errors
	errs := make(chan error)

	// Seed court types table
	go func() {
		courtTypes := make([]models.CourtType, 0)

		// Loop through the court types enum
		for courtType := enums.CourtType(0); courtType <= enums.Badminton; courtType++ {
			courtTypes = append(courtTypes, models.CourtType{
				Type: courtType.Label(),
			})
		}

		// Create the court types
		if err := Conn.Clauses(clause.Insert{Modifier: "ignore"}).Create(&courtTypes).Error; err != nil {
			errs <- err

			return
		}

		errs <- nil
	}()

	// Seed court types table
	go func() {
		paymentMethods := make([]models.PaymentMethod, 0)

		// Loop through the court types enum
		for paymentMethod := enums.PaymentMethod(0); enums.CourtType(paymentMethod) <= enums.Badminton; paymentMethod++ {
			paymentMethods = append(paymentMethods, models.PaymentMethod{
				Method: paymentMethod.Label(),
			})
		}

		// Create the court types
		if err := Conn.Clauses(clause.Insert{Modifier: "ignore"}).Create(&paymentMethods).Error; err != nil {
			errs <- err

			return
		}

		errs <- nil
	}()

	// Check if there is an error
	if err := <-errs; err != nil {
		return err
	}

	return nil
}

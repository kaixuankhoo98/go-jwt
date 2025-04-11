package initializers

import (
	"go-jwt/models"
	"os"
)

func SyncDatabase() {
	if os.Getenv("AUTOMIGRATE") == "auto" {
		DB.AutoMigrate(
			&models.User{},
			&models.Workout{},
			&models.Exercise{},
			&models.Set{},
			&models.ExerciseType{},
			&models.Group{},
		)
	}
}

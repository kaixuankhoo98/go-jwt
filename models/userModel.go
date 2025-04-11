package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email         string `gorm:"unique"`
	Password      string
	Workouts      []Workout      `gorm:"foreignKey:UserID"`
	ExerciseTypes []ExerciseType `gorm:"foreignKey:UserID"`
	Groups        []Group        `gorm:"foreignKey:UserID"`
}

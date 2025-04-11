package models

type Set struct {
	ID         uint `gorm:"primaryKey"`
	ExerciseID uint
	Exercise   Exercise `gorm:"foreignKey:ExerciseID"`
	Weight     *float64
	Reps       *int
	Duration   *int
}

type SetWeightResponse struct {
	ID     uint     `json:"id"`
	Weight *float64 `json:"weight"`
	Reps   *int     `json:"reps"`
}

type SetDurationResponse struct {
	ID       uint `json:"id"`
	Duration *int `json:"duration"`
}

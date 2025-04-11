package models

type Exercise struct {
	ID             uint `gorm:"primaryKey"`
	WorkoutID      uint
	Workout        Workout `gorm:"foreignKey:WorkoutID"`
	ExerciseTypeID uint
	ExerciseType   ExerciseType `gorm:"foreignKey:ExerciseTypeID"`
	Sets           []Set        `gorm:"foreignKey:ExerciseID"`
	IsTimeBased    bool         `gorm:"default:false"`
}

type ExerciseResponse struct {
	ID             uint        `json:"id"`
	ExerciseTypeId uint        `json:"exerciseTypeId"`
	ExerciseType   string      `json:"exerciseType"`
	Sets           interface{} `json:"sets"` // Can be either []SetWeightResponse or []SetDurationResponse
}

func NewExerciseResponse(exercise Exercise) ExerciseResponse {
	// Convert the exercise type ID to its name using the ExerciseType model
	exerciseType := NewExerciseTypeResponse(exercise.ExerciseType)

	var sets interface{}
	if exercise.IsTimeBased {
		// Convert to duration-based sets
		var durationSets []SetDurationResponse
		for _, s := range exercise.Sets {
			durationSets = append(durationSets, SetDurationResponse{
				ID:       s.ID,
				Duration: s.Duration,
			})
		}
		if len(durationSets) == 0 {
			durationSets = []SetDurationResponse{}
		}
		sets = durationSets
	} else {
		// Convert to weight-based sets
		var weightSets []SetWeightResponse
		for _, s := range exercise.Sets {
			weightSets = append(weightSets, SetWeightResponse{
				ID:     s.ID,
				Weight: s.Weight,
				Reps:   s.Reps,
			})
		}
		if len(weightSets) == 0 {
			weightSets = []SetWeightResponse{}
		}
		sets = weightSets
	}

	return ExerciseResponse{
		ID:             exercise.ID,
		ExerciseTypeId: exerciseType.ID,
		ExerciseType:   exerciseType.Name,
		Sets:           sets,
	}
}

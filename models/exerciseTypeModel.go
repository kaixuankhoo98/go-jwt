package models

type ExerciseType struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Name      string     `gorm:"not null"`
	GroupID   *uint      // Nullable in case user doesn’t assign a group
	Group     *Group     `gorm:"foreignKey:GroupID"`
	User      User       `gorm:"foreignKey:UserID"`
	Exercises []Exercise `gorm:"foreignKey:ExerciseTypeID"`
}

type ExerciseTypeResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	GroupId *uint  `json:"groupId"`
}

func NewExerciseTypeResponse(exerciseType ExerciseType) ExerciseTypeResponse {
	return ExerciseTypeResponse{
		ID:      exerciseType.ID,
		Name:    exerciseType.Name,
		GroupId: exerciseType.GroupID,
	}
}

func NewExerciseTypeResponseList(exerciseTypes []ExerciseType) []ExerciseTypeResponse {
	exerciseTypeResponses := make([]ExerciseTypeResponse, len(exerciseTypes))
	for i, exerciseType := range exerciseTypes {
		exerciseTypeResponses[i] = NewExerciseTypeResponse(exerciseType)
	}
	return exerciseTypeResponses
}

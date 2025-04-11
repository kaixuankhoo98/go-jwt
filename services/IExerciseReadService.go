package services

type IExerciseReadService interface {
	GetWorkoutIdByExerciseId(exerciseId uint) (uint, error)
	GetIsTimeBasedExercise(exerciseId uint) (bool, error)
}

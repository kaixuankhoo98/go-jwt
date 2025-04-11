package services

type IWorkoutReadService interface {
	VerifyWorkoutBelongsToUser(userId uint, workoutId uint) error
}

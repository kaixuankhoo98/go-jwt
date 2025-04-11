package services

type IExerciseTypeReadService interface {
	VerifyExerciseTypeBelongsToUser(userId uint, exerciseTypeId uint) error
}

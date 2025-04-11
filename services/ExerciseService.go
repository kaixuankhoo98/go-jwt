package services

import (
	"errors"
	"go-jwt/models"

	"gorm.io/gorm"
)

type exerciseService struct {
	db                      *gorm.DB
	workoutReadService      IWorkoutReadService
	exerciseTypeReadService IExerciseTypeReadService
}

func NewExerciseService(database *gorm.DB, wrs IWorkoutReadService, ets IExerciseTypeReadService) IExerciseService {
	return &exerciseService{db: database, workoutReadService: wrs, exerciseTypeReadService: ets}
}

func (service *exerciseService) CreateExercise(userId uint, workoutId uint, exerciseTypeId uint, isTimeBased bool) (uint, error) {
	err := service.workoutReadService.VerifyWorkoutBelongsToUser(userId, workoutId)
	if err != nil {
		return 0, err
	}
	err = service.exerciseTypeReadService.VerifyExerciseTypeBelongsToUser(userId, exerciseTypeId)
	if err != nil {
		return 0, err
	}

	exercise := models.Exercise{WorkoutID: workoutId, ExerciseTypeID: exerciseTypeId, IsTimeBased: isTimeBased}
	if err := service.db.Create(&exercise).Error; err != nil {
		return 0, errors.New("failed to create exercise")
	}

	return exercise.ID, nil
}

func (service *exerciseService) GetExercisesByWorkoutId(userId uint, workoutId uint) ([]models.ExerciseResponse, error) {
	err := service.workoutReadService.VerifyWorkoutBelongsToUser(userId, workoutId)
	if err != nil {
		return nil, err
	}

	var exercises []models.Exercise
	if err := service.db.Where("workout_id = ?", workoutId).Preload("ExerciseType").Preload("Sets").Find(&exercises).Error; err != nil {
		return nil, errors.New("failed to fetch exercises")
	}

	exerciseResponses := make([]models.ExerciseResponse, len(exercises))
	for i, exercise := range exercises {
		exerciseResponses[i] = models.NewExerciseResponse(exercise)
	}

	return exerciseResponses, nil
}

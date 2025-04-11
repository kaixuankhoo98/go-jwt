package services

import (
	"errors"
	"go-jwt/models"

	"gorm.io/gorm"
)

type setService struct {
	db                  *gorm.DB
	exerciseReadService IExerciseReadService
	workoutReadService  IWorkoutReadService
}

func NewSetService(database *gorm.DB, ers IExerciseReadService, wrs IWorkoutReadService) ISetService {
	return &setService{db: database, exerciseReadService: ers, workoutReadService: wrs}
}

func (s *setService) CreateWeightedSet(userId uint, exerciseId uint, weight float64, reps int) (*models.Set, error) {
	workoutId, err := s.exerciseReadService.GetWorkoutIdByExerciseId(exerciseId)
	if err != nil {
		return nil, errors.New("invalid exerciseId")
	}

	err = s.workoutReadService.VerifyWorkoutBelongsToUser(userId, workoutId)
	if err != nil {
		return nil, err
	}

	if isTimeBased, err := s.exerciseReadService.GetIsTimeBasedExercise(exerciseId); err != nil {
		return nil, errors.New("failed to check if exercise is time based")
	} else if isTimeBased {
		return nil, errors.New("exercise is time based, cannot create weighted set")
	}

	set := models.Set{ExerciseID: exerciseId, Weight: &weight, Reps: &reps}
	if err := s.db.Create(&set).Error; err != nil {
		return nil, errors.New("failed to create set")
	}

	return &set, nil
}

func (s *setService) CreateDurationSet(userId uint, exerciseId uint, duration int) (*models.Set, error) {
	workoutId, err := s.exerciseReadService.GetWorkoutIdByExerciseId(exerciseId)
	if err != nil {
		return nil, errors.New("invalid exerciseId")
	}

	err = s.workoutReadService.VerifyWorkoutBelongsToUser(userId, workoutId)
	if err != nil {
		return nil, err
	}

	if isTimeBased, err := s.exerciseReadService.GetIsTimeBasedExercise(exerciseId); err != nil {
		return nil, errors.New("failed to check if exercise is time based")
	} else if !isTimeBased {
		return nil, errors.New("exercise is not time based, cannot create duration set")
	}

	set := models.Set{ExerciseID: exerciseId, Duration: &duration}
	if err := s.db.Create(&set).Error; err != nil {
		return nil, errors.New("failed to create set")
	}

	return &set, nil
}

package services

import "gorm.io/gorm"

type exerciseReadService struct {
	db *gorm.DB
}

func NewExerciseReadService(database *gorm.DB) IExerciseReadService {
	return &exerciseReadService{db: database}
}

func (s *exerciseReadService) GetWorkoutIdByExerciseId(exerciseId uint) (uint, error) {
	var workoutId uint
	if err := s.db.Table("exercises").Select("workout_id").Where("id = ?", exerciseId).Scan(&workoutId).Error; err != nil {
		return 0, err
	}
	return workoutId, nil
}

func (s *exerciseReadService) GetIsTimeBasedExercise(exerciseId uint) (bool, error) {
	var isTimeBased bool
	if err := s.db.Table("exercises").Select("is_time_based").Where("id = ?", exerciseId).Scan(&isTimeBased).Error; err != nil {
		return false, err
	}
	return isTimeBased, nil
}

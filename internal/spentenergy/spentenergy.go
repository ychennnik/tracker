package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// validateParams проверяет, что входные параметры расчёта калорий имеют смысл:
// шаги, вес, рост и длительность должны быть строго положительными.
func validateParams(steps int, weight, height float64, duration time.Duration) error {
	switch {
	case steps <= 0:
		return errors.New("количество шагов должно быть больше 0")
	case weight <= 0:
		return errors.New("вес должен быть больше 0")
	case height <= 0:
		return errors.New("рост должен быть больше 0")
	case duration <= 0:
		return errors.New("продолжительность должна быть больше 0")
	}
	return nil
}

// spentCalories — общая часть расчёта для бега и ходьбы:
// (вес * средняя скорость * длительность в минутах) / минут в часе.
func spentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateParams(steps, weight, height, duration); err != nil {
		return 0, err
	}

	speed := MeanSpeed(steps, height, duration)
	return (weight * speed * duration.Minutes()) / minInH, nil
}

// WalkingSpentCalories возвращает количество калорий, потраченных при ходьбе.
// При некорректных параметрах возвращает 0 и ошибку.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := spentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return calories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories возвращает количество калорий, потраченных при беге.
// При некорректных параметрах возвращает 0 и ошибку.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	return spentCalories(steps, weight, height, duration)
}

// MeanSpeed возвращает среднюю скорость в км/ч.
// Если шаги или длительность не положительны, возвращает 0.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

// Distance возвращает пройденную дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

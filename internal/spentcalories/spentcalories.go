package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/errorsx"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	splitedData := strings.Split(data, ",")
	if len(splitedData) != 3 {
		return 0, "0", 0, errorsx.ErrWrongNumberOfElements
	}

	steps, err := strconv.Atoi(splitedData[0])
	if err != nil {
		return 0, "0", 0, err
	}
	if steps <= 0 {
		return 0, "0", 0, errorsx.ErrNegativeSteps
	}

	activityType := splitedData[1]

	duration, err := time.ParseDuration(splitedData[2])
	if err != nil {
		return 0, "0", 0, err
	}
	if duration <= 0 {
		return 0, "0", 0, errorsx.ErrNegativeDuration
	}

	return steps, activityType, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return (float64(steps) * height * stepLengthCoefficient) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0.0
	}
	distance := distance(steps, height)
	return distance / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	form := `Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f`
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	speed := meanSpeed(steps, height, duration)
	distance := distance(steps, height)
	switch activityType {
	case "Бег":
		spentCalories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(form+"\n", activityType, duration.Hours(), distance, speed, spentCalories), nil
	case "Ходьба":
		spentCalories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(form+"\n", activityType, duration.Hours(), distance, speed, spentCalories), nil
	default:
		return "", errorsx.ErrUnknownAcrivityType
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0.0, errorsx.ErrNegativeSteps
	}
	if weight <= 0 {
		return 0.0, errorsx.ErrNegativeWeight
	}
	if height <= 0 {
		return 0.0, errorsx.ErrNegativeHeight
	}
	if duration <= 0 {
		return 0.0, errorsx.ErrNegativeDuration
	}

	meanSpeed := meanSpeed(steps, height, duration)

	return (weight * meanSpeed * duration.Minutes()) / minInH, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0.0, errorsx.ErrNegativeSteps
	}
	if weight <= 0 {
		return 0.0, errorsx.ErrNegativeWeight
	}
	if height <= 0 {
		return 0.0, errorsx.ErrNegativeHeight
	}
	if duration <= 0 {
		return 0.0, errorsx.ErrNegativeDuration
	}

	meanSpeed := meanSpeed(steps, height, duration)

	return ((weight * meanSpeed * duration.Minutes()) / minInH) * walkingCaloriesCoefficient, nil

}

package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/errorsx"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	splitedData := strings.Split(data, ",")
	if len(splitedData) != 2 {
		return 0, 0, errorsx.ErrWrongNumberOfElements
	}
	steps, err := strconv.Atoi(splitedData[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, errorsx.ErrNegativeSteps
	}

	duration, err := time.ParseDuration(splitedData[1])
	if err != nil {
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, errorsx.ErrNegativeDuration
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	form := `Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.`
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps <= 0 {
		log.Println(err)
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm
	if duration < 0 {
		log.Println(err)
		return ""
	}
	spentcalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}
	return fmt.Sprintf(form+"\n", steps, distance, spentcalories)
}

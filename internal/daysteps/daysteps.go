package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {

	dataRaw := strings.Split(data, ",")
	if len(dataRaw) != 2 {
		return 0, 0, fmt.Errorf("невалидные данные о тренировке")
	}
	steps, err := strconv.Atoi(dataRaw[0])
	if err != nil {
		return 0, 0, fmt.Errorf("невалидные данные о тренировке")
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("невалидные данные о тренировке")
	}
	timeSpent, err := time.ParseDuration(dataRaw[1])
	if err != nil {
		return 0, 0, fmt.Errorf("невалидные данные о тренировке")
	}
	return steps, timeSpent, nil
}

func DayActionInfo(data string, weight, height float64) string {

	steps, duration, err := parsePackage(data)

	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distance := float64(steps) * float64(stepLength) / float64(mInKm)

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.",
		steps, distance, calories,
	)
	return result
}

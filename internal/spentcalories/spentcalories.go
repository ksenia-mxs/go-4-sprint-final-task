package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
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

	dataRaw := strings.Split(data, ",")
	if len(dataRaw) != 3 {
		return 0, "", 0, fmt.Errorf("невалидные данные о тренировке")
	}
	steps, err := strconv.Atoi(dataRaw[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("невалидные данные о тренировке")
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("невалидные данные о тренировке")
	}
	activity := dataRaw[1]
	timeSpent, err := time.ParseDuration(dataRaw[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("невалидные данные о тренировке")
	}
	if timeSpent <= 0 {
		return 0, "", 0, fmt.Errorf("невалидные данные о тренировке")
	}
	return steps, activity, timeSpent, nil

}

func distance(steps int, height float64) float64 {

	result := height * float64(stepLengthCoefficient) * float64(steps) / float64(mInKm)
	return result

}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration <= 0 {
		return 0
	}
	walkedDistance := distance(steps, height)
	result := walkedDistance / duration.Hours()
	return result
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("тренировка не состоялась")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	result := weight * meanSpeed * duration.Minutes() / float64(minInH)
	return result, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("тренировка не состоялась")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	result := weight * meanSpeed * duration.Minutes() / float64(minInH) * float64(walkingCaloriesCoefficient)
	return result, nil
}

func TrainingInfo(data string, weight, height float64) (string, error) {

	steps, activity, timeSpent, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	distance := distance(steps, height)
	speed := meanSpeed(steps, height, timeSpent)

	switch activity {
	case "Бег":
		{
			calories, err := RunningSpentCalories(steps, weight, height, timeSpent)
			if err != nil {
				return "", err
			}
			result := fmt.Sprintf(
				"Тип тренировки: %s\n"+
					"Длительность: %.2f ч.\n"+
					"Дистанция: %.2f км.\n"+
					"Скорость: %.2f км/ч\n"+
					"Сожгли калорий: %.2f\n",
				activity, timeSpent.Hours(), distance, speed, calories,
			)
			return result, nil
		}

	case "Ходьба":
		{
			calories, err := WalkingSpentCalories(steps, weight, height, timeSpent)
			if err != nil {
				return "", err
			}
			result := fmt.Sprintf(
				"Тип тренировки: %s\n"+
					"Длительность: %.2f ч.\n"+
					"Дистанция: %.2f км.\n"+
					"Скорость: %.2f км/ч\n"+
					"Сожгли калорий: %.2f\n",
				activity, timeSpent.Hours(), distance, speed, calories,
			)
			return result, nil
		}

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

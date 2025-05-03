package spentenergy

import (
	"errors"
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories-Функция принимает:количество шагов;вес(кг.) и рост(м.) пользователя;продолжительность бега и
// возвращает два значения:количество калорий, потраченных при беге и ошибку
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("длительность тренировки должна быть больше 0")
	}
	// speed- средняя скорость(результат функции meanSpeed)
	speed := MeanSpeed(steps, height, duration)
	// durationInMinutes- продолжительность в минутах
	durationInMinutes := duration.Minutes()
	return ((weight * speed * durationInMinutes) / minInH) * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories-Функция принимает:количество шагов;вес(кг.) и рост(м.) пользователя;продолжительность бега и
// возвращает два значения:количество калорий, потраченных при беге и ошибку
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("длительность тренировки должна быть больше 0")
	}
	//speed- средняя скорость(результат функции meanSpeed)
	speed := MeanSpeed(steps, height, duration)
	//durationInMinutes- продолжительность в минутах
	durationInMinutes := duration.Minutes()
	return (weight * speed * durationInMinutes) / minInH, nil
}

// meanSped принимает количество шагов steps, рост пользователя height и продолжительность активности duration  и возвращает среднюю скорость
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		fmt.Println("количество шагов должно быть больше  0")
		return 0
	}
	if duration <= 0 {
		fmt.Println("длительность тренировки должна быть больше 0")
		return 0
	}
	//hours- продолжительность в часах
	hours := duration.Hours()
	//distance - пройденная дистанция (результат функции Distance)
	distance := Distance(steps, height)
	//avarageSpeed - средняя скорость
	avarageSpeed := distance / hours
	return avarageSpeed
}

// distance принимает количество шагов и рост пользователя в метрах, а возвращает дистанцию в километрах
func Distance(steps int, height float64) float64 {
	//stepLength - длина шага
	stepLength := height * stepLengthCoefficient
	//walkDistance - пройденная дистанция,умножаем количество шагов на длину шага
	walkDistance := float64(steps) * stepLength
	walkDistance = walkDistance / float64(mInKm)
	return walkDistance
}

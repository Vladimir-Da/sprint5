package trainings

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training содержит все необходимые данные о тренировке: количество шагов, тип тренировки, длительность тренировки, а также
// данные из структуры personaldata.Personal, то есть имя, вес и рост пользователя.
type Training struct {
	// Steps - количетво шагов
	Steps int
	// TrainingType - тип тренировки(бег или ходьба).
	TrainingType string
	// Duration - продолжительность тренировки
	Duration time.Duration
	//
	personaldata.Personal
}

// Parse парсит строку с данными формата "3456,Ходьба,3h00m" и записывает данные в соответствующие поля структуры Training
func (t *Training) Parse(datastring string) (err error) {
	if len(datastring) == 0 {
		err := errors.New("wrong data")
		log.Println("error:", err)
		return err
	}
	str := strings.Split(datastring, ",")
	if len(str) != 3 {
		err := errors.New("want - steps,tranning type ,duration; but have ")
		log.Println("error:", err, datastring)
		return err
	}
	// steps преобразованное интовое значение кол-во шагов
	steps, err := strconv.Atoi(str[0])
	if err != nil {
		err := errors.New("want string to parse into int ; but have ")
		log.Println("error:", err, str[0])
		return err
	}
	if steps <= 0 {
		err := errors.New("steps must be positive")
		log.Println("error: steps = ", steps, err)
		return err
	}
	t.Steps = steps
	// typeTrain - тип тренировки (бег,ходьба)
	typeTrain := str[1]
	t.TrainingType = typeTrain
	// dutation - продолжительность тренировки (формат время)
	duration, err := time.ParseDuration(str[2])
	if err != nil {
		err := errors.New(" cant parse into duration")
		log.Println("error: ", str[2], err)
		return err
	}
	if duration == 0 {
		err := errors.New("the training lasted 0")
		log.Println("error: ", err)
		return err
	}
	if duration < 0 {
		err := errors.New("duration is negative")
		log.Println("error:", err)
		return err
	}
	t.Duration = duration
	return nil
}

// ActionInfo формирует и возвращает строку с данными о тренировке, исходя из того, какой тип тренировки был передан
func (t Training) ActionInfo() (string, error) {
	//distance -дистанция
	distance := spentenergy.Distance(t.Steps, t.Height)
	// durHours - продолжительность в часах
	durHours := t.Duration.Hours()
	// spentEnergy - затрачено калорий (отличается при Беге и Ходьбе)
	var spentEnergy float64
	switch t.TrainingType {
	case "Ходьба":
		{
			energy, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
			if err != nil {
				log.Println("error:", err)
			}
			spentEnergy = energy
		}

	case "Бег":
		{
			energy, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
			if err != nil {
				log.Println("error:", err)
			}
			spentEnergy = energy
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, durHours, distance, speed, spentEnergy), nil

}

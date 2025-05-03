package daysteps

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

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// метод Parse парсит строку с данными формата "678,0h50m" и записывает данные в соответствующие поля структуры DaySteps
func (ds *DaySteps) Parse(datastring string) (err error) {
	if len(datastring) == 0 {
		err := errors.New("no data")
		log.Println("error:", err)
		return err
	}
	//str слайс строк str[]- шаги , str[1] -продолжительность
	str := strings.Split(datastring, ",")
	if len(str) != 2 {
		err := errors.New("want - steps,duration; but have ")
		log.Println("error:", err, datastring)
		return err
	}
	// steps-преобразованное интовое значение кол-во шагов
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
	ds.Steps = steps
	// dutation - продолжительность тренировки (формат время)
	duration, err := time.ParseDuration(str[1])
	if err != nil {
		err := errors.New(" cant parse into duration")
		log.Println("error: ", str[1], err)
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
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// ActionInfo формирует и возвращает строку с данными о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	// distance - вычисляем дисстанцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	// spentcalories - вычисляем потраченные калории
	spentcalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentcalories), nil

}

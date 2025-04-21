package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	. "github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	strings := strings.Split(data, ",")
	if len(strings) != 2 {
		log.Println(errors.New("parsePackage: len(strings) !=2"))
		return 0, 0, errors.New("parsePackage: len(strings) !=2")
	}
	steps, err := strconv.Atoi(strings[0])
	if err != nil {
		log.Println(err)
		return 0, 0, errors.New("parsePackage: cant convert steps to string")
	}
	if steps < 1 {
		log.Println(err)
		return 0, 0, errors.New("parsePackage: steps <= 0")
	}
	duration, err := time.ParseDuration(strings[1])
	if err != nil {
		log.Println(err)
		return 0, 0, errors.New("parsePackage: cant convert duration to time.Duration")
	}
	if duration <= time.Duration(0) {
		log.Println(errors.New("DayActionInfo: duration <= 0"))
		return 0, 0, errors.New("DayActionInfo: duration <= 0")
	}
	return steps, duration, err
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return ""
	}
	if steps < 1 {
		log.Println(errors.New("DayActionInfo: steps <= 0"))
		return ""
	}
	distance := stepLength * float64(steps) / float64(mInKm)
	calories, err := WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}
	if duration <= time.Duration(0) {
		log.Println(errors.New("DayActionInfo: duration <= 0"))
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}

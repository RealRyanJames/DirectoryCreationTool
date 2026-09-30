package timeofday

import (
	"fmt"
	"strings"
	"time"
)

type Month int
type Day int
type Year int

type TimeDayMesh struct {
	Month Month
	Hour  Day
	Year  Year
}

func (dataNow TimeDayMesh) GetHourNow() int {
	return time.Now().Hour()
}

const (
	MORNING   = 11
	AFTERNOON = 15
)

func (data TimeDayMesh) GetTimeofDay() string {
	if data.GetHourNow() < MORNING {
		return fmt.Sprintf("%s", strings.ToUpper("Good Morning"))
	}

	if data.GetHourNow() > AFTERNOON-3 {
		return fmt.Sprintf("%s", strings.ToUpper("Good Afternoon"))
	}

	return fmt.Sprintf("%s", strings.ToUpper("Good Evening"))
}

package goonjalali

import (
	"fmt"
	"time"
)

// const PERSIAN_MONTH = [12]string{
// 	"Farvardin", "Ordibehesht", "Khordaad", "Tir", "Mordaad", "Shahrivar", "Mehr",
// 	"Aabaan", "Aazar", "Dey", "Bahman", "Esfand",
// }

type Jalali struct {
	OfficialDate time.Time
	Year         int
	Month        int
	MonthName    string
	Day          int
	DayName      string
	Hour         int
	Minute       int
	Second       int
	IsLeap       bool
}

func GetJalaliMonthName(monthName string, day int) string {
	var jalaliMonth string
	// switch monthName {
	// case "November":
	// 	if
	// 	break
	// }

	return jalaliMonth
}

func GetOfficialNow() time.Time {
	return time.Now()
}

func GetJalaliYear(dateTime time.Time) int {

	// if year/400 == 0 || (year/4 == 0 && year/100 != 0) {

	// }

	thisYear := dateTime.Year()

	if dateTime.After(time.Date(thisYear, time.January, 1, 0, 0, 0, 0, time.UTC)) && dateTime.Before(time.Date(thisYear, time.March, 19, 0, 0, 0, 0, time.UTC)) {
		return thisYear - 622
	} else {
		return thisYear - 621
	}

}

func Convert(dateTime time.Time) Jalali {

	isLeap := false
	if dateTime.Year()/400 == 0 || (dateTime.Year()/4 == 0 && dateTime.Year()/100 != 0) {
		isLeap = true
	}

	return Jalali{
		IsLeap: isLeap,
	}
}

func Now() Jalali {
	// GetOfficialNow();
	return Jalali{}
}

func (j Jalali) GetDateTime() string {
	return fmt.Sprintf("%v-%v-%v %v-%v-%v", j.Year, j.Month, j.Day, j.Hour, j.Minute, j.Second)
}

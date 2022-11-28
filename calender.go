package goonjalali

import (
	"fmt"
	"time"
)

var PERSIAN_MONTH [12]string = [12]string{
	"Farvardin", "Ordibehesht", "Khordaad", "Tir", "Mordaad", "Shahrivar", "Mehr",
	"Aabaan", "Azar", "Dey", "Bahman", "Esfand",
}

var FARSI_MONTH_NAME map[string]string = map[string]string{
	"Farvardin": "فروردین", "Ordibehesht": "اردیبهشت", "Khordaad": "خرداد",
	"Tir": "تیر", "Mordaad": "مرداد", "Shahrivar": "شهریور", "Mehr": "مهر",
	"Aabaan": "آبان", "Azar": "آذر", "Dey": "دی", "Bahman": "بهمن", "Esfand": "اسفند",
}

type Jalali struct {
	OfficialDate time.Time
	Year         int
	Month        int
	MonthName    string
	Day          int
	DayName      string
	Time         string
}

func GetJalaliMonthName(dateTime time.Time) (int, string, error) {

	if dateTime.After(time.Date(dateTime.Year(), time.March, 21, 0, 0, 0, 0, time.Local)) &&
		dateTime.Before(time.Date(dateTime.Year(), time.April, 20, 23, 59, 59, 999999999, time.Local)) {
		return 1, PERSIAN_MONTH[0], nil
	}

	if dateTime.After(time.Date(dateTime.Year(), time.April, 21, 0, 0, 0, 0, time.Local)) &&
		dateTime.Before(time.Date(dateTime.Year(), time.May, 21, 0, 0, 0, 0, time.Local)) {
		return 2, PERSIAN_MONTH[1], nil
	}

	if dateTime.After(time.Date(dateTime.Year(), time.May, 22, 0, 0, 0, 0, time.Local)) &&
		dateTime.Before(time.Date(dateTime.Year(), time.June, 21, 0, 0, 0, 0, time.Local)) {
		return 3, PERSIAN_MONTH[2], nil
	}

	if dateTime.After(time.Date(dateTime.Year(), time.June, 22, 0, 0, 0, 0, time.Local)) &&
		dateTime.Before(time.Date(dateTime.Year(), time.July, 22, 0, 0, 0, 0, time.Local)) {
		return 4, PERSIAN_MONTH[3], nil
	}

	if dateTime.After(time.Date(dateTime.Year(), time.July, 23, 0, 0, 0, 0, time.Local)) &&
		dateTime.Before(time.Date(dateTime.Year(), time.August, 22, 0, 0, 0, 0, time.Local)) {
		return 5, PERSIAN_MONTH[4], nil
	}

	if dateTime.After(time.Date(dateTime.Year(), time.August, 23, 0, 0, 0, 0, time.Local)) &&
		dateTime.Before(time.Date(dateTime.Year(), time.September, 22, 0, 0, 0, 0, time.Local)) {
		return 6, PERSIAN_MONTH[5], nil
	}

	if dateTime.After(time.Date(dateTime.Year(), time.September, 23, 0, 0, 0, 0, time.Local)) &&
		dateTime.Before(time.Date(dateTime.Year(), time.October, 22, 0, 0, 0, 0, time.Local)) {
		return 7, PERSIAN_MONTH[6], nil
	}

	if dateTime.After(time.Date(dateTime.Year(), time.October, 23, 0, 0, 0, 0, time.Local)) &&
		dateTime.Before(time.Date(dateTime.Year(), time.November, 21, 0, 0, 0, 0, time.Local)) {
		return 8, PERSIAN_MONTH[7], nil
	}

	if dateTime.After(time.Date(dateTime.Year(), time.November, 22, 0, 0, 0, 0, time.Local)) &&
		dateTime.Before(time.Date(dateTime.Year(), time.December, 21, 0, 0, 0, 0, time.Local)) {
		return 9, PERSIAN_MONTH[8], nil
	}

	if dateTime.After(time.Date(dateTime.Year(), time.December, 22, 0, 0, 0, 0, time.Local)) &&
		dateTime.Before(time.Date(dateTime.Year(), time.January, 20, 0, 0, 0, 0, time.Local)) {
		return 10, PERSIAN_MONTH[9], nil
	}

	if dateTime.After(time.Date(dateTime.Year(), time.January, 21, 0, 0, 0, 0, time.Local)) &&
		dateTime.Before(time.Date(dateTime.Year(), time.February, 29, 0, 0, 0, 0, time.Local)) {
		return 11, PERSIAN_MONTH[10], nil
	}

	if dateTime.After(time.Date(dateTime.Year(), time.February, 20, 0, 0, 0, 0, time.Local)) &&
		dateTime.Before(time.Date(dateTime.Year(), time.March, 20, 0, 0, 0, 0, time.Local)) {
		return 12, PERSIAN_MONTH[11], nil
	}

	return 0, "", nil
}

func GetOfficialNow() time.Time {
	return time.Now()
}

func GetJalaliYear(dateTime time.Time) int {

	thisYear := dateTime.Year()

	if dateTime.After(time.Date(thisYear, time.January, 1, 0, 0, 0, 0, time.Local)) && dateTime.Before(time.Date(thisYear, time.March, 19, 0, 0, 0, 0, time.Local)) {
		return thisYear - 622
	} else {
		return thisYear - 621
	}

}

func Convert(dateTime time.Time) Jalali {

	// isLeap := false
	// if dateTime.Year()/400 == 0 || (dateTime.Year()/4 == 0 && dateTime.Year()/100 != 0) {
	// 	isLeap = true
	// }
	monthNumber, monthName, _ := GetJalaliMonthName(dateTime)

	return Jalali{
		OfficialDate: dateTime,
		Year:         GetJalaliYear(dateTime),
		Month:        monthNumber,
		MonthName:    monthName,
		Time:         dateTime.Format("03:04:05"),
	}
}

func Now() Jalali {
	return Convert(time.Now())
}

func (j Jalali) GetDateTime() string {
	return fmt.Sprintf("%v-%v-%v %v", j.Year, j.Month, j.Day, j.Time)
}

func (j Jalali) GetDate() string {
	return fmt.Sprintf("%v-%v-%v", j.Year, j.Month, j.Day)
}

func (j Jalali) GetTime() string {
	return fmt.Sprintf("%v", j.Time)
}

func (j Jalali) GetPersianMonthName() string {
	return FARSI_MONTH_NAME[j.MonthName]
}

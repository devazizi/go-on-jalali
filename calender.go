package goonjalali

import (
	"fmt"
	"strconv"
	"strings"
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

var MONTHS_DAY_COUNT map[string]int = map[string]int{
	"January":   31,
	"February":  28,
	"March":     31,
	"April":     30,
	"May":       31,
	"June":      30,
	"July":      31,
	"August":    31,
	"September": 30,
	"October":   31,
	"November":  30,
	"December":  31,
}

var JALALI_MONTH_TO_GLOBAL map[string]string = map[string]string{
	"Farvardin":   "21 March-20 April",
	"Ordibehesht": "21 April-21 May",
	"Khordaad":    "22 May-21 June",
	"Tir":         "22 June-22 July",
	"Mordaad":     "23 July-22 August",
	"Shahrivar":   "23 August-22 September",
	"Mehr":        "23 September-22 October",
	"Aabaan":      "23 October-21 November",
	"Azar":        "22 November-21 December",
	"Dey":         "22 December-20 January",
	"Bahman":      "21 January-19 February",
	"Esfand":      "20 February-20 March",
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

func IsLeapYear(year int) bool {

	if year/400 == 0 || (year/4 == 0 && year/100 != 0) {
		return true
	}

	return false
}

func GetJalaliDay(monthName string, dateTime time.Time) int {
	monthDetails := JALALI_MONTH_TO_GLOBAL[monthName]

	jalaliMonthToGlobal := strings.Split(monthDetails, "-")

	firstMonthDetails := strings.Split(jalaliMonthToGlobal[0], " ")
	secondMonthDetails := strings.Split(jalaliMonthToGlobal[1], " ")

	firstMonthCount := MONTHS_DAY_COUNT[firstMonthDetails[1]]
	secondMonthCount := MONTHS_DAY_COUNT[secondMonthDetails[1]]

	if monthName == "Bahman" {
		if IsLeapYear(dateTime.Year()) {
			secondMonthCount++
		}
	}

	if monthName == "Esfand" {
		if IsLeapYear(dateTime.Year()) {
			firstMonthCount++
		}
	}

	var day int
	if dateTime.Month().String() == firstMonthDetails[1] {
		diff, _ := strconv.Atoi(firstMonthDetails[0])
		day = dateTime.Day() - diff + 1
	} else {
		diff, _ := strconv.Atoi(firstMonthDetails[0])
		day = dateTime.Day() + (firstMonthCount - diff) + 1
	}

	return day
}

func Convert(dateTime time.Time) Jalali {
	monthNumber, monthName, _ := GetJalaliMonthName(dateTime)

	return Jalali{
		OfficialDate: dateTime,
		Year:         GetJalaliYear(dateTime),
		Day:          GetJalaliDay(monthName, dateTime),
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

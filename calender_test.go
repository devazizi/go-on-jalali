package goonjalali_test

import (
	"fmt"
	"testing"
	"time"

	goonjalali "github.com/devazizi/go-on-jalali"
)

func TestNowDateIsCurrect(t *testing.T) {
	officialDateTime := goonjalali.GetOfficialNow()
	now := time.Now()
	if now.Day() != officialDateTime.Day() {
		t.Errorf("date didn't equal %v %v", now.Format("2006-1-2"), officialDateTime.Format("2006-1-2"))
	}
}

func TestConvertYearIsCorrect(t *testing.T) {
	firstJalaliYear := goonjalali.GetJalaliYear(time.Date(2022, time.November, 28, 1, 1, 1, 1, time.UTC))
	if firstJalaliYear != 1401 {
		t.Errorf("error data incorrect %v must %v", firstJalaliYear, 1401)
	}

	secondJalaliYear := goonjalali.GetJalaliYear(time.Date(2020, time.March, 29, 1, 1, 1, 1, time.UTC))
	if secondJalaliYear != 1399 {
		t.Errorf("error data incorrect %v must %v", secondJalaliYear, 1399)
	}
}

func TestConvertYearInFutureIsCorrect(t *testing.T) {
	futureYear := goonjalali.GetJalaliYear(time.Date(2026, time.September, 29, 1, 1, 1, 1, time.UTC))
	if futureYear != 1405 {
		t.Errorf("error data incorrect %v must %v", futureYear, 1405)
	}
}

func TestGetJalaliMonthName(t *testing.T) {
	monthNumber, monthName, _ := goonjalali.GetJalaliMonthName(time.Date(2022, time.November, 28, 1, 1, 1, 1, time.UTC))

	if monthNumber != 9 && monthName != "Azar" {
		t.Errorf("month must equal Azar but equal %v", monthName)
	}
}

// func TestCorrectConvertOfficailMonthToJalali(t *testing.T) {
// 	// firstDateTime := time.Date(2022, 11, 28, 15, 20 , 20, 0, time.Now().Local().Location())
// 	// firstConvert := goonjalali.GetJalaliMonthName(firstDateTime.Month().String(), firstDateTime.Day())

// }

func TestFullConvertDateToJalali(t *testing.T) {
	exampleTime := time.Date(2022, time.November, 28, 1, 23, 45, 32, time.Local)

	fmt.Println(goonjalali.Convert(exampleTime).GetPersianMonthName())
}

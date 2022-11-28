package goonjalali_test

import (
	"testing"
	"time"

	goonjalali "github.com/devazizi/go-on-jalali"
)

func TestNowDateIsCurrect(t *testing.T) {
	officialDateTime := goonjalali.GetOfficialNow()
	now := time.Now()
	if now.Day() != officialDateTime.Day() {
		t.Errorf("dates didn't equal %v %v", now.Format("2006-1-2"), officialDateTime.Format("2006-1-2"))
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

// func TestCorrectConvertOfficailMonthToJalali(t *testing.T) {
// 	// firstDateTime := time.Date(2022, 11, 28, 15, 20 , 20, 0, time.Now().Local().Location())
// 	// firstConvert := goonjalali.GetJalaliMonthName(firstDateTime.Month().String(), firstDateTime.Day())

// }

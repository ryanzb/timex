package timex

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	tm := Now()
	if !timeEqual(tm, time.Now()) {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestNowUnix(t *testing.T) {
	ts := NowUnix()
	tm := time.Unix(ts, 0)
	if !timeEqual(tm, time.Now()) {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func timeEqual(tm, tm2 time.Time) bool {
	return tm.Year() == tm2.Year() && tm.Month() == tm2.Month() && tm.Day() == tm2.Day() &&
		tm.Hour() == tm2.Hour() && tm.Minute() == tm2.Minute() && tm.Second() == tm2.Second() &&
		tm.Location() == tm2.Location()
}

func TestDate(t *testing.T) {
	tm := Date(2019, 9, 1)
	if tm.Year() != 2019 || tm.Month() != 9 || tm.Day() != 1 ||
		tm.Hour() != 0 || tm.Minute() != 0 || tm.Second() != 0 || tm.Location() != time.Local {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestDateOf(t *testing.T) {
	tm := time.Date(2019, 9, 1, 1, 1, 1, 1, time.UTC)
	tm = DateOf(tm)
	if tm.Year() != 2019 || tm.Month() != 9 || tm.Day() != 1 ||
		tm.Hour() != 0 || tm.Minute() != 0 || tm.Second() != 0 || tm.Location() != time.Local {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestDateHourOf(t *testing.T) {
	tm := time.Date(2019, 9, 1, 1, 1, 1, 1, time.UTC)
	tm = DateHourOf(tm)
	if tm.Year() != 2019 || tm.Month() != 9 || tm.Day() != 1 ||
		tm.Hour() != 1 || tm.Minute() != 0 || tm.Second() != 0 || tm.Location() != time.Local {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestDateHalfHourOf(t *testing.T) {
	tm := time.Date(2019, 9, 1, 1, 31, 1, 1, time.UTC)
	tm = DateHalfHourOf(tm)
	if tm.Year() != 2019 || tm.Month() != 9 || tm.Day() != 1 ||
		tm.Hour() != 1 || tm.Minute() != 30 || tm.Second() != 0 || tm.Location() != time.Local {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestDateHourMinOf(t *testing.T) {
	tm := time.Date(2019, 9, 1, 1, 1, 1, 1, time.UTC)
	tm = DateHourMinOf(tm)
	if tm.Year() != 2019 || tm.Month() != 9 || tm.Day() != 1 ||
		tm.Hour() != 1 || tm.Minute() != 1 || tm.Second() != 0 || tm.Location() != time.Local {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestDateOfTimestamp(t *testing.T) {
	tm := DateOfTimestamp(1586448800)
	tm2 := Date(2020, 4, 10)
	if !timeEqual(tm, tm2) {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestDateOfToday(t *testing.T) {
	now := time.Now()
	tm := DateOfToday()
	if tm.Year() != now.Year() || tm.Month() != now.Month() || tm.Day() != now.Day() ||
		tm.Hour() != 0 || tm.Minute() != 0 || tm.Second() != 0 || tm.Location() != time.Local {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestDateOfTodayUnix(t *testing.T) {
	date := DateOfTodayUnix()
	if date != DateOfToday().Unix() {
		t.Error(date)
	} else {
		t.Log(date)
	}
}

func TestZeroOfToday(t *testing.T) {
	zero := ZeroOfToday()
	if zero != DateOfToday() {
		t.Error(zero)
	} else {
		t.Log(zero)
	}
}

func TestZeroOfTodayUnix(t *testing.T) {
	zero := ZeroOfTodayUnix()
	if zero != ZeroOfToday().Unix() {
		t.Error(zero)
	} else {
		t.Log(zero)
	}
}

func TestYesterday(t *testing.T) {
	tm := Yesterday()
	tm2 := time.Now().AddDate(0, 0, -1)
	if !timeEqual(tm, tm2) {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestDateOfYesterday(t *testing.T) {
	tm := DateOfYesterday()
	if tm.AddDate(0, 0, 1) != DateOfToday() {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestDateOfYesterdayUnix(t *testing.T) {
	date := DateOfYesterdayUnix()
	if date != DateOfYesterday().Unix() {
		t.Error(date)
	} else {
		t.Log(date)
	}
}

func TestZeroOfYesterday(t *testing.T) {
	zero := ZeroOfYesterday()
	if zero != DateOfYesterday() {
		t.Error(zero)
	} else {
		t.Log(zero)
	}
}

func TestZeroOfYesterdayUnix(t *testing.T) {
	zero := ZeroOfYesterdayUnix()
	if zero != ZeroOfYesterday().Unix() {
		t.Error(zero)
	} else {
		t.Log(zero)
	}
}

func TestTomorrow(t *testing.T) {
	tm := Tomorrow()
	tm2 := time.Now().AddDate(0, 0, 1)
	if !timeEqual(tm, tm2) {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestDateOfTomorrow(t *testing.T) {
	tm := DateOfTomorrow()
	if tm.AddDate(0, 0, -1) != DateOfToday() {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestDateOfTomorrowUnix(t *testing.T) {
	date := DateOfTomorrowUnix()
	if date != DateOfTomorrow().Unix() {
		t.Error(date)
	} else {
		t.Log(date)
	}
}

func TestZeroOfTomorrow(t *testing.T) {
	zero := ZeroOfTomorrow()
	if zero != DateOfTomorrow() {
		t.Error(zero)
	} else {
		t.Log(zero)
	}
}

func TestZeroOfTomorrowUnix(t *testing.T) {
	zero := ZeroOfTomorrowUnix()
	if zero != ZeroOfTomorrow().Unix() {
		t.Error(zero)
	} else {
		t.Log(zero)
	}
}

func TestDaysLater(t *testing.T) {
	tm := DaysLater(1)
	tm2 := Tomorrow()
	if !timeEqual(tm, tm2) {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestUnix(t *testing.T) {
	tm := Unix(1586448000)
	tm2 := Date(2020, 4, 10)
	if !timeEqual(tm, tm2) {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestParse(t *testing.T) {
	tm, err := Parse("2020-04-03 11:22:33")
	if err != nil {
		t.Error(err)
		return
	}

	if tm.Year() != 2020 || tm.Month() != 4 || tm.Day() != 3 ||
		tm.Hour() != 11 || tm.Minute() != 22 || tm.Second() != 33 ||
		tm.Location() != time.Local {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

func TestParseDate(t *testing.T) {
	tm, err := ParseDate("2020-04-03")
	if err != nil {
		t.Error(err)
		return
	}

	if tm.Year() != 2020 || tm.Month() != 4 || tm.Day() != 3 ||
		tm.Hour() != 0 || tm.Minute() != 0 || tm.Second() != 0 ||
		tm.Location() != time.Local {
		t.Error(tm)
	} else {
		t.Log(tm)
	}
}

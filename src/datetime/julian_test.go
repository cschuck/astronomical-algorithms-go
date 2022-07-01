package datetime

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addJulianDayTest struct {
	M         uint32
	D         float64
	Y         int
	gregorian bool
	expected  float64
}

var addJulianDayTests = []addJulianDayTest{
	addJulianDayTest{10, 4.81, 1957, true, 2436116.31},
	addJulianDayTest{1, 27.5, 333, false, 1842713.0},
}

func TestJulianDay(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	for _, test := range addJulianDayTests {
		if output := JulianDay(test.M, test.Y, test.D, test.gregorian); output != test.expected {
			t.Errorf("Output %f not equal to expected %f", output, test.expected)
		}
	}
}

func TestModifiedJulianDay(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	var M uint32 = 11
	var D = 17.0
	var Y = 1858
	var gregorian = true
	var JD = ModifiedJulianDay(M, Y, D, gregorian)
	log.Debug("JD = " + fmt.Sprint(JD))
	if JD != 0.0 {
		t.Fatal(JD)
	}
}

type addGregorianDateTest struct {
	JD             float64
	expected_month uint32
	expected_day   float64
	expected_year  int
}

var addGregorianDateTests = []addGregorianDateTest{
	addGregorianDateTest{2436116.31, 10, 4.81, 1957},
	addGregorianDateTest{1842713.0, 1, 27.5, 333},
	addGregorianDateTest{1507900.13, 5, 28.63, -584},
}

func TestGregorianDate(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	for _, test := range addGregorianDateTests {
		var m, d, y = GregorianDate(test.JD)
		if m != test.expected_month {
			t.Errorf("Output %d not equal to expected %d", m, test.expected_month)
		}
		if d != test.expected_day {
			t.Errorf("Output %f not equal to expected %f", d, test.expected_day)
		}
		if y != test.expected_year {
			t.Errorf("Output %d not equal to expected %d", y, test.expected_year)
		}
	}
}

//type addDaysBetweenDatesTest struct {
//	m1       uint32
//	d1       float64
//	y1       int
//	m2       uint32
//	d2       float64
//	y2       int
//	expected float64
//}
//
//var addDaysBetweenDatesTests = []addDaysBetweenDatesTest{
//	addDaysBetweenDatesTest{4, 20.0, 1910, 2, 9.0, 1986, 27689.0},
//}

func TestDaysBetweenDates(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	var days = DaysBetweenDates(4, 1910, 20.0, 2, 1986, 9.0)
	if days != 27689 {
		t.Fatal(days)
	}
}

func TestAdjustedDateByDays(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	var m, d, y = AdjustedDateByDays(7, 1991, 11.0, 10000)
	if m != 11 {
		t.Errorf("Output %d not equal to expected %d", m, 11)
	}
	if d != 26.0 {
		t.Errorf("Output %f not equal to expected %f", d, 26.0)
	}
	if y != 2018 {
		t.Errorf("Output %d not equal to expected %d", y, 2018)
	}
}

package datetime

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

type addDateOfEasterForGregorianTest struct {
	Y          int
	expected_M int
	expected_D int
}

var addDateOfEasterForGregorianTests = []addDateOfEasterForGregorianTest{
	addDateOfEasterForGregorianTest{1991, 3, 31},
	addDateOfEasterForGregorianTest{1992, 4, 19},
	addDateOfEasterForGregorianTest{1993, 4, 11},
	addDateOfEasterForGregorianTest{1954, 4, 18},
	addDateOfEasterForGregorianTest{2000, 4, 23},
	addDateOfEasterForGregorianTest{1818, 3, 22},
}

func TestDateOfEasterForGregorian(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	for _, test := range addDateOfEasterForGregorianTests {
		var m, d = DateOfEasterForGregorian(test.Y)
		//log.Debug("m = " + strconv.Itoa(m))
		//log.Debug("d = " + strconv.Itoa(d))
		if m != test.expected_M {
			t.Errorf("Month %d not equal to expected %d", m, test.expected_M)
		}
		if d != test.expected_D {
			t.Errorf("Day %d not equal to expected %d", d, test.expected_D)
		}
	}
}

type addDateOfEasterForJulianTest struct {
	Y          int
	expected_M int
	expected_D int
}

var addDateOfEasterForJulianTests = []addDateOfEasterForJulianTest{
	addDateOfEasterForJulianTest{179, 4, 12},
	addDateOfEasterForJulianTest{711, 4, 12},
	addDateOfEasterForJulianTest{1243, 4, 12},
}

func TestDateOfEasterForJulian(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	for _, test := range addDateOfEasterForJulianTests {
		var m, d = DateOfEasterForJulian(test.Y)
		//log.Debug("m = " + strconv.Itoa(m))
		//log.Debug("d = " + strconv.Itoa(d))
		if m != test.expected_M {
			t.Errorf("Month %d not equal to expected %d", m, test.expected_M)
		}
		if d != test.expected_D {
			t.Errorf("Day %d not equal to expected %d", d, test.expected_D)
		}
	}
}

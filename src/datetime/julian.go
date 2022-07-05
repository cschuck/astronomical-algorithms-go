package datetime

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"strconv"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func isLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 {
		return true
	} else if year%4 == 0 && year%100 == 0 && year%400 == 0 {
		return true
	} else {
		return false
	}
}

// JulianDay converts either a Gregorian or Julian Date to its corresponding Julian Day.
// M - month of the date
// Y - year of the date
// D - day of the month
// gregorian - true if the date is Gregorian, false if Julian
func JulianDay(M uint32, Y int, D float64, gregorian bool) float64 {
	var B = 0
	if gregorian {
		var A = uint32(Y / 100)
		log.Debug("A = " + strconv.Itoa(int(A)))
		B = 2 - int(A) + int(int(A)/4)
		log.Debug("B = " + strconv.Itoa(B))
	}
	var JD float64 = 0
	if M > 2 {
		JD = float64(uint32(365.25*float64(Y+4716))) + float64(uint32(30.6001*float64(M+1))) + D + float64(B) - 1524.5
	} else {
		var M1 = M + 12
		var Y1 = Y - 1
		JD = float64(uint32(365.25*float64(Y1+4716))) + float64(uint32(30.6001*float64(M1+1))) + D + float64(B) - 1524.5
	}
	return JD
}

// ModifiedJulianDay converts either a Gregorian or Julian Date to its corresponding Julian Day.
// M - month of the date
// Y - year of the date
// D - day of the month
// gregorian - true if the date is Gregorian, false if Julian
func ModifiedJulianDay(M uint32, Y int, D float64, gregorian bool) float64 {
	return JulianDay(M, Y, D, gregorian) - 2400000.5
}

func GregorianDate(JD float64) (uint32, float64, int) {
	var Z, F = math.Modf(JD + 0.5)
	Fs := fmt.Sprintf("%f", F)
	log.Debug("Z = " + strconv.Itoa(int(Z)) + " F = " + Fs)
	var A = 0
	if Z < 2299161 {
		A = int(Z)
	} else {
		var alpha = int((Z - 1867216.25) / 36524.25)
		A = int(Z) + 1 + alpha - int(alpha/4)
	}
	var B = A + 1524
	var C = int((float64(B) - 122.1) / 365.25)
	var D = int(365.25 * float64(C))
	var E = int((float64(B-D) / 30.6001))
	var d = roundFloat(float64(B)-float64(D)-float64(int(30.6001*float64(E)))+F, 2)
	var m = 0
	if E < 14 {
		m = E - 1
	} else if E == 14 || E == 15 {
		m = E - 13
	}
	var y = 0
	if m > 2 {
		y = C - 4716
	} else if m == 1 || m == 2 {
		y = C - 4715
	}
	return uint32(m), d, y
}

func DaysBetweenDates(M1 uint32, Y1 int, D1 float64, M2 uint32, Y2 int, D2 float64) float64 {
	var JD1 = JulianDay(M1, Y1, D1, true)
	var JD2 = JulianDay(M2, Y2, D2, true)
	return JD2 - JD1
}

func AdjustedDateByDays(M uint32, Y int, D float64, Days float64) (uint32, float64, int) {
	var JD = JulianDay(M, Y, D, true)
	return GregorianDate(JD + Days)
}

func DayOfWeek(M uint32, Y int, D float64) uint32 {
	var JD = JulianDay(M, Y, D, true)
	var JD1 = JD + 1.5
	return uint32(int(roundFloat(JD1, 0)) % 7)
}

func DayOfTheYear(M uint32, Y int, D float64) uint32 {
	var K = 2.0
	if isLeapYear(Y) {
		log.Debug("leap year")
		K = 1.0
	} else {
		log.Debug("not leap year")
	}
	//var I1 = roundFloat(float64(275*M)/9.0, 0)
	var I1 = int(float64(275*M) / 9.0)
	//log.Debug("I1 = " + fmt.Sprintf("%f", I1))
	//var I2 = roundFloat(float64(M+9)/12.0, 0)
	var I2 = int(float64(M+9) / 12.0)
	//log.Debug("I2 = " + fmt.Sprintf("%f", I2))
	var N = uint32(roundFloat(float64(I1)-float64(K)*float64(I2)+D-30.0, 0))
	log.Debug("N = " + strconv.Itoa(int(N)))
	return N
}

func MonthAndDay(N uint32, Y int) (uint32, uint32) {
	var K = 2
	if isLeapYear(Y) {
		K = 1
	}
	var M = uint32(1)
	if N >= 32 {
		M = uint32((9.0*(float64(K)+float64(N)))/275.0 + 0.98)
	}
	var D = uint32(int(N) - int(float64(275*M)/9.0) + K*int(float64(M+9)/12.0) + 30)
	return M, D
}

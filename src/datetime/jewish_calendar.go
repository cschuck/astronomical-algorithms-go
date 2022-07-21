package datetime

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func FifteenNisan(X int, isJulian bool) (uint32, uint32, int) {
	var C = int(X / 100)
	log.Debug("C = " + strconv.Itoa(C))
	var S = int(float64(3*C-5) / float64(4))
	if X < 1583 || isJulian {
		S = 0
	}
	log.Debug("S = " + strconv.Itoa(S))
	var A = X + 3760
	log.Debug("A = " + strconv.Itoa(A))
	var a = (12*X + 12) % 19
	log.Debug("a = " + strconv.Itoa(a))
	var b = X % 4
	log.Debug("b = " + strconv.Itoa(b))
	var Q = -1.904412361576 + 1.554241796621*float64(a) + 0.25*float64(b) - 0.003177794022*float64(X) + float64(S)
	Qs := fmt.Sprintf("%f", Q)
	log.Debug("Q = " + Qs)
	var j = (int(Q) + 3*X + 5*b + 2 - S) % 7
	log.Debug("j = " + strconv.Itoa(j))
	var r = Q - float64(int(Q))
	rs := fmt.Sprintf("%f", r)
	log.Debug("r = " + rs)
	var D = 0
	if j == 2 || j == 4 || j == 6 {
		D = int(Q) + 23
	} else if j == 1 && a > 6 && r >= 0.632870370 {
		D = int(Q) + 24
	} else if j == 0 && a > 11 && r >= 0.897723765 {
		D = int(Q) + 23
	} else {
		D = int(Q) + 22
	}
	if D <= 31 {
		return uint32(D), 3, A
	} else {
		return uint32(D - 31), 4, A
	}
}

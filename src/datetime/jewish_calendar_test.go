package datetime

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"testing"
)

func TestFifteenNisan(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	var X = 1990
	var isJulian = false
	var D, M, A = FifteenNisan(X, isJulian)
	log.Debug("D = " + strconv.Itoa(int(D)))
	log.Debug("M = " + strconv.Itoa(int(M)))
	log.Debug("A = " + strconv.Itoa(A))
}

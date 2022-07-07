package datetime

func DateOfEasterForGregorian(x int) (int, int) {
	//log.Debugf("x = " + strconv.Itoa(x))
	var a = x % 19
	//log.Debugf("a = " + strconv.Itoa(a))
	var b = x / 100
	//log.Debugf("b = " + strconv.Itoa(b))
	var c = x % 100
	//log.Debugf("c = " + strconv.Itoa(c))
	var d = b / 4
	//log.Debugf("d = " + strconv.Itoa(d))
	var e = b % 4
	//log.Debugf("e = " + strconv.Itoa(e))
	var f = (b + 8) / 25
	//log.Debugf("f = " + strconv.Itoa(f))
	var g = (b - f + 1) / 3
	//log.Debugf("g = " + strconv.Itoa(g))
	var h = (19*a + b - d - g + 15) % 30
	//log.Debugf("h = " + strconv.Itoa(h))
	var i = c / 4
	//log.Debugf("i = " + strconv.Itoa(i))
	var k = c % 4
	//log.Debugf("k = " + strconv.Itoa(k))
	var l = (32 + 2*e + 2*i - h - k) % 7
	//log.Debugf("l = " + strconv.Itoa(l))
	var m = (a + 11*h + 22*l) / 451
	//log.Debugf("m = " + strconv.Itoa(m))
	var n = (h + l - 7*m + 114) / 31
	//log.Debugf("n = " + strconv.Itoa(n))
	var p = (h + l - 7*m + 114) % 31
	//log.Debugf("p = " + strconv.Itoa(p))
	return n, p + 1
}

func DateOfEasterForJulian(x int) (int, int) {
	//log.Debugf("x = " + strconv.Itoa(x))
	var a = x % 4
	var b = x % 7
	var c = x % 19
	var d = (19*c + 15) % 30
	var e = (2*a + 4*b - d + 34) % 7
	var f = (d + e + 114) / 31
	var g = (d + e + 114) % 31
	return f, g + 1
}

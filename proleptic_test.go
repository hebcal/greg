package greg_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hebcal/greg"
	"github.com/stretchr/testify/assert"
)

func TestProlepticGreg2RD(t *testing.T) {
	assert := assert.New(t)
	rataDie := greg.ProlepticToRD(1995, time.December, 17)
	assert.Equal(int64(728644), rataDie)
	rataDie = greg.ProlepticToRD(1888, time.December, 31)
	assert.Equal(int64(689578), rataDie)
	rataDie = greg.ProlepticToRD(2005, time.April, 2)
	assert.Equal(int64(732038), rataDie)
}

func TestProlepticGreg2RDEarlyCE(t *testing.T) {
	assert := assert.New(t)
	rataDie := greg.ProlepticToRD(88, time.December, 30)
	assert.Equal(int64(32141), rataDie)
	rataDie = greg.ProlepticToRD(1, time.January, 1)
	assert.Equal(int64(1), rataDie)
}

func TestProlepticGreg2RDNegative(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int64(0), greg.ProlepticToRD(0, time.December, 31))
	assert.Equal(int64(-1), greg.ProlepticToRD(0, time.December, 30))
	assert.Equal(int64(-2), greg.ProlepticToRD(0, time.December, 29))
	assert.Equal(int64(-48), greg.ProlepticToRD(0, time.November, 13))
	assert.Equal(int64(-61), greg.ProlepticToRD(0, time.October, 31))
	assert.Equal(int64(-91), greg.ProlepticToRD(0, time.October, 1))
	assert.Equal(int64(-305), greg.ProlepticToRD(0, time.March, 1))
	assert.Equal(int64(-306), greg.ProlepticToRD(0, time.February, 29))
	assert.Equal(int64(-307), greg.ProlepticToRD(0, time.February, 28))
	assert.Equal(int64(-308), greg.ProlepticToRD(0, time.February, 27))
	assert.Equal(int64(-334), greg.ProlepticToRD(0, time.February, 1))
	assert.Equal(int64(-335), greg.ProlepticToRD(0, time.January, 31))
	assert.Equal(int64(-350), greg.ProlepticToRD(0, time.January, 16))
	assert.Equal(int64(-365), greg.ProlepticToRD(0, time.January, 1))
	assert.Equal(int64(-366), greg.ProlepticToRD(-1, time.December, 31))
}

func TestProlepticGreg2RDNegative2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int64(-396), greg.ProlepticToRD(-1, time.December, 1))
	assert.Equal(int64(-730), greg.ProlepticToRD(-1, time.January, 1))
	assert.Equal(int64(-1095), greg.ProlepticToRD(-2, time.January, 1))
	assert.Equal(int64(-1460), greg.ProlepticToRD(-3, time.January, 1))
	assert.Equal(int64(-36171), greg.ProlepticToRD(-99, time.December, 20))
	assert.Equal(int64(-365077), greg.ProlepticToRD(-999, time.June, 15))
	assert.Equal(int64(-36536), greg.ProlepticToRD(-100, time.December, 20))
}

func TestProlepticRD2Greg(t *testing.T) {
	assert := assert.New(t)
	year, month, day := greg.ProlepticFromRD(737553)
	assert.Equal(2020, year)
	assert.Equal(time.May, month)
	assert.Equal(8, day)
	year2, month2, day2 := greg.ProlepticFromRD(689578)
	assert.Equal(1888, year2)
	assert.Equal(time.December, month2)
	assert.Equal(31, day2)
	gy, gm, gd := greg.ProlepticFromRD(732038)
	assert.Equal(2005, gy)
	assert.Equal(time.April, gm)
	assert.Equal(2, gd)
}

func TestProlepticRD2Greg88ce(t *testing.T) {
	assert := assert.New(t)
	var year int
	var month time.Month
	var day int
	year, month, day = greg.ProlepticFromRD(32141)
	assert.Equal(88, year)
	assert.Equal(time.December, month)
	assert.Equal(30, day)
	year, month, day = greg.ProlepticFromRD(32142)
	assert.Equal(88, year)
	assert.Equal(time.December, month)
	assert.Equal(31, day)
	year, month, day = greg.ProlepticFromRD(32143)
	assert.Equal(89, year)
	assert.Equal(time.January, month)
	assert.Equal(1, day)
}

func TestProlepticRD2Greg1ce(t *testing.T) {
	assert := assert.New(t)
	year, month, day := greg.ProlepticFromRD(1)
	assert.Equal(1, year)
	assert.Equal(time.January, month)
	assert.Equal(1, day)
}

func TestProlepticRD2GregNegative(t *testing.T) {
	assert := assert.New(t)
	var year int
	var month time.Month
	var day int

	year, month, day = greg.ProlepticFromRD(-730)
	assert.Equal(-1, year)
	assert.Equal(time.January, month)
	assert.Equal(1, day)

	year, month, day = greg.ProlepticFromRD(-36536)
	assert.Equal(-100, year)
	assert.Equal(time.December, month)
	assert.Equal(20, day)

	year, month, day = greg.ProlepticFromRD(0)
	assert.Equal(0, year)
	assert.Equal(time.December, month)
	assert.Equal(31, day)
	year, month, day = greg.ProlepticFromRD(-1)
	assert.Equal(0, year)
	assert.Equal(time.December, month)
	assert.Equal(30, day)
	year, month, day = greg.ProlepticFromRD(-2)
	assert.Equal(0, year)
	assert.Equal(time.December, month)
	assert.Equal(29, day)
	year, month, day = greg.ProlepticFromRD(-48)
	assert.Equal(0, year)
	assert.Equal(time.November, month)
	assert.Equal(13, day)
	year, month, day = greg.ProlepticFromRD(-61)
	assert.Equal(0, year)
	assert.Equal(time.October, month)
	assert.Equal(31, day)

	year, month, day = greg.ProlepticFromRD(-91)
	assert.Equal(0, year)
	assert.Equal(time.October, month)
	assert.Equal(1, day)

	year, month, day = greg.ProlepticFromRD(-92)
	assert.Equal(0, year)
	assert.Equal(time.September, month)
	assert.Equal(30, day)

	year, month, day = greg.ProlepticFromRD(-365)
	assert.Equal(0, year)
	assert.Equal(time.January, month)
	assert.Equal(1, day)

	year, month, day = greg.ProlepticFromRD(-366)
	assert.Equal(-1, year)
	assert.Equal(time.December, month)
	assert.Equal(31, day)
}

func ExampleProlepticToRD() {
	rataDie := greg.ProlepticToRD(1995, time.December, 17)
	fmt.Println(rataDie)
	// Output: 728644
}

func ExampleProlepticFromRD() {
	year, month, day := greg.ProlepticFromRD(737553)
	fmt.Println(year, month, day)
	// Output: 2020 May 8
}

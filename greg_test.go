package greg_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hebcal/greg"
	"github.com/stretchr/testify/assert"
)

func TestGreg2RD(t *testing.T) {
	assert := assert.New(t)
	rataDie := greg.ToRD(1995, time.December, 17)
	assert.Equal(int64(728644), rataDie)
	rataDie = greg.ToRD(1888, time.December, 31)
	assert.Equal(int64(689578), rataDie)
	rataDie = greg.ToRD(2005, time.April, 2)
	assert.Equal(int64(732038), rataDie)
}

func TestGreg2RDEarlyCE(t *testing.T) {
	assert := assert.New(t)
	rataDie := greg.ToRD(88, time.December, 30)
	assert.Equal(int64(32139), rataDie)
	rataDie = greg.ToRD(1, time.January, 1)
	assert.Equal(int64(-1), rataDie)
}

func TestGreg2RDNegative(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int64(0), greg.ToRD(1, time.January, 2))
	assert.Equal(int64(-1), greg.ToRD(1, time.January, 1))
	assert.Equal(int64(-2), greg.ToRD(-1, time.December, 31))
	assert.Equal(int64(-3), greg.ToRD(-1, time.December, 30))
	assert.Equal(int64(-4), greg.ToRD(-1, time.December, 29))
	assert.Equal(int64(-50), greg.ToRD(-1, time.November, 13))
	assert.Equal(int64(-63), greg.ToRD(-1, time.October, 31))
	assert.Equal(int64(-93), greg.ToRD(-1, time.October, 1))
	assert.Equal(int64(-307), greg.ToRD(-1, time.March, 1))
	assert.Equal(int64(-308), greg.ToRD(-1, time.February, 29))
	assert.Equal(int64(-309), greg.ToRD(-1, time.February, 28))
	assert.Equal(int64(-310), greg.ToRD(-1, time.February, 27))
	assert.Equal(int64(-336), greg.ToRD(-1, time.February, 1))
	assert.Equal(int64(-337), greg.ToRD(-1, time.January, 31))
	assert.Equal(int64(-352), greg.ToRD(-1, time.January, 16))
	assert.Equal(int64(-367), greg.ToRD(-1, time.January, 1))
	assert.Equal(int64(-368), greg.ToRD(-2, time.December, 31))
}

func TestGreg2RDNegative2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int64(-398), greg.ToRD(-2, time.December, 1))
	assert.Equal(int64(-732), greg.ToRD(-2, time.January, 1))
	assert.Equal(int64(-1097), greg.ToRD(-3, time.January, 1))
	assert.Equal(int64(-1462), greg.ToRD(-4, time.January, 1))
	assert.Equal(int64(-36173), greg.ToRD(-100, time.December, 20))
	assert.Equal(int64(-365086), greg.ToRD(-1000, time.June, 15))
	assert.Equal(int64(-36538), greg.ToRD(-101, time.December, 20))
	assert.Equal(int64(-36892), greg.ToRD(-101, time.January, 1))
}

func TestRD2Greg(t *testing.T) {
	assert := assert.New(t)
	year, month, day := greg.FromRD(737553)
	assert.Equal(2020, year)
	assert.Equal(time.May, month)
	assert.Equal(8, day)
	year2, month2, day2 := greg.FromRD(689578)
	assert.Equal(1888, year2)
	assert.Equal(time.December, month2)
	assert.Equal(31, day2)
	gy, gm, gd := greg.FromRD(732038)
	assert.Equal(2005, gy)
	assert.Equal(time.April, gm)
	assert.Equal(2, gd)
}

func TestRD2Greg88ce(t *testing.T) {
	assert := assert.New(t)
	var year int
	var month time.Month
	var day int
	year, month, day = greg.FromRD(32139)
	assert.Equal(88, year)
	assert.Equal(time.December, month)
	assert.Equal(30, day)
	year, month, day = greg.FromRD(32140)
	assert.Equal(88, year)
	assert.Equal(time.December, month)
	assert.Equal(31, day)
	year, month, day = greg.FromRD(32141)
	assert.Equal(89, year)
	assert.Equal(time.January, month)
	assert.Equal(1, day)
}

func TestRD2GregAtTransition(t *testing.T) {
	assert := assert.New(t)
	var year int
	var month time.Month
	var day int
	year, month, day = greg.FromRD(639796)
	assert.Equal(1752, year)
	assert.Equal(time.September, month)
	assert.Equal(2, day)
	year, month, day = greg.FromRD(639797)
	assert.Equal(1752, year)
	assert.Equal(time.September, month)
	assert.Equal(14, day)
}

func TestGreg2RDAtTransition(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int64(639796), greg.ToRD(1752, time.September, 2))
	assert.Equal(int64(639797), greg.ToRD(1752, time.September, 14))
	assert.Equal(int64(639798), greg.ToRD(1752, time.September, 15))
}

func TestRD2Greg1ce(t *testing.T) {
	assert := assert.New(t)
	var year int
	var month time.Month
	var day int
	year, month, day = greg.FromRD(-1)
	assert.Equal(1, year)
	assert.Equal(time.January, month)
	assert.Equal(1, day)
	year, month, day = greg.FromRD(0)
	assert.Equal(1, year)
	assert.Equal(time.January, month)
	assert.Equal(2, day)
	year, month, day = greg.FromRD(1)
	assert.Equal(1, year)
	assert.Equal(time.January, month)
	assert.Equal(3, day)
}

func TestRD2GregNegative(t *testing.T) {
	assert := assert.New(t)
	var year int
	var month time.Month
	var day int

	year, month, day = greg.FromRD(-732)
	assert.Equal(-2, year)
	assert.Equal(time.January, month)
	assert.Equal(1, day)

	year, month, day = greg.FromRD(-36538)
	assert.Equal(-101, year)
	assert.Equal(time.December, month)
	assert.Equal(20, day)

	year, month, day = greg.FromRD(-2)
	assert.Equal(-1, year)
	assert.Equal(time.December, month)
	assert.Equal(31, day)
	year, month, day = greg.FromRD(-3)
	assert.Equal(-1, year)
	assert.Equal(time.December, month)
	assert.Equal(30, day)
	year, month, day = greg.FromRD(-4)
	assert.Equal(-1, year)
	assert.Equal(time.December, month)
	assert.Equal(29, day)
	year, month, day = greg.FromRD(-50)
	assert.Equal(-1, year)
	assert.Equal(time.November, month)
	assert.Equal(13, day)
	year, month, day = greg.FromRD(-63)
	assert.Equal(-1, year)
	assert.Equal(time.October, month)
	assert.Equal(31, day)

	year, month, day = greg.FromRD(-93)
	assert.Equal(-1, year)
	assert.Equal(time.October, month)
	assert.Equal(1, day)

	year, month, day = greg.FromRD(-94)
	assert.Equal(-1, year)
	assert.Equal(time.September, month)
	assert.Equal(30, day)

	year, month, day = greg.FromRD(-367)
	assert.Equal(-1, year)
	assert.Equal(time.January, month)
	assert.Equal(1, day)

	year, month, day = greg.FromRD(-368)
	assert.Equal(-2, year)
	assert.Equal(time.December, month)
	assert.Equal(31, day)
}

func ExampleDaysIn() {
	days := greg.DaysIn(time.February, 2004)
	fmt.Println(days)
	// Output: 29
}

func ExampleDateToRD() {
	t := time.Date(2014, time.February, 19, 0, 0, 0, 0, time.UTC)
	rataDie := greg.DateToRD(t)
	fmt.Println(rataDie)
	// Output: 735283
}

func ExampleToRD() {
	rataDie := greg.ToRD(1995, time.December, 17)
	fmt.Println(rataDie)
	// Output: 728644
}

func ExampleFromRD() {
	year, month, day := greg.FromRD(737553)
	fmt.Println(year, month, day)
	// Output: 2020 May 8
}

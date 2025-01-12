// Hebcal's greg package converts between Gregorian dates
// and R.D. (Rata Die) day numbers.
// An R.D. number consists of an absolute number equal the number of days
// that elapsed from a fixed point in time. We use negative numbers to represent
// days before that fixed point. For this library, we use the Gregorian date 12/31/1 BCE
// as our fixed point.
package greg

// Hebcal - A Jewish Calendar Generator
// Copyright (c) 2022 Michael J. Radwin
// Derived from original C version, Copyright (C) 1994-2004 Danny Sadinoff
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

import (
	"math"
	"time"
)

// 1-based month lengths
var monthLen = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// DaysIn returns the number of days in the Gregorian month.
func DaysIn(m time.Month, year int) int {
	if m == time.February && IsLeapYear(year) {
		return 29
	}
	return monthLen[m]
}

// Returns true if the Gregorian year is a leap year.
func IsLeapYear(year int) bool {
	gyear := int64(year)
	if mod(gyear, 4) == 0 {
		n := mod(gyear, 400)
		if n != 100 && n != 200 && n != 300 {
			return true
		}
	}
	return false
}

// Converts Gregorian date to absolute R.D. (Rata Die) days.
// Hours, minutes and seconds are ignored
func DateToRD(t time.Time) int64 {
	year, month, day := t.Date()
	abs := ToRD(year, month, day)
	return abs
}

// This function helps with a trick we later use to find the absolute number
// of the current day this year. The idea is that if we are in Jan or Feb,
// our approximation is exactly accurate. But for later months we need to subtract 1
// if it's a leap year, and 2 if it isn't
func monthOffset(year int, month time.Month) int {
	if month <= time.February {
		return 0
	} else if IsLeapYear(year) {
		return -1
	} else {
		return -2
	}
}

// Converts Gregorian date to absolute R.D. (Rata Die) days.
//
// Panics if Gregorian year is 0.
func ToRD(year int, month time.Month, day int) int64 {
	py := int64(year - 1)
	abs := 365*py + // days up to preceding year
		quotient(py, 4) - // add in Julian leap years
		quotient(py, 100) + // subtract out century leap years
		quotient(py, 400) + // add in Gregorian leap years
		quotient((367*int64(month)-362), 12) + // add in the days so far this year
		int64(monthOffset(year, month)) + int64(day)
	return abs
}

// the next two functions allow us to deal with negative R.D. numbers
func mod(x, y int64) int64 {
	X := float64(x)
	Y := float64(y)
	return int64(X - Y*math.Floor(X/Y))
}

func quotient(x, y int64) int64 {
	return int64(math.Floor(float64(x) / float64(y)))
}

// Finds the year in which a given R.D. occurs
func yearFromFixed(rataDie int64) int {
	l0 := int64(rataDie) - 1     // subtract 1 because we are counting from Gregorian date 12/31/1 BCE
	n400 := quotient(l0, 146097) // number of 400 years periods
	d1 := mod(l0, 146097)        // days into current period
	n100 := quotient(d1, 36524)  // number of 100 year periods
	d2 := mod(d1, 36524)
	n4 := quotient(d2, 1461) // number of 4 year periods
	d3 := mod(d2, 1461)
	n1 := quotient(d3, 365)                 // number of years into current period
	year := 400*n400 + 100*n100 + 4*n4 + n1 // total years
	yy := int(year)

	// if we didn't get a 4-year block (with a leap day in it), but we did get 4 separate years
	// it must be December 31 on the previous year (because there was a leap day). So don't add 1
	if n100 == 4 || n1 == 4 {
		return yy
	}

	// otherwise it is the next year (because there is no year 0, so generally add 1)
	return yy + 1
}

/*
Converts from Rata Die (R.D. number) to Gregorian date.

See the footnote on page 384 of “Calendrical Calculations, Part II:
Three Historical Calendars” by E. M. Reingold,  N. Dershowitz, and S. M.
Clamen, Software--Practice and Experience, Volume 23, Number 4
(April, 1993), pages 383-404 for an explanation.
*/
func FromRD(rataDie int64) (int, time.Month, int) {
	year := yearFromFixed(rataDie) // get the year we are in
	var correction int64
	if rataDie < ToRD(year, time.March, 1) {
		correction = 0
	} else if IsLeapYear(year) {
		correction = 1
	} else {
		correction = 2
	}
	priorDays := rataDie - ToRD(year, time.January, 1)    // how many days into the year are we?
	month := quotient(12*(priorDays+correction)+373, 367) // trick to find month
	day := rataDie - ToRD(year, time.Month(month), 1) + 1 // find remaining days
	return year, time.Month(month), int(day)
}

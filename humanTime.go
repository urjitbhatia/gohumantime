package gohumantime

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	unitRegexpPattern      = "(second|minute|hour|day|week|month|year)s?"
	connectorRegexpPattern = "and|,"
	Second                 = 1000
	Minute                 = Second * 60
	Hour                   = Minute * 60
	Day                    = Hour * 24
	Week                   = Day * 7
	Month                  = Day * 30
	Year                   = Day * 365
)

/*
 * humanTime holds internal data and methods to help converting a human readable time string into milliseconds
 */
type humanTime struct {
	languageMap     map[string]string // Word to number map
	unitRegexp      *regexp.Regexp    // Cached unit regex
	connectorRegexp *regexp.Regexp    // Cached connector regex
}

/*
 * processUnits converts time unit words like "minute" into the correct millisecond multiplier
 */
func processUnits(time string) (int, error) {

	if strings.TrimSpace(time) == "" {
		return 0, nil
	}

	fields := strings.Fields(time)
	if len(fields) < 2 {
		return 0, errors.New("No usable time literals found")
	}
	num, err := strconv.ParseFloat(fields[0], 10)
	if err != nil {
		num = 1
	}
	unit := fields[1]
	var unitNum float64
	switch unit {
	case "second":
		unitNum = Second
	case "minute":
		unitNum = Minute
	case "hour":
		unitNum = Hour
	case "day":
		unitNum = Day
	case "week":
		unitNum = Week
	case "month":
		unitNum = Month
	case "year":
		unitNum = Year
	}

	return int(unitNum * num), nil
}

/*
 * ToMilliseconds converts a human readable interval string into milliseconds.
 * Example: ToMilliseconds("three minutes and five seconds") returns 3 * 60 * 1000 + 5 * 1000
 */
func ToMilliseconds(humanReadableTime string) (int, error) {
	return humanTime{map[string]string{
		"one":     "1",
		"two":     "2",
		"three":   "3",
		"four":    "4",
		"five":    "5",
		"six":     "6",
		"seven":   "7",
		"eight":   "8",
		"nine":    "9",
		"ten":     "10",
		"fifteen": "15",
		"twenty":  "20",
		"thirty":  "30",
		"forty":   "40",
		"fifty":   "50",
		"sixty":   "60",
		"seventy": "70",
		"eighty":  "80",
		"ninety":  "90",
		"hundred": "100",
	},
		regexp.MustCompile(unitRegexpPattern),
		regexp.MustCompile(connectorRegexpPattern),
	}.toMilliseconds(humanReadableTime)
}

/*
 * toMilliseconds converts a humanReadableTime string to milliseconds
 */
func (h humanTime) toMilliseconds(humanReadableTime string) (sum int, err error) {

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered panic in humanTime::ToMilliseconds", r)
			err = errors.New("Malformed input")
		}
	}()

	if strings.TrimSpace(humanReadableTime) == "" {
		return 0, nil
	}
	timeString := h.wordNumbersToDecimals(strings.ToLower(humanReadableTime))
	timeString = h.unitRegexp.ReplaceAllString(timeString, "$1,")
	for _, s := range h.connectorRegexp.Split(timeString, -1) {
		s, err := processUnits(s)
		if err == nil {
			sum += s
		}
	}
	return sum, err
}

/*
 * wordNumbersToDecimals replaces word numbers like "one", "two" into numeric literals like "1", "2" etc
 */
func (h humanTime) wordNumbersToDecimals(timeString string) string {

	fields := strings.Fields(timeString)
	for _, f := range fields {
		if val, ok := h.languageMap[f]; ok {
			timeString = strings.Replace(timeString, f, val, 1)
		}
	}
	return timeString
}

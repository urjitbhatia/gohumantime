package main

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
	second                 = 1000
	minute                 = second * 60
	hour                   = minute * 60
	day                    = hour * 24
	week                   = day * 7
	month                  = day * 30
	year                   = day * 365
)

/*
 * HumanInterval holds internal data and methods to help converting a human readable time string into milliseconds
 */
type humanInterval struct {
	languageMap     map[string]int // Word to number map
	unitRegexp      *regexp.Regexp // Cached unit regex
	connectorRegexp *regexp.Regexp // Cached connector regex
}

/*
 * processUnits converts time unit words like "minute" into the correct millisecond multiplier
 */
func processUnits(time string) int {

	if strings.TrimSpace(time) == "" {
		return 0
	}

	fields := strings.Fields(time)
	num, err := strconv.Atoi(fields[0])
	unit := fields[1]
	if err != nil {
		num = 1
	}
	var unitNum int
	switch unit {
	case "second":
		unitNum = second
	case "minute":
		unitNum = minute
	case "hour":
		unitNum = hour
	case "day":
		unitNum = day
	case "week":
		unitNum = week
	case "month":
		unitNum = month
	case "year":
		unitNum = year
	}

	return unitNum * num
}

/*
 * NewHumanInterval
 */
func ToMilliseconds(humanReadableTime string) (int, error) {
	return humanInterval{map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
	},
		regexp.MustCompile(unitRegexpPattern),
		regexp.MustCompile(connectorRegexpPattern),
	}.toMilliseconds(humanReadableTime)
}

/*
 * toMilliseconds converts a humanReadableTime string to milliseconds
 */
func (h humanInterval) toMilliseconds(humanReadableTime string) (sum int, err error) {

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered panic in HumanInterval::ToMilliseconds", r)
			err = errors.New("Malformed input")
		}
	}()

	if strings.TrimSpace(humanReadableTime) == "" {
		return 0, errors.New("Cannot parse empty humanReadable value")
	}

	timeString := h.wordNumbersToDecimals(humanReadableTime)
	timeString = h.unitRegexp.ReplaceAllString(timeString, "$1,")
	for _, s := range h.connectorRegexp.Split(timeString, -1) {
		sum += processUnits(s)
	}
	return sum, err
}

/*
 * wordNumbersToDecimals replaces word numbers like "one", "two" into numeric literals like "1", "2" etc
 */
func (h humanInterval) wordNumbersToDecimals(time string) string {

	fields := strings.Fields(time)
	for _, f := range fields {
		if val, ok := h.languageMap[f]; ok {
			var matchStr string
			if val > 1 {
				matchStr = strconv.Itoa(val)
			} else {
				matchStr = "" // ignore 1 since its unity
			}
			time = strings.Replace(time, f, matchStr, 1)
		}
	}
	return time
}

func main() {
	val, err := ToMilliseconds("1 minute2 second 22 years")
	if err == nil {
		log.Printf("valuie: %d", val)
	} else {
		log.Printf("errir:", err)
	}
}

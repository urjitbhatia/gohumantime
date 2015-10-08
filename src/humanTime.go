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
)

type HumanInterval struct {
	languageMap     map[string]int
	unitRegexp      *regexp.Regexp
	connectorRegexp *regexp.Regexp
}

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
		unitNum = 1000
	case "minute":
		unitNum = 1000 * 60
	case "hour":
		unitNum = 1000 * 60 * 60
	case "day":
		unitNum = 1000 * 60 * 60 * 24
	case "week":
		unitNum = 1000 * 60 * 60 * 24 * 7
	case "month":
		unitNum = 1000 * 60 * 60 * 24 * 30
	case "year":
		unitNum = 1000 * 60 * 60 * 24 * 365
	}

	return unitNum * num
}

func NewHumanInterval() HumanInterval {
	return HumanInterval{map[string]int{
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
	}
}

func (h HumanInterval) ToMilliseconds(humanReadableTime string) (int, error) {

	if strings.TrimSpace(humanReadableTime) == "" {
		return 0, errors.New("Cannot parse empty humanReadable value")
	}

	timeString := h.wordNumbersToDecimals(humanReadableTime)
	timeString = h.unitRegexp.ReplaceAllString(timeString, "$1,")
	sum := 0
	for _, s := range h.connectorRegexp.Split(timeString, -1) {
		sum += processUnits(s)
	}
	return sum, nil
}

func (hi HumanInterval) wordNumbersToDecimals(time string) string {

	fields := strings.Fields(time)
	for _, f := range fields {
		if val, ok := hi.languageMap[f]; ok {
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
	hi := NewHumanInterval()
	val, err := hi.ToMilliseconds("1 minute, 0 minutes")
	if err == nil {
		log.Printf("valuie: %d", val)
	} else {
		log.Printf("errir:", err)
	}
}

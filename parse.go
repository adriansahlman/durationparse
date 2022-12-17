package durationparse

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	Nanosecond  = time.Nanosecond
	Microsecond = time.Microsecond
	Millisecond = time.Millisecond
	Second      = time.Second
	Minute      = time.Minute
	Hour        = time.Hour
	Day         = 24 * Hour
	Week        = 7 * Day
	Month       = 30 * Day
	Year        = 365 * Day
)

var unitPatterns = map[time.Duration]*regexp.Regexp{
	Nanosecond:  regexp.MustCompile(`n(ano)?s(ec(onds?)?)?`),
	Microsecond: regexp.MustCompile(`([μµu]|(micro))s(ec(onds?)?)?`),
	Millisecond: regexp.MustCompile(`m(ill?i)?s(ec(onds?)?)?`),
	Second:      regexp.MustCompile(`s(ec(onds?)?)?`),
	Minute:      regexp.MustCompile(`m(in(utes?)?)?`),
	Hour:        regexp.MustCompile(`h(r|(ours?)?)`),
	Day:         regexp.MustCompile(`d(ays?)?`),
	Week:        regexp.MustCompile(`w(eeks?)?`),
	Month:       regexp.MustCompile(`months?`),
	Year:        regexp.MustCompile(`y(ears?)?`),
}

var numberPattern = regexp.MustCompile(`\d+(\.\d+)?`)
var valuePatterns map[time.Duration]*regexp.Regexp
var validationPattern *regexp.Regexp

func init() {
	unitExpressions := make([]string, 0, len(unitPatterns))
	valuePatterns = make(map[time.Duration]*regexp.Regexp)
	for key, value := range unitPatterns {
		valuePatterns[key] = regexp.MustCompile(
			fmt.Sprintf(
				`%s\s*%s(\.|,|\s|\d|$)`,
				numberPattern.String(),
				value.String(),
			),
		)
		unitExpressions = append(unitExpressions, value.String())
	}
	validationPattern = regexp.MustCompile(
		fmt.Sprintf(
			`^\s*(\d+(\.\d+)?\s*(%s)\s*(\.|,|\s*|$)?\s*)*$`,
			strings.Join(
				unitExpressions,
				"|",
			),
		),
	)
}

// Parse a duration string.
// When error is not nil the duration returned
// contains all the time that could be parsed
// instead of always being 0.
func Parse(
	input string,
) (time.Duration, error) {
	var duration time.Duration
	for unit, pattern := range valuePatterns {
		if match := pattern.FindString(input); match != "" {
			factor, err := strconv.ParseFloat(
				numberPattern.FindString(match),
				64,
			)
			if err != nil {
				return 0, err
			}
			duration += time.Duration(factor * float64(unit))
		}
	}
	if !validationPattern.MatchString(input) {
		return duration, fmt.Errorf("not a valid duration format")
	}
	return duration, nil
}

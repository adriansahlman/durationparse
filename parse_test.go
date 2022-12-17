package durationparse_test

import (
	"testing"
	"time"

	"github.com/adriansahlman/durationparse"
)

func TestParseSucceeds(t *testing.T) {
	for _, testCase := range []struct {
		input    string
		duration time.Duration
	}{
		{
			input:    "1d",
			duration: durationparse.Day,
		},
		{
			input:    "1h21m3ns",
			duration: durationparse.Hour + 21*durationparse.Minute + 3*durationparse.Nanosecond,
		},
		{
			input:    "1 years 2d ,5 min",
			duration: durationparse.Year + 2*durationparse.Day + 5*durationparse.Minute,
		},
	} {
		t.Run(
			testCase.input,
			func(t *testing.T) {
				got, err := durationparse.Parse(testCase.input)
				if err != nil {
					t.Fatalf(
						"unexpected error for input %q: %v",
						testCase.input,
						err,
					)
				}
				if got != testCase.duration {
					t.Errorf(
						"expected duration %d, got %d",
						testCase.duration,
						got,
					)
				}
			},
		)
	}

}

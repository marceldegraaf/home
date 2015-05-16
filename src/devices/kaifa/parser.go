package kaifa

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
)

var (
	totalInLowRegexp    = regexp.MustCompile("1.8.1\\(([0-9.]+)")
	totalInHighRegexp   = regexp.MustCompile("1.8.2\\(([0-9.]+)")
	totalOutLowRegexp   = regexp.MustCompile("2.8.1\\(([0-9.]+)")
	totalOutHighRegexp  = regexp.MustCompile("2.8.2\\(([0-9.]+)")
	currentTariffRegexp = regexp.MustCompile("96.14.0\\(([0-9]+)")
	currentlyInRegexp   = regexp.MustCompile("1.7.0\\(([0-9.]+)")
	currentlyOutRegexp  = regexp.MustCompile("2.7.0\\(([0-9.]+)")
	timestampRegexp     = regexp.MustCompile("1.0.0\\(([0-9]+)")
)

type message struct {
	body string
}

func parse(m string) (*Point, error) {
	msg := message{m}

	// Sometimes a message is empty; probably due to a
	// bug in the serial code. We detect that here, and
	// stop parsing if the timestamp is empty.
	if msg.findFirstMatch(timestampRegexp) == "" {
		return nil, errors.New("message was empty")
	}

	return &Point{
		CurrentlyIn:  msg.matchAsFloat(currentlyInRegexp),
		CurrentlyOut: msg.matchAsFloat(currentlyOutRegexp),
		TotalInLow:   msg.matchAsFloat(totalInLowRegexp),
		TotalInHigh:  msg.matchAsFloat(totalInHighRegexp),
		TotalOutLow:  msg.matchAsFloat(totalOutLowRegexp),
		TotalOutHigh: msg.matchAsFloat(totalOutHighRegexp),
		Tariff:       msg.matchAsInt(currentTariffRegexp),
		Timestamp:    msg.parsedTimestamp(timestampRegexp),
	}, nil
}

func (m message) findFirstMatch(r *regexp.Regexp) string {
	match := r.FindStringSubmatch(m.body)

	if len(match) == 0 {
		return ""
	}

	return match[1]
}

func (m message) matchAsInt(r *regexp.Regexp) int64 {
	match := m.findFirstMatch(r)

	i, err := strconv.ParseInt(match, 0, 0)
	if err != nil {
		log.Errorf("could not parse %q to int: %s", match, err)
		return 0
	}

	return i
}

func (m message) matchAsFloat(r *regexp.Regexp) float64 {
	match := m.findFirstMatch(r)

	flt, err := strconv.ParseFloat(match, 64)
	if err != nil {
		log.Errorf("could not parse %q to float64: %s", match, err)
		return 0
	}

	return flt
}

func (m message) parsedTimestamp(r *regexp.Regexp) time.Time {
	match := m.findFirstMatch(r)

	t, err := time.Parse("060102150405", match)
	if err != nil {
		log.Errorf("could not parse time %q: %s", match, err)
		return time.Now()
	}

	return t
}

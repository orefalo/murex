package mkarray

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	rxEngAlphaLower *regexp.Regexp = regexp.MustCompile(`^[a-z]$`)
	rxEngAlphaUpper *regexp.Regexp = regexp.MustCompile(`^[A-Z]$`)
)

func rangeToArray(b []byte) ([]string, error) {
	split := strings.Split(string(b), "..")
	if len(split) > 2 {
		return nil, errors.New("Invalid syntax. Too many double periods, `..`, in range`" + string(b) + "`. Please escape periods, `\\.`, if you wish to include period in your range.")
	}

	if len(split) < 2 {
		return nil, errors.New("Invalid syntax. Range periods, `..`, found but cannot determine start and end range in `" + string(b) + "`.")
	}

	i1, e1 := strconv.Atoi(split[0])
	i2, e2 := strconv.Atoi(split[1])

	if e1 == nil && e2 == nil {
		switch {
		case i1 < i2:
			a := make([]string, i2-i1+1)
			for i := range a {
				a[i] = strconv.Itoa(i + i1)
			}
			return a, nil

		case i1 > i2:
			a := make([]string, i1-i2+1)
			for i := range a {
				a[i] = strconv.Itoa(i1 - i)
			}
			return a, nil

		default:
			return nil, errors.New("Invalid range. Start and end of range are the same in `" + string(b) + "`.")
		}
	}

	if rxEngAlphaLower.MatchString(split[0]) && rxEngAlphaLower.MatchString(split[1]) {
		switch {
		case split[0] < split[1]:
		case split[0] > split[1]:
		default:
		}
		return nil, errors.New("TODO: Code me")
	}

	if rxEngAlphaUpper.MatchString(split[0]) && rxEngAlphaUpper.MatchString(split[1]) {
		switch {
		case split[0] < split[1]:
		case split[0] > split[1]:
		default:
		}
		return nil, errors.New("TODO: Code me")
	}

	// Mapped lists. See consts.go
	c := getCase(split[0])
	start := strings.ToLower(split[0])
	end := strings.ToLower(split[1])
	for i := range mapRanges {
		matched, array := mapArray(mapRanges[i][start], mapRanges[i][end], mapRanges[i], c)
		if matched {
			return array, nil
		}
	}

	return nil, errors.New("Unable to auto-detect range in `" + string(b) + "`.")
}

func mapArray(start, end int, constMap map[string]int, c int) (matched bool, array []string) {
	if start == 0 || end == 0 {
		return
	}

	matched = true

	consts := make([]string, len(constMap))
	for s, i := range constMap {
		consts[i-1] = setCase(s, c)
	}

	switch {
	case start < end:
		array = consts[start-1 : end]

	case start >= end:
		array = consts[start-1:]
		array = append(array, consts[:end]...)
	}

	return
}
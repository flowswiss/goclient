package flow

import (
	"fmt"
	"regexp"
	"strconv"
)

var identifierRegex = regexp.MustCompile("^flow://(\\d+)$")

type Id uint

func (i Id) String() string {
	return fmt.Sprintf("flow://%d", i)
}

func (i Id) Valid() bool {
	return i != 0
}

func ParseIdentifier(identifier string) Id {
	match := identifierRegex.FindStringSubmatch(identifier)
	if len(match) == 0 {
		return 0
	}

	val, err := strconv.ParseInt(match[1], 10, 32)
	if err != nil {
		return 0
	}

	return Id(val)
}

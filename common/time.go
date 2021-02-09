package common

import (
	"encoding/json"
	"time"
)

const (
	legacyTimeFormat = "2006-01-02T15:04:05-07:00"
	timeFormat       = "2006-01-02T15:04:05-0700"
)

type Time time.Time

func (t Time) AsTime() time.Time {
	return time.Time(t)
}

func (t Time) String() string {
	return t.AsTime().Format(timeFormat)
}

func (t *Time) UnmarshalJSON(b []byte) error {
	var formatted string

	err := json.Unmarshal(b, &formatted)
	if err != nil {
		return err
	}

	val, err := time.Parse(timeFormat, formatted)
	if _, ok := err.(*time.ParseError); ok {
		val, err = time.Parse(legacyTimeFormat, formatted)
	}

	if err != nil {
		return err
	}

	*t = Time(val)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

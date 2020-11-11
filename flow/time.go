package flow

import (
	"encoding/json"
	"time"
)

const (
	dateTimeFormat = "2006-01-02T15:04:05-0700"
)

type DateTime time.Time

func (d DateTime) Time() time.Time {
	return time.Time(d)
}

func (d DateTime) String() string {
	return d.Time().Format(dateTimeFormat)
}

func (d *DateTime) UnmarshalJSON(b []byte) error {
	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	t, err := time.Parse(dateTimeFormat, s)
	if err != nil {
		return err
	}

	*d = DateTime(t)
	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

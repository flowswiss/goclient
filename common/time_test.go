package common

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	now := time.Date(2020, time.April, 15, 17, 39, 51, 0, time.UTC)

	timeStr := now.Format(timeFormat)
	timeJSON, err := json.Marshal(timeStr)
	if err != nil {
		t.Fatal(err)
	}

	var parsed Time
	err = json.Unmarshal(timeJSON, &parsed)
	if err != nil {
		t.Fatal(err)
	}

	if !parsed.AsTime().Equal(now) {
		t.Errorf("expected parsed time to equal %v, got %v", now, parsed.AsTime())
	}

	timeStr = now.Format(legacyTimeFormat)
	timeJSON, err = json.Marshal(timeStr)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(timeJSON, &parsed)
	if err != nil {
		t.Fatal(err)
	}

	if !parsed.AsTime().Equal(now) {
		t.Errorf("expected parsed time to equal %v, got %v", now, parsed.AsTime())
	}
}

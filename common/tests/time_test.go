package commontests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/flowswiss/goclient/common"
)

func TestTimeFormat(t *testing.T) {
	testCases := []struct {
		Formatted string
		Expected  time.Time
	}{
		{Formatted: "2020-12-23T17:53:11+0000", Expected: time.Unix(1608745991, 0)},
		{Formatted: "2020-12-23T19:53:11+0200", Expected: time.Unix(1608745991, 0)},
		{Formatted: "2020-12-23T12:53:11-0500", Expected: time.Unix(1608745991, 0)},
		{Formatted: "2020-12-23T19:38:11+0145", Expected: time.Unix(1608745991, 0)},
		{Formatted: "2021-02-03T12:37:42+0000", Expected: time.Unix(1612355862, 0)},
		{Formatted: "2020-12-23T17:53:11+00:00", Expected: time.Unix(1608745991, 0)},
		{Formatted: "2020-12-23T19:53:11+02:00", Expected: time.Unix(1608745991, 0)},
		{Formatted: "2020-12-23T12:53:11-05:00", Expected: time.Unix(1608745991, 0)},
	}

	for _, testCase := range testCases {
		t.Run("parse timestamp "+testCase.Formatted, func(t *testing.T) {
			var val common.Time

			err := json.Unmarshal([]byte(`"`+testCase.Formatted+`"`), &val)
			if err != nil {
				t.Fatal(err)
			}

			if !val.AsTime().Equal(testCase.Expected) {
				t.Errorf("expected %v to equal %v", val.AsTime(), testCase.Expected)
			}
		})
	}
}

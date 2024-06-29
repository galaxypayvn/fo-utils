package customtype

import (
	"encoding/json"
	"errors"
	"time"
)

var timeLayouts = []string{time.RFC3339Nano, time.RFC3339, time.DateTime, time.DateOnly}

type Time struct {
	time.Time
}

func (t *Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return json.Marshal(nil)
	}

	return t.Time.UTC().MarshalJSON()
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || (len(data) == 2 && data[0] == '"' && data[1] == '"') {
		return nil
	}

	data = data[len(`"`) : len(data)-len(`"`)]
	var tm time.Time
	var err error
	for _, layout := range timeLayouts {
		tm, err = time.Parse(string(layout), string(data))
		if err == nil {
			t.Time = tm
			break
		}
	}

	if err != nil {
		return errors.New("unknown time layout")
	}

	return nil
}

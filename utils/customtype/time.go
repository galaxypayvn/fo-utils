package customtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
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
			t.Time = tm.UTC()
			break
		}
	}

	if err != nil {
		return errors.New("unknown time layout")
	}

	return nil
}

func (d *Time) Scan(value any) error {
	var err error
	switch s := value.(type) {
	case time.Time:
		d.Time = s
	case string:
		d.Time, err = time.Parse(time.RFC3339Nano, s)
		if err != nil {
			return err
		}
	case []byte:
		d.Time, err = time.Parse(time.RFC3339Nano, string(s))
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan type %T into CustomTime", value)
	}

	d.Time = d.Time.UTC()
	return nil
}

func (d Time) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil
	}
	return d.Time.UTC().Format(time.RFC3339Nano), nil
}

package uttime

import "time"

func ParseTimeLocation(t time.Time, locationName string) time.Time {
	// Load the target location (timezone)
	location, err := time.LoadLocation(locationName)
	if err != nil {
		return t
	}

	// Convert the given time to the target timezone
	localTime := t.In(location)

	return localTime
}

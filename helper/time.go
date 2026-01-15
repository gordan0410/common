package helper

import (
	"fmt"
	"time"
)

const (
	second = 1
	minute = 60 * second
	hour   = 60 * minute
)

var locationCache map[int]*time.Location
var availableOffsets []float64
var availableOffsetMinutes []int

func init() {
	locationCache = make(map[int]*time.Location)

	// 所有支援的時區偏移
	availableOffsets = []float64{
		-12.0, -11.0, -10.0, -9.5, -9.0,
		-8.0, -7.0, -6.0, -5.0, -4.5,
		-4.0, -3.5, -3.0, -2.0, -1.0,
		0.0, 1.0, 2.0, 3.0, 3.5,
		4.0, 4.5, 5.0, 5.5, 5.75,
		6.0, 6.5, 7.0, 8.0, 8.75,
		9.0, 9.5, 10.0, 10.5, 11.0,
		12.0, 12.75, 13.0, 14.0,
	}

	for _, offset := range availableOffsets {
		totalMinutes := int(offset * 60)
		availableOffsetMinutes = append(availableOffsetMinutes, totalMinutes)
		locationCache[totalMinutes] = time.FixedZone(OffsetMinutesToLabel(totalMinutes), totalMinutes*60)
	}
}

// GetLocation returns a pointer to a time.Location corresponding to the provided offset in hours.
//
// Parameters:
//   - offset: The timezone offset in minutes. Valid offsets are between -12 and 14.
//
// Returns:
//   - A pointer to the corresponding time.Location if offset is within the valid range;
//     otherwise, it returns time.UTC.
func GetLocation(offsetMinutes int) *time.Location {
	if location, ok := locationCache[offsetMinutes]; ok {
		return location
	}

	return time.UTC
}

// GetUTCOffsetHours returns the UTC offset in hours for the given time.Time object.
//
// Parameters:
//   - t: A time.Time object.
//
// Returns:
//   - The UTC offset in hours as an integer.
func GetUTCOffsetHours(t time.Time) int {
	_, offset := t.Zone()
	return offset / hour
}

func GetUTCOffsetMinutes(t time.Time) int {
	_, offset := t.Zone()
	return offset / 60
}

func OffsetMinutesToLabel(offsetMinutes int) string {
	sign := "+"
	absMinutes := offsetMinutes
	if offsetMinutes < 0 {
		sign = "-"
		absMinutes = -offsetMinutes
	}

	hours := absMinutes / 60
	minutes := absMinutes % 60

	return fmt.Sprintf("utc%s%d:%02d", sign, hours, minutes)
}

func GetAvailableOffsetMinutes() []int {
	return availableOffsetMinutes
}

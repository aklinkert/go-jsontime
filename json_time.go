package jsontime

import (
	"fmt"
	"time"
)

// JSONTime handles parsing and formatting timestamps according the RFC3339/ISO8601 standard
type JSONTime struct {
	time.Time
}

// String returns a string representation of the time.
func (t JSONTime) String() string {
	return t.Format(time.RFC3339)
}

// MarshalJSON formats the timestamp as JSON
func (t JSONTime) MarshalJSON() ([]byte, error) {
	date := fmt.Sprintf("%q", t.String())
	return []byte(date), nil
}

// Now returns the current time as JSONTime
func Now() JSONTime {
	return JSONTime{
		Time: time.Now(),
	}
}

// NowPtr returns the current time as pointer to JSONTime
func NowPtr() *JSONTime {
	return Ptr(Now())
}

// Ptr is a convenience method that returns a pointer to the given JSONTime
func Ptr(t JSONTime) *JSONTime {
	return &t
}

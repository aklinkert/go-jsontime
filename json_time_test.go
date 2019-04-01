package jsontime

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

type testJSON struct {
	Test JSONTime `json:"test"`
}

type testPtrJSON struct {
	Test *JSONTime `json:"test"`
}

var (
	jsonDate     = []byte("{\"test\":\"2018-08-18T10:31:17+02:00\"}")
	jsonDateNull = []byte("{\"test\":null}")
	rawDate      = "2018-08-18T10:31:17+02:00"
)

func getTestTime(t *testing.T) JSONTime {
	date, err := time.Parse(time.RFC3339, rawDate)
	if err != nil {
		t.Fatal(err)
	}

	return JSONTime{Time: date}
}

func getTestTimePtr(t *testing.T) *JSONTime {
	return Ptr(getTestTime(t))
}

func TestJSONTime_MarshalJSON(t *testing.T) {
	test := &testJSON{
		Test: JSONTime(getTestTime(t)),
	}

	b, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(b, jsonDate) {
		t.Errorf("Expected to get %q, got %q", string(jsonDate), string(b))
	}
}

func TestJSONTime_UnmarshalJSON(t *testing.T) {
	test := &testJSON{}
	testTime := getTestTime(t)

	if err := json.Unmarshal(jsonDate, test); err != nil {
		t.Fatal(err)
	}

	if !test.Test.Equal(testTime.Time) {
		t.Errorf("Dates are not equal: %v, %v", test.Test, testTime)
	}
}

func TestJSONTime_UnmarshalJSON_Null(t *testing.T) {
	test := &testJSON{}

	if err := json.Unmarshal(jsonDateNull, test); err != nil {
		t.Fatal(err)
	}

	if !test.Test.IsZero() {
		t.Errorf("Expected  time to be zero, got %v", test.Test)
	}
}

func TestJSONTime_UnmarshalJSON_Ptr(t *testing.T) {
	test := &testPtrJSON{}
	testTime := getTestTime(t)

	if err := json.Unmarshal(jsonDate, test); err != nil {
		t.Fatal(err)
	}

	if !test.Test.Equal(testTime.Time) {
		t.Errorf("Dates are not equal: %v, %v", test.Test, testTime)
	}
}

func TestJSONTime_UnmarshalJSON_PtrNull(t *testing.T) {
	test := &testPtrJSON{}

	if err := json.Unmarshal(jsonDateNull, test); err != nil {
		t.Fatal(err)
	}

	if test.Test != nil {
		t.Errorf("Expected  time to be nil, got %v", test.Test)
	}
}

func TestJSONTime_MarshalJSON_Ptr(t *testing.T) {
	test := &testPtrJSON{
		Test: getTestTimePtr(t),
	}

	b, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(b, jsonDate) {
		t.Errorf("Expected to get %q, got %q", string(jsonDate), string(b))
	}
}

func TestNow(t *testing.T) {
	now := Now()

	if now.Time.IsZero() {
		t.Error("Expected time on now to be set, but is zero")
	}
}

func TestNowPtr(t *testing.T) {
	now := NowPtr()

	if now.Time.IsZero() {
		t.Error("Expected time on now to be set, but is zero")
	}
}

func TestPtr(t *testing.T) {
	test := Now()
	now := Ptr(test)

	if *now != test {
		t.Error("Expected Ptr to return pointer to given object, but doesn't")
	}
}

package resterr

import "testing"

func TestErrorWithDefaultStatusCode(t *testing.T) {
	err := New("TEST ERROR")
	if err.Error() != "TEST ERROR" {
		t.Error("Error message is different from expected")
	}
	if err.StatusCode() != 500 {
		t.Error("Error status code is different from expected")
	}
}

func TestErrorCreationWithStatusCode(t *testing.T) {
	err := New("TEST ERROR").WithStatusCode(404)
	if err.StatusCode() != 404 {
		t.Error("Error status code is different from expected")
	}
}

func TestErrorWithFormat(t *testing.T) {
	err := Errorf("%s ERROR", "TEST")
	if err.Error() != "TEST ERROR" {
		t.Error("Error message is different from expected")
	}
}

func TestErrorWithFormatAndStatusCode(t *testing.T) {
	err := ErrorfWithStatusCode(404, "%s ERROR", "TEST")
	if err.Error() != "TEST ERROR" {
		t.Error("Error message is different from expected")
	}
	if err.StatusCode() != 404 {
		t.Error("Error status code is different from expected")
	}
}

func TestSettingInvalidStatusCode(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else if r.(string) != "Invalid status code" {
			t.Error("Invalid error message")
		}
	}()
	New("TEST ERROR").WithStatusCode(0)
}

func TestWithStatusCodeInError(t *testing.T) {
	ShowStatusCodeInError(true)
	err := New("TEST ERROR")
	if err.Error() != "500: TEST ERROR" {
		t.Error("Error message is different from expected")
	}
}

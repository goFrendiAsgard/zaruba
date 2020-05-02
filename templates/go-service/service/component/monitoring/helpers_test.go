package monitoring

import "testing"

// TestGetHTTPCodeByStatusTrue
func TestGetHTTPCodeByStatusTrue(t *testing.T) {
	httpStatus := getHTTPCodeByStatus(true)
	if httpStatus != 200 {
		t.Errorf("Unexpected result: %d", httpStatus)
	}
}

// TestGetHTTPCodeByStatusFalse
func TestGetHTTPCodeByStatusFalse(t *testing.T) {
	httpStatus := getHTTPCodeByStatus(false)
	if httpStatus != 500 {
		t.Errorf("Unexpected result: %d", httpStatus)
	}
}

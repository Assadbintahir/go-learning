package api_v1

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHello(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleHello)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v expected %v",
			status, http.StatusOK)
	}

	expected := `Hello, Asad — From RYK`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v expected %v",
			rr.Body.String(), expected)
	}

}

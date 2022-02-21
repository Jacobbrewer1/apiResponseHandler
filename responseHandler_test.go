package apiResponseHandler

import "testing"

func Test_DecodeResponse(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		expected response
	}{
		{"Response code 200", 200, Response200},
		{"Response code 400", 400, Response400},
		{"Response code 403", 403, Response403},
		{"Response code 404", 404, Response404},
		{"Response code 414", 414, Response414},
		{"Response code 429", 429, Response429},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse := DecodeResponse(tt.code)
			if gotResponse != tt.expected {
				t.Errorf("DecodeResponse(tt.code) = %v, expected %v", gotResponse, tt.expected)
			}
		})
	}
}

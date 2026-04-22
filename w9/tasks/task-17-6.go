package main

import (
	"errors"
	"fmt"
)

type RequestError struct {
	StatusCode int
	Cause      error
}

func (re *RequestError) Error() string {
	return fmt.Sprintf("request error status code: %d, cause: %s\n", re.StatusCode, re.Cause.Error())
}

func (re *RequestError) Unwrap() error {
	return re.Cause
}

var (
	ErrDatabaseDown     = errors.New("database connection error")
	ErrWrongToken       = errors.New("wrong token error")
	ErrMissingUserID    = errors.New("on userID is given")
	ErrEndpointNotFound = errors.New("wrong api URL")
)

var validEndpoints = []string{"/users", "/posts", "/settings"}

func handleRequest(endpoint, token, userID string) error {
	isValid := false
	for _, ep := range validEndpoints {
		if endpoint == ep {
			isValid = true
			break
		}
	}
	if !isValid {
		return &RequestError{StatusCode: 404, Cause: ErrEndpointNotFound}
	}

	if token != "valid-token" {
		return &RequestError{StatusCode: 401, Cause: ErrWrongToken}
	}

	if userID == "" {
		return &RequestError{StatusCode: 404, Cause: ErrMissingUserID}
	}

	if userID == "500" {
		return fmt.Errorf("Server error: %w\n", ErrDatabaseDown)
	}

	fmt.Println("Successful request handling")
	return nil
}

//func main() {
//
//	tests := []struct {
//		endpoint string
//		token    string
//		userID   string
//	}{
//		{"/users", "valid-token", "200"},
//		{"/users", "wrong-token", "200"},
//		{"/adams", "valid-token", "200"},
//		{"/users", "valid-token", ""},
//		{"/users", "valid-token", "500"},
//	}
//
//	for _, tt := range tests {
//		err := handleRequest(tt.endpoint, tt.token, tt.userID)
//		if err != nil {
//			if errors.Is(err, ErrDatabaseDown) {
//				fmt.Println("Check database connection")
//			}
//
//			var reqErr *RequestError
//			if errors.As(err, &reqErr) {
//				fmt.Println(reqErr.StatusCode, reqErr.Cause.Error())
//			}
//		}
//	}
//}

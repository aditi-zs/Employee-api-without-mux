package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//	func TestHandler(t *testing.T) {
//		req, err := http.NewRequest("GET", "/emp", nil)
//		if err != nil {
//			t.Fatal(err)
//		}
//		rr := httptest.NewRecorder()
//		handler := http.HandlerFunc(Employee)
//		handler.ServeHTTP(rr, req)
//		if status := rr.Code; status != http.StatusOK {
//			t.Errorf("handler returned wrong status code: got %v want %v",
//				status, http.StatusOK)
//		}
//		expected := `[{"id":"1","name":"Aditi","age":22,"address":"UP"}]`
//		if strings.TrimSpace(rr.Body.String()) != expected {
//			t.Errorf("handler returned unexpected body: got (%v) want (%v)",
//				rr.Body.String(), expected)
//		}
//	}
//
//	func TestPostEmployeeData(t *testing.T) {
//		tests := []struct {
//			description string
//			input       Emp
//			expRes      Emp
//			statusCode  int
//		}{
//			{"All entries are present",
//				Emp{
//					"1", "Aditi", 22, "UP",
//				},
//				Emp{
//					"1", "Aditi", 22, "UP",
//				},
//				201,
//			},
//		}
//
//		for _, tc := range tests {
//			val, _ := json.Marshal(tc.input) //go to json
//			req, err := http.NewRequest("POST", "/emp", bytes.NewReader(val))
//			if err != nil {
//				t.Errorf(err.Error())
//			}
//			//response recorder
//			rr := httptest.NewRecorder()
//			handler := http.HandlerFunc(Employee)
//			handler.ServeHTTP(rr, req)
//			//PostEmployeeData(resRec, req)
//			var actRes Emp
//			_ = json.Unmarshal(rr.Body.Bytes(), &actRes) //json to go
//			assert.Equal(t, tc.statusCode, rr.Code)
//			assert.Equal(t, tc.expRes, actRes)
//		}
//	}

func TestEmployee(t *testing.T) {
	tests := []struct {
		description string
		method      string
		input       string
		expRes      string
		statusCode  int
	}{
		{"for get method",
			"GET",
			"",
			`[{"id":"1","name":"Aditi","age":22,"address":"UP"}]`,
			200,
		},
		{"for post method",
			"POST",
			`{"id":"2","name":"Monika","age":23,"address":"UP"}`,
			`{"id":"2","name":"Monika","age":23,"address":"UP"}`,
			201,
		},
	}
	for _, tc := range tests {
		if tc.method == "GET" {
			req, err := http.NewRequest("GET", "/emp", nil)
			if err != nil {
				t.Errorf(err.Error())
			}
			//response recorder
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Employee)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tc.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			//expected := `[{"id":"1","name":"Aditi","age":22,"address":"UP"}]`
			if strings.TrimSpace(rr.Body.String()) != tc.expRes {
				t.Errorf("handler returned unexpected body: got (%v) want (%v)",
					rr.Body.String(), tc.expRes)
			}

		}
		if tc.method == "POST" {
			//val, _ := json.Marshal(tc.input) //go to json
			req, err := http.NewRequest("POST", "/emp", strings.NewReader(tc.input))
			if err != nil {
				t.Errorf(err.Error())
			}
			//response recorder
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Employee)
			handler.ServeHTTP(rr, req)

			assert.Equal(t, tc.statusCode, rr.Code)
			assert.Equal(t, tc.expRes, strings.TrimSpace(rr.Body.String()))
		}
	}
}

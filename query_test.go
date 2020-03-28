package main

import (
	"testing"
)

type testCase struct {
	text     string
	expected bool
}

func TestValidQuery(t *testing.T) {
	testSuite := []testCase{
		testCase{"add s", true},
		testCase{"create", false},
		testCase{"toggle 23", true},
		testCase{"edit 3", false},
		testCase{"toggle hsjdhfsjd", false},
		testCase{"new 34", true},
	}

	for _, tc := range testSuite {
		isValid, _ := validQuery(tc.text)
		if isValid != tc.expected {
			t.Error("expected:", tc.expected, "with text:", tc.text)
		}
	}

}

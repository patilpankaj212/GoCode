package main

import (
	"fmt"
	"strings"
	"testing"
)

type testCase struct {
	value    int
	expected string
}

func TestManyFizzbuzz(t *testing.T) {

	testCases := []testCase{
		{1, "1"},
		{3, "fizz"},
		{5, "buzz"},
		{8, "8"},
		{15, "fizzbuzz"},
		{17, "17"},
		{300, "fizzbuzz"},
		{93, "fizz"},
		{50, "buzz"},
	}

	for _, testCaseValue := range testCases {
		t.Run(fmt.Sprintf("%d", testCaseValue.value), func(t *testing.T) {
			actualValue := fizzbuzz(testCaseValue.value)
			if !strings.EqualFold(actualValue, testCaseValue.expected) {
				t.Logf("expected %s, but got: %s", testCaseValue.expected, actualValue)
				t.Fail()
			}
		})
	}
}

func TestFizzbuzz(t *testing.T) {

	value := fizzbuzz(1)
	if !strings.EqualFold(value, "1") {
		t.Log("expected 1, but got: ", value)
		t.Fail()
	}

	value = fizzbuzz(3)
	if !strings.EqualFold(value, "fizz") {
		t.Log("expected fizz, but got: ", value)
		t.Fail()
	}

	value = fizzbuzz(5)
	if !strings.EqualFold(value, "buzz") {
		t.Log("expected buzz, but got: ", value)
		t.Fail()
	}

	value = fizzbuzz(8)
	if !strings.EqualFold(value, "8") {
		t.Log("expected 8, but got: ", value)
		t.Fail()
	}

	value = fizzbuzz(15)
	if !strings.EqualFold(value, "fizzbuzz") {
		t.Log("expected fizzbuzz, but got: ", value)
		t.Fail()
	}

	value = fizzbuzz(17)
	if !strings.EqualFold(value, "17") {
		t.Log("expected 17, but got: ", value)
		t.Fail()
	}

	value = fizzbuzz(300)
	if !strings.EqualFold(value, "fizzbuzz") {
		t.Log("expected fizzbuzz, but got: ", value)
		t.Fail()
	}
}

func BenchmarkFizzbuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzbuzz(i)
	}
}

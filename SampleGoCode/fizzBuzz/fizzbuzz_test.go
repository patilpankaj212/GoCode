package main

import (
	"strings"
	"testing"
)

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

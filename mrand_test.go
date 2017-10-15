// Copyright (c) 2017 Mark K Mueller (markmueller.com). All rights reserved.
// GNU GPL version 3.

package mrand_test

import (
    "mrand"
	"testing"
)

const iterations = 4000000

func TestInt_noArgs(t *testing.T) {

	t.Logf("Test for a 32 bit value using no arguments")

	var i int

	i = mrand.Int()

	t.Logf("value: %d", i)

}

func TestInt64_noArgs(t *testing.T) {

	t.Logf("Test for a 64 bit value using no arguments")

	var i int64

	i = mrand.Int64()

	t.Logf("value: %d", i)

}

func TestInt_Outside_Range(t *testing.T) {
	min := 90
	max := min + 9
	t.Logf("Test for value outside the supplied range: %d to %d", min, max)
	for i := 0; i < iterations; i++ {
		var r int
		r = mrand.Int(min, max)
		if r < min || r > max {
			t.Errorf("Returned value out of range. Got %d", r)
			break
		}
	}
}

func TestInt_min_max(t *testing.T) {
	min := 90
	max := min + 9

	t.Logf("Test for minimum value: %d ", min)
	ok := false
	for i := 0; i < iterations; i++ {
		var r int
		r = mrand.Int(min, max)
		if r == min {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Never got the minimum value after %d iterations", iterations)
	}

	t.Logf("Test for maximum value: %d ", max)
	ok = false
	for i := 0; i < iterations; i++ {
		var r int
		r = mrand.Int(min, max)
		if r == max {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Never got the maximum value after %d iterations", iterations)
	}

}

func TestInt64_Outside_Range(t *testing.T) {
	min := int64(90)
	max := min + 9
	t.Logf("Test for value outside the supplied range: %d to %d", min, max)
	for i := 0; i < iterations; i++ {
		var r int64
		r = mrand.Int64(min, max)
		if r < min || r > max {
			t.Errorf("Returned value out of range. Got %d", r)
			break
		}
	}
}

func TestInt64_min_max (t *testing.T) {
	min := int64(90)
	max := min + 9

	t.Logf("Test for minimum value: %d ", min)
	ok := false
	for i := 0; i < iterations; i++ {
		var r int64
		r = mrand.Int64(min, max)
		if r == min {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Never got the minimum value after %d iterations", iterations)
	}

	t.Logf("Test for maximum value: %d", max)
	ok = false
	for i := 0; i < iterations; i++ {
		var r int64
		r = mrand.Int64(min, max)
		if r == max {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Never got the maximum value after %d iterations", iterations)
	}

}

func TestString_zero(t *testing.T) {

	t.Logf(`after %d iterations, we should eventually get a "0"`, iterations)
	ok := false
	for i := 0; i < iterations; i++ {
		var r string
		r = mrand.String(1)
		if r == "0" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf(`Never got a "0" after %d iterations`, iterations)
	}

	t.Logf(`after %d iterations, we should eventually get a "z"`, iterations)
	ok = false
	for i := 0; i < iterations; i++ {
		var r string
		r = mrand.String(1)
		if r == "z" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf(`Never got a "z" after %d iterations`, iterations)
	}

}

// Copyright (c) 2017 Mark K Mueller (markmueller.com). All rights reserved.
// GNU GPL version 3.

// Examples for mrand

package mrand

import (
    "testing"
    "math"
)

const iterations = 1000000

func TestInt_Outside_Range ( t *testing.T ) {
    t.Logf("Iterations: %d", iterations)
    min := 99
    max := 100
    t.Logf("Test for value outside the supplied range: %d to %d", min, max)
    for i:=0; i<iterations; i++ {
        var r int
        r = Int(min, max)
        if r < min || r > max {
            t.Errorf("Returned value out of range. Got %d", r)
            break
        }
    }
}

func TestInt_Minimum ( t *testing.T ) {
    ok := false
    min := math.MinInt32
    max := min + 1
    t.Logf("Test for minumum value: %d ", min)
    for i:=0; i<iterations; i++ {
        var r int
        r = Int(min, max)
        if r == min {
            ok = true
            break
        }
    }
    if !ok {
        t.Errorf("Never got the minumum value after %d iterations", iterations)
    }
}

func TestInt_Maximum ( t *testing.T ) {
    ok := false
    max := math.MaxInt32
    min := max - 1
    t.Logf("Test for maximum value: %d", max)
    for i:=0; i<iterations; i++ {
        var r int
        r = Int(min, max)
        if r == max {
            ok = true
            break
        }
    }
    if !ok {
        t.Errorf("Never got the maximum value after %d iterations", iterations)
    }
}








//  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //

func TestInt64_Outside_Range ( t *testing.T ) {
    t.Logf("Iterations: %d", iterations)
    min := int64(99)
    max := min + 1
    t.Logf("Test for value outside the supplied range: %d to %d", min, max)
    for i:=0; i<iterations; i++ {
        var r int64
        r = Int64(min, max)
        if r < min || r > max {
            t.Errorf("Returned value out of range. Got %d", r)
            break
        }
    }
}

func TestInt64_Minimum ( t *testing.T ) {
    ok := false
    min := int64(math.MinInt64)
    max := min + 1
    t.Logf("Test for minumum value: %d ", min)
    for i:=0; i<iterations; i++ {
        var r int64
        r = Int64(min, max)
        if r == min {
            ok = true
            break
        }
    }
    if !ok {
        t.Errorf("Never got the minumum value after %d iterations", iterations)
    }
}

func TestInt64_Maximum ( t *testing.T ) {
    max := int64(math.MaxInt64)
    min := int64(max - 1)
    t.Logf("Test for maximum value: %d", max)
    ok := false
    for i:=0; i<iterations; i++ {
        var r int64
        r = Int64(min, max)
        if r == max {
            ok = true
            break
        }
    }
    if !ok {
        t.Errorf("Never got the maximum value after %d iterations", iterations)
    }
}

//  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //

func TestString_Zero ( t *testing.T ) {
    t.Logf(`Test random string for "0"`)
    ok := false
    for i:=0; i<iterations; i++ {
        var r string
        // test is we ever get the first character "0"
        r = String(1)
        if r == "0" {
            ok = true
            break
        }
    }
    if !ok {
        t.Errorf(`Never got a "0" after %d iterations`, iterations)
    }
}

func TestString_Zed ( t *testing.T ) {
    t.Logf(`Test random string for "z"`)
    ok := false
    for i:=0; i<iterations; i++ {
        var r string
        // test is we ever get the last character "z"
        r = String(1)
        if r == "z" {
            ok = true
            break
        }
    }
    if !ok {
        t.Errorf(`Never got a "z" after %d iterations`, iterations)
    }
}


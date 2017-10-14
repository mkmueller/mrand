// Copyright (c) 2017 Mark K Mueller (markmueller.com). All rights reserved.
// GNU GPL version 3.

// Package mrand is a simple wrapper for math/rand providing a few functions to
// make things a little easier.
package mrand

import (
    "math/rand"
    "time"
)

const (
    characterSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    characterSetLen = 62
)

func init () {
    rand.Seed(time.Now().UnixNano())
}

// Int returns a random 32 bit integer between min and max inclusive.
func Int (min, max int ) int {
//    return int(min) + (rand.Int31() % int(max-min))
    return int(Int64(int64(min), int64(max)))
}

// Int64 returns a random 64 bit integer between min and max inclusive.
func Int64 (min, max int64 ) int64 {
    return min + (rand.Int63() % (1+max-min))
}

// String returns a random alpha-numeric string of length n.
func String (n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = characterSet[Int(0,characterSetLen-1)]
    }
    return string(b)
}

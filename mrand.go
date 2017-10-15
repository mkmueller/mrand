// Copyright (c) 2017 Mark K Mueller (markmueller.com). All rights reserved.
// GNU GPL version 3.

// Package mrand is a simple wrapper for math/rand. Int() and Int64() accept
// optional range parameters.
//
package mrand

import (
    "math/rand"
    "time"
)

const (
    characterSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    clen = 62
	MaxInt32  = 1<<31 - 1
	MaxInt64  = 1<<63 - 1
)

func init () {
    rand.Seed(time.Now().UnixNano())
}

// Int returns a random non-negative 32 bit integer.
//   Example:
//   rnd := mrand.Int()        // return a random int between 0 and MaxInt32 inclusive
//   rnd := mrand.Int(90,99)   // return a random int between 90 and 99 inclusive
//
func Int (a... int ) int {
    var min, max int
    if len(a) >= 2 {
        min = a[0]
        max = a[1]
        if min < 0 {
            min = 0
        }
    } else {
        min = 0
        max = MaxInt32
    }
    return int(Int64(int64(min), int64(max)))
}

// Int64 returns a random non-negative 64 bit integer.
//   Example:
//   rnd := mrand.Int64()        // return a random int64 between 0 and MaxInt64 inclusive
//   rnd := mrand.Int64(90,99)   // return a random int64 between 90 and 99 inclusive
//
func Int64 (a... int64 ) int64 {
    var min, max int64
    if len(a) >= 2 {
        min = a[0]
        max = a[1]
        if min < 0 {
            min = 0
        }
    } else {
        min = 0
        max = MaxInt64
    }
    return min + (rand.Int63() % (1+max-min))
}

// String returns a random alpha-numeric string of length n.
//   Example:
//   str := mrand.String(64)
//
func String (n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = characterSet[Int(0,clen-1)]
    }
    return string(b)
}

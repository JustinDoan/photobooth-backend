package main

import "testing"

func TestSample(t *testing.T) {
    if 1+1 != 2 {
        t.Error("Test failed: 1+1 did not equal 2")
    }
}
package test

import (
    "testing"
)

// Sample function to test
func add(x, y int) int {
    return x + y
}

// Test function
func TestAdd(t *testing.T) {
    result := add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("add(2, 3) = %d; want %d", result, expected)
    }
}

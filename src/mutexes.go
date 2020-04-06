package main

import (
	"fmt"
	"testing"
)

func Intmin(a, b int)  int {
	if a < b {
		return a
	}else {
		return b
	}
}

func TestIntmin(t *testing.T)  {
	ans := Intmin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}
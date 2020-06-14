package models

import (
	"testing"
)

func TestTenisBall_Length(t *testing.T) {
	c := Containers{1: 0, 2: 0}

	if c.Length() != 2 {
		t.Errorf("Expected: %v, Got: %v\n", 2, c.Length())
	}
}

func TestTenisBall_Load(t *testing.T) {
	c := Containers{1: 0, 2: 0}

	expectedFull := c.Load()
	actualFull := 0
	for k, v := range c {
		actualFull = k
		if v >= 3 {
			break
		}
	}
	if expectedFull != actualFull {
		t.Errorf("Expected: %v, Got: %v\n", expectedFull, actualFull)
	}
}

func TestTenisBall_GetTotalBall(t *testing.T) {
	c := Containers{1: 0, 2: 0}

	total := c.GetTotalBall(2)

	if total != 0 {
		t.Errorf("Expected: %v, Got: %v\n", 0, total)
	}
}

func TestTenisBall_GetRandContainer(t *testing.T) {
	c := Containers{1: 0, 2: 0}

	number := c.GetRandContainer(1, c.Length())

	if number < 1 && number > c.Length() {
		t.Errorf("Expected: %v, Got: %v\n", c.Length(), number)
	}
}

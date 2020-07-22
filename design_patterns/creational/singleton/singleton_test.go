package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil{
		t.Error("expected pointer to singleton, not nil")
	}

	counter2 := GetInstance()
	if counter2 != counter1{
		t.Errorf("expected same instance in counter2 but got a different one")
	}
}

func TestSingleton_AddOne(t *testing.T) {
	counter := GetInstance()
	currentCount := counter.AddOne()

	if currentCount != 1{
		t.Errorf("after calling the first time to count shoud get 1 but it is: %d\n", currentCount)
	}
}
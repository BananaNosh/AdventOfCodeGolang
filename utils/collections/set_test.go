package collections

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSet(t *testing.T) {
	rand.Seed(2022)
	intSet := NewComparingSet[int]()
	for i := 0; i < 19; i++ {
		randInt := rand.Intn(100)
		intSet.Add(randInt)
	}
	fmt.Println(intSet)
	intSet.Add(10)
	if !intSet.Has(10) {
		t.Errorf("Number not found %v", 10)
	}
	max := intSet.Max()
	min := intSet.Min()
	if max != 98 || min != 3 {
		t.Errorf("Max or min wrong %v %v", max, min)
	}
	all := CheckAll[int](intSet.Set, func(i int) bool {
		return i > 10
	})
	if all {
		t.Errorf("Not all numbers are greater than 10")
	}
	all = CheckAll[int](intSet.Set, func(i int) bool {
		return i < 100
	})
	if !all {
		t.Errorf("All numbers are smaller 100")
	}
}

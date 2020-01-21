package main

import (
	"reflect"
	"testing"
)

func testAssert(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSum(t *testing.T) {

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}

	testAssert(t, got, want)
}

func TestSumAllTails(t *testing.T) {

	t.Run("Collections with valid ints", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{3, 4})
		want := []int{2, 4}

		testAssert(t, got, want)
	})

	t.Run("Collections with blank values", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		testAssert(t, got, want)
	})

}

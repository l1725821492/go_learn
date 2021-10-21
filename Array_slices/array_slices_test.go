package Array_slices

import (
	"reflect"
	"testing"
)

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func Sum_2(numbers []string) string {
	var sum string
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumAllTails(numbers_tosum ...[]int) []int {
	var sums []int
	for _, numbers := range numbers_tosum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums

}
func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5, 8}

	got := Sum(numbers)
	want := 23

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
func TestStr(t *testing.T) {

	numbers := []string{"g", "o", "l", "a", "n", "g"}

	got := Sum_2(numbers)
	want := "golang"

	if got != want {
		t.Errorf("got %s want %s given, %v", got, want, numbers)
	}
}
func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{4, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

package iteration

import "testing"

func Repeat(charter string) string {
	var repeat_str string
	for i := 0; i < 5; i++ {
		repeat_str += charter
	}
	return repeat_str

}
func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

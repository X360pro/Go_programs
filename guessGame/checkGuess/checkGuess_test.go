package checkGuess

import "testing"

func TestCheckGueess(t *testing.T) {
	actual, _ := CheckGuess(47, 49)
	expected := "TRY HIGHER"
	if actual != expected {
		t.Error("Actual is :", actual, "But expected is : ", expected)
	}
}

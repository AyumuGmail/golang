package word

import "testing"

func TestPalidrome(t *testing.T) {
	if !IsPalindrome("detarated") {
		t.Error(`IsPalidrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalidrome("kayak") = false`)
	}
}

func TestNonPalidrome(t *testing.T) {
	if IsPalindrome("palidrome") {
		t.Error(`IsPalidrome("palidrome") = true`)
	}
}

func TestFrenchPalidrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalidrome("été") = false`)
	}
}

func TestCanalPalidrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalidrome(%q) = false`, input)
	}
}

func TestIsPalidrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detarated", true},
		{"A man, a plan, a canal: Panama", true},
		{"été", true},
		{"desserts", false},
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalidrome(%q) = %v", test.input, got)
		}
	}
}

func BenchmarkIsPalidrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man,a plan,a canal:Panama")
	}
}

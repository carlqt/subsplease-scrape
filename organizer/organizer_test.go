package organizer

import (
	"testing"
)

func TestFuzzy(t *testing.T) {
	tests := []struct {
		Input1 string
		Input2 string
	}{
		{Input1: "[SubsPlease] Vampire Dormitory - 06 (720p) [11903712].mkv", Input2: "[SubsPlease] Vampire Dormitory - 10 (720p) [A376EF02].mkv"},
		{Input1: "[SubsPlease] Vampire Dormitory - 06 (720p) [11903712].mkv", Input2: "[SubsPlease] Vampire Dormitory - 08 (720p) [6D861FC5].mkv"},
	}

	for _, test := range tests {
		actual := fuzzySearch(test.Input1, test.Input2)
		if actual > 25 {
			t.Fatalf("Actual (%d) is not less than or equal 25", actual)
		}
	}
}

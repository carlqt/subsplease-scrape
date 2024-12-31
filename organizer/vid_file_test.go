package organizer

import (
	"path"
	"testing"
)

type testData struct {
	Input    string
	Expected string
}

func TestTranslationGroup(t *testing.T) {
	tests := []testData{
		{Input: "[SubsPlease] Vampire Dormitory - 10 (720p) [A376EF02].mkv", Expected: "[SubsPlease]"},
		{Input: "[EMBER] Kimetsu no Yaiba - Hashira Geiko-hen - 07.mkv", Expected: "[EMBER]"},
		{Input: "[SubsPlease] Tower of God S2 - 14 (720p) [9FAE3662].mkv", Expected: "[SubsPlease]"},
	}

	for _, test := range tests {
		vFilePath := path.Join("Users", test.Input)
		vFile := vidFile{Path: vFilePath}

		actual := vFile.translationGroup()
		if test.Expected != actual {
			t.Fatalf("Expected = %s | Actual = %s", test.Expected, actual)
		}
	}
}

func TestTitle(t *testing.T) {
	tests := []testData{
		{Input: "[EMBER] Kimetsu no Yaiba - Hashira Geiko-hen - 07.mkv", Expected: "Kimetsu no Yaiba - Hashira Geiko-hen"},
		{Input: "[SubsPlease] Vampire Dormitory - 10 (720p) [A376EF02].mkv", Expected: "Vampire Dormitory"},
		{Input: "[SubsPlease] Tower of God S2 - 14 (720p) [9FAE3662].mkv", Expected: "Tower of God S2"},
	}

	for _, test := range tests {
		vFilePath := path.Join("Users", test.Input)
		vFile := vidFile{Path: vFilePath}

		actual := vFile.Title()
		if test.Expected != actual {
			t.Fatalf("Expected = %s | Actual = %s", test.Expected, actual)
		}
	}
}

func TestEndSection(t *testing.T) {
	tests := []testData{
		{Input: "[EMBER] Kimetsu no Yaiba - Hashira Geiko-hen - 07.mkv", Expected: "- 07.mkv"},
		{Input: "[SubsPlease] Vampire Dormitory - 10 (720p) [A376EF02].mkv", Expected: "- 10 (720p) [A376EF02].mkv"},
	}

	for _, test := range tests {
		vFilePath := path.Join("Users", test.Input)
		vFile := vidFile{Path: vFilePath}

		actual := vFile.endSection()
		if test.Expected != actual {
			t.Fatalf("Expected = %s | Actual = %s", test.Expected, actual)
		}
	}
}

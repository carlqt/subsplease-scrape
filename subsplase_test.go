package main

import "testing"

func TestGetFilename(t *testing.T) {
	testData := `inline; filename="%5BSubsPlease%5D%207th%20Time%20Loop%20-%2002%20%28480p%29%20%5B312B6329%5D.mkv.torrent"; filename*=UTF-8''%5BSubsPlease%5D%207th%20Time%20Loop%20-%2002%20%28480p%29%20%5B312B6329%5D.mkv.torrent`
	expected := "[SubsPlease] 7th Time Loop - 02 (480p) [312B6329].mkv.torrent"

	actual := GetFilename(testData)
	if expected != actual {
		t.Fatalf("Expected = %s\nActual = %s", expected, actual)
	}
}

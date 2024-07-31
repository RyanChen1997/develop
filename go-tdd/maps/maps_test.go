package maps

import "testing"

func TestSearch(t *testing.T) {
	m := map[string]string{"test": "This is Test"}
	got := Search(m, "test")
	assertString(t, got, "This is Test")
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

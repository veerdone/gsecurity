package gsecurity

import "testing"

func TestKeyMatch(t *testing.T) {
	t.Log(KeyMatch("user.add", "user.*"))
	t.Log(KeyMatch("user.add", "admin.*"))
}

func TestKeyMatch2(t *testing.T) {
	t.Log(KeyMatch2("ab.txt", "*.txt"))
	t.Log(KeyMatch2("ab.word", "*.txt"))
}

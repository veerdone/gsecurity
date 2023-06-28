package gsecurity

import "strings"

// KeyMatch
// Example: "user.add" match "user.*" return true
func KeyMatch(key1 string, key2 string) bool {
	if key2 == "*" {
		return true
	}
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}

	if len(key1) > i {
		return key1[:i] == key2[:i]
	}

	return key1 == key2[:i]
}

// KeyMatch2
// Example: "user.add" match "*.add" return true
func KeyMatch2(key1 string, key2 string) bool {
	if key2 == "*" {
		return true
	}

	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}

	return strings.HasSuffix(key1, key2[i+1:])
}

package gsecurity

import (
	"testing"
)

func TestRand32(t *testing.T) {
	t.Log(rand32())
}

func TestRand64(t *testing.T) {
	t.Log(rand64())
}

func TestRand128(t *testing.T) {
	t.Log(rand128())
}

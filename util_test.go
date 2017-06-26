package utils

import "testing"

func TestIsEmpty(t *testing.T) {
	t.Logf("string: a, %v", IsEmpty("a"))
	t.Logf("string: , %v", IsEmpty(""))
	t.Logf("int: 0, %v", IsEmpty(int(0)))
	t.Logf("int: 1, %v", IsEmpty(int(1)))
	t.Logf("int32: 0, %v", IsEmpty(int32(0)))
	t.Logf("int32: 1, %v", IsEmpty(int32(1)))
	t.Logf("int64: 0, %v", IsEmpty(int64(0)))
	t.Logf("int64: 1, %v", IsEmpty(int64(1)))
	t.Logf("byte: 0, %v", IsEmpty(byte(0)))
	t.Logf("byte: 1, %v", IsEmpty(byte(1)))

}
